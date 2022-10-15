package services

import (
	"context"
	"database/sql"
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/repository"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"github.com/go-playground/validator/v10"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	Db              *sql.DB
	Validate        *validator.Validate
}

func NewLoginServiceImpl(loginRepository repository.LoginRepository, db *sql.DB, validate *validator.Validate) *LoginServiceImpl {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		Db:              db,
		Validate:        validate,
	}
}

func (service *LoginServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.Db.Begin()
	defer helpers.CommitOrRollback(tx, err)

	login, err := service.LoginRepository.Login(ctx, tx, request.Username)

	if err != nil {
		panic(exceptions.NewErrorNotFound(err.Error()))
	}

	return web.ToLoginResponse(login)
}
