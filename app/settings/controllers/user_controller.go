package controllers

import "net/http"

type UserController interface {
	Update(writer http.ResponseWriter, request *http.Request)
	UpdatePassword(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request)
	UploadImage(writer http.ResponseWriter, request *http.Request)
}
