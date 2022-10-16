package controllers

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/services"
	"esaku-project/helpers"
	"net/http"
	"time"
)

type LoginControllerImpl struct {
	LoginService services.LoginService
}

func NewLoginControllerImpl(loginService services.LoginService) *LoginControllerImpl {
	return &LoginControllerImpl{LoginService: loginService}
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

	helpers.SetCookieToken(writer, loginResponse.CookieAccess, loginResponse.Token, loginResponse.ExpirationAccess)
	helpers.SetCookieToken(writer, loginResponse.CookieRefresh, loginResponse.RefreshToken, loginResponse.ExpirationRefresh)

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *LoginControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	cookieAccess, cookieRefresh := controller.LoginService.Logout(request.Context())

	helpers.SetCookieToken(writer, cookieAccess, "", time.Unix(0, 0))
	helpers.SetCookieToken(writer, cookieRefresh, "", time.Unix(0, 0))

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
