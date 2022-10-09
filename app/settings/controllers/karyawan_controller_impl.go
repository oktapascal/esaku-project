package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

type KaryawanControllerImpl struct {
	KaryawanService services.KaryawanService
}

func NewKaryawanControllerImpl(karyawanService services.KaryawanService) *KaryawanControllerImpl {
	return &KaryawanControllerImpl{KaryawanService: karyawanService}
}

func (controller *KaryawanControllerImpl) Save(writer http.ResponseWriter, request *http.Request) {
	karyawanRequest := web.KaryawanSaveRequest{}
	helpers.ReadFromRequestBodyJson(request, karyawanRequest)

	karyawanResponse := controller.KaryawanService.Save(request.Context(), karyawanRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   karyawanResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KaryawanControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	karyawanRequest := web.KaryawanUpdateRequest{}
	helpers.ReadFromRequestBodyJson(request, karyawanRequest)

	vars := mux.Vars(request)
	nik := vars["nik"]
	karyawanRequest.Nik = nik

	karyawanResponse := controller.KaryawanService.Update(request.Context(), karyawanRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   karyawanResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KaryawanControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nik := vars["nik"]

	controller.KaryawanService.Delete(request.Context(), nik)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KaryawanControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	nik := vars["nik"]

	karyawanResponse := controller.KaryawanService.FindById(request.Context(), nik)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   karyawanResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KaryawanControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	karyawanResponses := controller.KaryawanService.FindAll(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   karyawanResponses,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *KaryawanControllerImpl) UploadImage(writer http.ResponseWriter, request *http.Request) {
	karyawanRequest := web.KaryawanUploadRequest{}
	helpers.ReadFromRequestBodyMultipart(request)

	_, fileHeader, err := request.FormFile("foto")
	if err != nil {
		panic(exceptions.NewErrorBadRequest(err.Error()))
	}

	karyawanRequest.Nik = request.PostFormValue("nik")
	karyawanRequest.Foto = fileHeader

	karyawanResponse := controller.KaryawanService.UploadImage(request.Context(), karyawanRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   karyawanResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
