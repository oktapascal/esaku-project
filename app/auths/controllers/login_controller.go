package controllers

import "net/http"

type LoginController interface {
	Login(writer http.ResponseWriter, request *http.Request)
	Logout(writer http.ResponseWriter, request *http.Request)
}
