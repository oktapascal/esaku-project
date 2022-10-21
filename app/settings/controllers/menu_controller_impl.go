package controllers

import (
	"esaku-project/app/settings/services"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type MenuControllerImpl struct {
	MenuService services.MenuService
}

func NewMenuControllerImpl(menuService services.MenuService) *MenuControllerImpl {
	return &MenuControllerImpl{MenuService: menuService}
}

func (controller *MenuControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *MenuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller *MenuControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeKlp := vars["kodeKlp"]

	menuResponse := controller.MenuService.FindById(request.Context(), kodeKlp)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   menuResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
