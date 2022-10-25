package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type KelompokMenuControllerImpl struct {
	KelompokMenuService services.KelompokMenuService
}

func NewKelompokMenuControllerImpl(kelompokMenuService services.KelompokMenuService) *KelompokMenuControllerImpl {
	return &KelompokMenuControllerImpl{KelompokMenuService: kelompokMenuService}
}

func (controller *KelompokMenuControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	kelompokMenuRequest := web.KelompokMenuSaveRequest{}
	helpers.ReadFromRequestBodyJson(request, &kelompokMenuRequest)

	kelompokMenuResponse := controller.KelompokMenuService.Save(request.Context(), kelompokMenuRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   kelompokMenuResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KelompokMenuControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	kelompokMenuRequest := web.KelompokMenuUpdateRequest{}
	helpers.ReadFromRequestBodyJson(request, &kelompokMenuRequest)

	vars := mux.Vars(request)
	kodeKlp := vars["kodeKlp"]
	kelompokMenuRequest.KodeKlp = kodeKlp

	kelompokMenuResponse := controller.KelompokMenuService.Update(request.Context(), kelompokMenuRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   kelompokMenuResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KelompokMenuControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeKlp := vars["kodeKlp"]

	controller.KelompokMenuService.Delete(request.Context(), kodeKlp)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KelompokMenuControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	kodeKlp := vars["kodeKlp"]

	klpMenuResponse := controller.KelompokMenuService.FindById(request.Context(), kodeKlp)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   klpMenuResponse,
	}
	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KelompokMenuControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	klpMenuResponses := controller.KelompokMenuService.FindAll(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   klpMenuResponses,
	}
	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KelompokMenuControllerImpl) Filter(writer http.ResponseWriter, request *http.Request) {
	klpMenuResponses := controller.KelompokMenuService.Filter(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   klpMenuResponses,
	}
	helpers.WriteToResponseBodyJson(writer, webResponse)
}
