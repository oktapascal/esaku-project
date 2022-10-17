package controllers

import (
	"esaku-project/app/auths/models/web"
	"esaku-project/app/auths/services"
	"esaku-project/helpers"
	"esaku-project/types"
	"net/http"
	"os"
	"time"
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

	cookieAccess := controller.Get("COOKIE_ACCESS_TOKEN")
	cookieRefresh := controller.Get("COOKIE_ACCESS_REFRESH_TOKEN")

	cookieConfig := helpers.NewCookieConfigImpl()

	dataAccess := types.M{"value": loginResponse.Token}
	cookieConfig.SetCookieToken(writer, cookieAccess, dataAccess, loginResponse.ExpirationAccess)

	dataRefresh := types.M{"value": loginResponse.RefreshToken}
	cookieConfig.SetCookieToken(writer, cookieRefresh, dataRefresh, loginResponse.ExpirationRefresh)

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *LoginControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
	cookieAccess := controller.Get("COOKIE_ACCESS_TOKEN")
	cookieRefresh := controller.Get("COOKIE_ACCESS_REFRESH_TOKEN")

	cookieConfig := helpers.NewCookieConfigImpl()

	dataAccess := types.M{}
	cookieConfig.SetCookieToken(writer, cookieAccess, dataAccess, time.Unix(0, 0))

	dataRefresh := types.M{}
	cookieConfig.SetCookieToken(writer, cookieRefresh, dataRefresh, time.Unix(0, 0))

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
