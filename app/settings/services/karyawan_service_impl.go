package services

import (
	"bytes"
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/repository"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"esaku-project/types"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type KaryawanServiceImpl struct {
	KaryawanRepository repository.KaryawanRepository
	Db                 *sql.DB
	Validate           *validator.Validate
	S3                 *s3.Client
}

func NewKaryawanServiceImpl(karyawanRepository repository.KaryawanRepository, db *sql.DB, validate *validator.Validate, s3 *s3.Client) *KaryawanServiceImpl {
	return &KaryawanServiceImpl{KaryawanRepository: karyawanRepository, Db: db, Validate: validate, S3: s3}
}

func (service *KaryawanServiceImpl) Save(ctx context.Context, request web.KaryawanSaveRequest) web.KaryawanListResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawan := domain.Karyawan{
		Nik:        request.Nik,
		KodeLokasi: kodeLokasi,
		Nama:       request.Nama,
		KodeUnit:   request.KodeUnit,
		FlagAktif:  request.FlagAktif,
		Jabatan:    request.Jabatan,
		Alamat:     request.Alamat,
		NoTelp:     request.NoTelp,
		NoHp:       request.NoHp,
		Email:      request.Email,
	}

	karyawan = service.KaryawanRepository.Save(ctx, tx, karyawan)

	return web.ToKaryawanListResponse(karyawan)
}

func (service *KaryawanServiceImpl) Update(ctx context.Context, request web.KaryawanUpdateRequest) web.KaryawanListResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, request.Nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	karyawan.Nama = request.Nama
	karyawan.Alamat = request.Alamat
	karyawan.NoHp = request.NoHp
	karyawan.NoTelp = request.NoTelp
	karyawan.Email = request.Email
	karyawan.KodeUnit = request.KodeUnit
	karyawan.FlagAktif = request.FlagAktif
	karyawan.Jabatan = request.Jabatan

	karyawan = service.KaryawanRepository.Update(ctx, tx, karyawan)

	return web.ToKaryawanListResponse(karyawan)
}

func (service *KaryawanServiceImpl) Delete(ctx context.Context, nik string) {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	_, err = service.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fmt.Sprintf("dev/%s", karyawan.Foto)),
	})

	helpers.PanicIfError(err)

	service.KaryawanRepository.Delete(ctx, tx, karyawan)
}

func (service *KaryawanServiceImpl) FindById(ctx context.Context, nik string) web.KaryawanDetailResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	if karyawan.Foto != "" {
		AwsBucket := os.Getenv("AWS_BUCKET")
		AwsRegion := os.Getenv("AWS_REGION")

		karyawan.Foto = "https://" + AwsBucket + "." + "s3-" + AwsRegion + ".amazonaws.com/" + karyawan.Foto
	}

	return web.ToKaryawanDetailResponse(karyawan)
}

func (service *KaryawanServiceImpl) FindAll(ctx context.Context) []web.KaryawanListResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawans := service.KaryawanRepository.FindAll(ctx, tx, kodeLokasi)

	return web.ToKaryawanListResponses(karyawans)
}

func (service *KaryawanServiceImpl) UploadImage(ctx context.Context, request web.KaryawanUploadRequest) web.KaryawanUploadResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	file, err := request.Foto.Open()
	helpers.PanicIfError(err)

	buff := new(bytes.Buffer)
	_, err = buff.ReadFrom(file)
	if err != nil {
		panic(exceptions.NewErrorBadRequest(err.Error()))
	}

	bytesString := buff.Bytes()

	_, err = helpers.CheckOnlyImage(bytesString)
	if err != nil {
		panic(exceptions.NewErrorBadRequest(err.Error()))
	}

	kodeLokasi := claims.KodeLokasi
	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, request.Nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	_, err = service.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fmt.Sprintf("dev/%s", karyawan.Foto)),
	})

	helpers.PanicIfError(err)

	fileName := fmt.Sprintf("profile-%s-%s", request.Nik, request.Foto.Filename)

	uploader := manager.NewUploader(service.S3)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(os.Getenv("AWS_BUCKET")),
		ACL:           "bucket-owner-full-control",
		Body:          file,
		ContentLength: request.Foto.Size,
		Key:           aws.String(fmt.Sprintf("dev/%s", fileName)),
	})

	helpers.PanicIfError(err)

	//noinspection GoUnhandledErrorResult
	defer file.Close()

	karyawan.Foto = fileName

	service.KaryawanRepository.UploadImage(ctx, tx, karyawan)

	return web.ToKaryawanUploadResponse(fileName)
}

func (service *KaryawanServiceImpl) Filter(ctx context.Context) []web.FilterKaryawanResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi

	karyawans := service.KaryawanRepository.FindAll(ctx, tx, kodeLokasi)

	return web.ToFilterKaryawanResponses(karyawans)
}
