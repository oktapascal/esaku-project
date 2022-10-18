package controllers

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/services"
	"esaku-project/bootstraps"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"esaku-project/types"
	"net/http"
)

type LoginControllerImpl struct {
	LoginService services.LoginService
	ConfigCookie bootstraps.Cookie
	JWTConfig    bootstraps.JWT
}

func NewLoginControllerImpl(loginService services.LoginService, configCookie bootstraps.Cookie, JWTConfig bootstraps.JWT) *LoginControllerImpl {
	return &LoginControllerImpl{
		LoginService: loginService,
		ConfigCookie: configCookie,
		JWTConfig:    JWTConfig,
	}
}

func (controller *LoginControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	loginRequest := web.LoginRequest{}
	helpers.ReadFromRequestBodyJson(request, &loginRequest)

	loginResponse := controller.LoginService.Login(request.Context(), loginRequest)

	accessToken, expirationAccess, err := controller.JWTConfig.GenerateAccessToken(loginResponse)
	if err != nil {
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	refreshToken, expirationRefresh, err := controller.JWTConfig.GenerateRefreshToken(loginResponse)
	if err != nil {
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	dataAccess := types.M{"value": accessToken}
	dataRefresh := types.M{"value": refreshToken}

	err = controller.ConfigCookie.CreateTokenAndCookie(loginResponse, dataAccess, dataRefresh, expirationAccess, expirationRefresh, writer)
	helpers.PanicIfError(err)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"token":         accessToken,
			"refresh_token": refreshToken,
		},
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *LoginControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	err := controller.ConfigCookie.DeleteCookie(writer)
	helpers.PanicIfError(err)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
