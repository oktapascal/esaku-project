package controllers

import "net/http"

type KelompokMenuController interface {
	Save(writer http.ResponseWriter, request *http.Request)
	Update(writer http.ResponseWriter, request *http.Request)
	Delete(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request)
	FindAll(writer http.ResponseWriter, request *http.Request)
}
