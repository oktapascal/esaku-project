package services

import (
	"context"
	"database/sql"
	"esaku-project/app/settings/models/domain"
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/repository"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-playground/validator/v10"
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

	karyawan := domain.Karyawan{
		Nik:        request.Nik,
		KodeLokasi: "99",
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

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, request.Nik, "99")

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

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, nik, "99")

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	_, err = service.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("devsai-s3"),
		Key:    aws.String(fmt.Sprintf("dev/%s", karyawan.Foto)),
	})

	helpers.PanicIfError(err)

	service.KaryawanRepository.Delete(ctx, tx, karyawan)
}

func (service *KaryawanServiceImpl) FindById(ctx context.Context, nik string) web.KaryawanDetailResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, nik, "99")

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToKaryawanDetailResponse(karyawan)
}

func (service *KaryawanServiceImpl) FindAll(ctx context.Context) []web.KaryawanListResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	karyawans := service.KaryawanRepository.FindAll(ctx, tx, "99")

	return web.ToKaryawanListResponses(karyawans)
}

func (service *KaryawanServiceImpl) UploadImage(ctx context.Context, request web.KaryawanUploadRequest) web.KaryawanUploadResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	karyawan, err := service.KaryawanRepository.FindById(ctx, tx, request.Nik, "99")

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	_, err = service.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String("devsai-s3"),
		Key:    aws.String(fmt.Sprintf("dev/%s", karyawan.Foto)),
	})

	helpers.PanicIfError(err)

	file, err := request.Foto.Open()
	helpers.PanicIfError(err)

	fileName := fmt.Sprintf("profile-%s-%s", request.Nik, request.Foto.Filename)

	uploader := manager.NewUploader(service.S3)

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String("devsai-s3"),
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
