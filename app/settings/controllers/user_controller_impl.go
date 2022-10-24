package controllers

import (
	"esaku-project/app/settings/models/web"
	"esaku-project/app/settings/services"
	"esaku-project/exceptions"
	"esaku-project/helpers"
	"net/http"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func NewUserControllerImpl(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.UserRequest{}
	helpers.ReadFromRequestBodyJson(request, &userRequest)

	controller.UserService.Update(request.Context(), userRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UserControllerImpl) UpdatePassword(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.PasswordRequest{}
	helpers.ReadFromRequestBodyJson(request, &userRequest)

	controller.UserService.UpdatePassword(request.Context(), userRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	userResponse := controller.UserService.FindById(request.Context())

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}

func (controller *UserControllerImpl) UploadImage(writer http.ResponseWriter, request *http.Request) {
	userRequest := web.UserUploadRequest{}
	helpers.ReadFromRequestBodyMultipart(request)

	_, fileHeader, err := request.FormFile("foto")
	if err != nil {
		panic(exceptions.NewErrorBadRequest(err.Error()))
	}

	userRequest.Foto = fileHeader

	controller.UserService.UploadImage(request.Context(), userRequest)

	webResponse := helpers.JsonResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
