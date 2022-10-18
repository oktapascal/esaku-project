package controllers

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/services"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"esaku-project/middlewares"
	"net/http"
	"os"
)

type LoginControllerImpl struct {
	LoginService services.LoginService
}

func NewLoginControllerImpl(loginService services.LoginService) *LoginControllerImpl {
	return &LoginControllerImpl{LoginService: loginService}
}

func (controller *LoginControllerImpl) Get(key string) string {
	return os.Getenv(key)
}

func (controller *LoginControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	loginRequest := web.LoginRequest{}
	helpers.ReadFromRequestBodyJson(request, &loginRequest)

	loginResponse := controller.LoginService.Login(request.Context(), loginRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]string{
			"token":         loginResponse.Token,
			"refresh_token": loginResponse.RefreshToken,
		},
	}

	err := middlewares.GenerateTokenAndCookie(loginResponse, writer)

	if err != nil {
		panic(exceptions.NewErrorUnauthorized("token is incorrect"))
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *LoginControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	err := middlewares.DeleteCookie(writer)

	if err != nil {
		panic(err.Error())
	}

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
