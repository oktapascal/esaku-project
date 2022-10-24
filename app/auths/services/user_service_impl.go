package services

import (
	"context"
	"database/sql"
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/repository"
	"esaku-project/app/settings/models/domain"
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

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Db             *sql.DB
	Validate       *validator.Validate
	S3             *s3.Client
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate, s3 *s3.Client) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Db:             db,
		Validate:       validate,
		S3:             s3,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserRequest) {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi
	Nik := claims.Nik

	user := domain.Karyawan{
		Nik:        Nik,
		KodeLokasi: kodeLokasi,
		Nama:       request.Nama,
		Jabatan:    request.Jabatan,
		NoTelp:     request.NoTelp,
		Email:      request.Email,
	}

	service.UserRepository.Update(ctx, tx, user)
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, request web.PasswordRequest) {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi
	Nik := claims.Nik

	password, err := helpers.HashPassword(request.Password)
	helpers.PanicIfError(err)

	karyawan := domain.Karyawan{
		Nik:        Nik,
		KodeLokasi: kodeLokasi,
	}
	user := domain.HakAkses{
		Password: password,
		Karyawan: karyawan,
	}

	service.UserRepository.UpdatePassword(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context) web.UserResponse {
	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi
	Nik := claims.Nik

	user, err := service.UserRepository.FindById(ctx, tx, Nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	if user.Foto != "" {
		AwsBucket := os.Getenv("AWS_BUCKET")
		AwsRegion := os.Getenv("AWS_REGION")

		user.Foto = "https://" + AwsBucket + "." + "s3-" + AwsRegion + ".amazonaws.com/" + user.Foto
	}

	return web.ToUserResponse(user)
}

func (service *UserServiceImpl) UploadImage(ctx context.Context, request web.UserUploadRequest) {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	token := ctx.Value("token").(*jwt.Token)
	claims := token.Claims.(*types.Claims)

	kodeLokasi := claims.KodeLokasi
	Nik := claims.Nik

	user, err := service.UserRepository.FindById(ctx, tx, Nik, kodeLokasi)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	_, err = service.S3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET")),
		Key:    aws.String(fmt.Sprintf("dev/%s", user.Foto)),
	})

	helpers.PanicIfError(err)

	file, err := request.Foto.Open()
	helpers.PanicIfError(err)

	fileName := fmt.Sprintf("profile-%s-%s", Nik, request.Foto.Filename)

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

	user.Foto = fileName

	service.UserRepository.UploadImage(ctx, tx, user)
}
