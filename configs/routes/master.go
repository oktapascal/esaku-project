package routes

import (
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeMasterRoutes(router *mux.Router) {
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		json := map[string]string{
			"status": "OK",
		}

		writer.WriteHeader(http.StatusOK)
		helpers.WriteToResponseBodyJson(writer, json)
	})
}
