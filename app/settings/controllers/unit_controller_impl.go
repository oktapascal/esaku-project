package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type UnitControllerImpl struct {
	UnitService services.UnitService
}

func NewUnitControllerImpl(unitService services.UnitService) *UnitControllerImpl {
	return &UnitControllerImpl{UnitService: unitService}
}

func (controller *UnitControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	unitRequest := web.UnitSaveRequest{}
	helpers.ReadFromRequestBodyJson(request, unitRequest)

	unitResponse := controller.UnitService.Save(request.Context(), unitRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   unitResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UnitControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	unitRequest := web.UnitUpdateRequest{}
	helpers.ReadFromRequestBodyJson(request, unitRequest)

	unitResponse := controller.UnitService.Update(request.Context(), unitRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   unitResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UnitControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeUnit := vars["kodeUnit"]

	controller.UnitService.Delete(request.Context(), kodeUnit)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UnitControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeUnit := vars["kodeUnit"]

	unitResponse := controller.UnitService.FindById(request.Context(), kodeUnit)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   unitResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UnitControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	unitResponse := controller.UnitService.FindAll(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   unitResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UnitControllerImpl) Filter(writer http.ResponseWriter, request *http.Request) {
	unitResponse := controller.UnitService.Filter(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   unitResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
