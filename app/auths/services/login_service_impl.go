package services

import (
	"context"
	"database/sql"
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/repository"
	"esaku-project/configs"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"github.com/go-playground/validator/v10"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	Db              *sql.DB
	Validate        *validator.Validate
	Config          configs.Config
	JwtConfig       helpers.ConfigJwt
}

func NewLoginServiceImpl(loginRepository repository.LoginRepository, db *sql.DB, validate *validator.Validate, config configs.Config, jwt helpers.ConfigJwt) *LoginServiceImpl {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		Db:              db,
		Validate:        validate,
		Config:          config,
		JwtConfig:       jwt,
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

	checkPassword := helpers.CheckPasswordHash(request.Password, login.Password)

	if !checkPassword {
		panic(exceptions.NewErrorBadRequest("password incorrect"))
	}

	jwtSecret := service.Config.Get("JWT_KEY_TOKEN")
	jwtRefresh := service.Config.Get("JWT_REFRESH_KEY_TOKEN")
	cookieAccess := service.Config.Get("COOKIE_ACCESS_TOKEN")
	cookieRefresh := service.Config.Get("COOKIE_ACCESS_REFRESH_TOKEN")

	tokenJwt, timeJwt, err := service.JwtConfig.GenerateJwtToken(ctx, jwtSecret, login)
	helpers.PanicIfError(err)

	tokenRefresh, timeRefresh, err := service.JwtConfig.GenerateJwtRefreshToken(ctx, jwtRefresh, login)
	helpers.PanicIfError(err)

	login.Token = tokenJwt
	login.RefreshToken = tokenRefresh
	login.ExpirationAccess = timeJwt
	login.ExpirationRefresh = timeRefresh
	login.CookieAccess = cookieAccess
	login.CookieRefresh = cookieRefresh

	return web.ToLoginResponse(login)
}
