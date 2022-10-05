package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type FormControllerImpl struct {
	FormService services.FormService
}

func NewFormControllerImpl(formService services.FormService) *FormControllerImpl {
	return &FormControllerImpl{FormService: formService}
}

func (controller *FormControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	formRequest := web.FormSaveRequest{}
	helpers.ReadFromRequestBodyJson(request, &formRequest)

	formResponse := controller.FormService.Save(request.Context(), formRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   formResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *FormControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	formRequest := web.FormUpdateRequest{}
	helpers.ReadFromRequestBodyJson(request, &formRequest)

	vars := mux.Vars(request)
	kodeForm := vars["kodeForm"]
	formRequest.KodeForm = kodeForm

	formResponse := controller.FormService.Update(request.Context(), formRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   formResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *FormControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeForm := vars["kodeForm"]

	controller.FormService.Delete(request.Context(), kodeForm)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *FormControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeForm := vars["kodeForm"]

	formResponse := controller.FormService.FindById(request.Context(), kodeForm)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   formResponse,
	}
	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *FormControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	formResponses := controller.FormService.FindAll(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   formResponses,
	}
	helpers.WriteToResponseBodyJson(writer, webResponse)
}
