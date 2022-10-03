package exceptions

import (
	"esaku-project/helpers"
	"esaku-project/models/web"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if badRequestError(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}

	if validationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorBadRequest)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.JsonResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   exception.Error,
		}

		helpers.WriteToResponseBodyJson(writer, webResponse)

		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var errorFields []web.FieldError
		var msg string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				msg = fmt.Sprintf("%s is required", err.Field())
			default:
				msg = exception.Error()
			}

			errorField := web.FieldError{
				Param:   err.Field(),
				Message: msg,
			}

			errorFields = append(errorFields, errorField)
		}

		webResponse := web.JsonResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   errorFields,
		}

		helpers.WriteToResponseBodyJson(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(ErrorNotFound)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.JsonResponse{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   exception.Error,
		}

		helpers.WriteToResponseBodyJson(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	logger := helpers.WriteLogging()

	logger.Error(err)
	webResponse := web.JsonResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err,
	}

	helpers.WriteToResponseBodyJson(writer, webResponse)
}
