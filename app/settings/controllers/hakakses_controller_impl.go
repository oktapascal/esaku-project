package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type HakAksesControllerImpl struct {
	HakAksesService services.HakAksesService
}

func NewHakAksesControllerImpl(hakAksesService services.HakAksesService) *HakAksesControllerImpl {
	return &HakAksesControllerImpl{HakAksesService: hakAksesService}
}

func (controller *HakAksesControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	aksesRequest := web.HakAksesSaveRequest{}
	helpers.ReadFromRequestBodyJson(request, aksesRequest)

	aksesResponse := controller.HakAksesService.Save(request.Context(), aksesRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   aksesResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *HakAksesControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	aksesRequest := web.HakAksesUpdateRequest{}
	helpers.ReadFromRequestBodyJson(request, aksesRequest)

	vars := mux.Vars(request)
	nik := vars["nik"]
	aksesRequest.Nik = nik

	aksesResponse := controller.HakAksesService.Update(request.Context(), aksesRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   aksesResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *HakAksesControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nik := vars["nik"]

	controller.HakAksesService.Delete(request.Context(), nik)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *HakAksesControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nik := vars["nik"]

	aksesResponse := controller.HakAksesService.FindById(request.Context(), nik)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   aksesResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *HakAksesControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	aksesResponses := controller.HakAksesService.FindAll(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   aksesResponses,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
