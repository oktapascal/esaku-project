package routes

import (
	controllers2 "esaku-project/app/auths/controllers"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeAuthRoutes(
	router *mux.Router,
	loginController controllers2.LoginController,
) {
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		json := map[string]string{
			"status": "OK",
		}

		writer.WriteHeader(http.StatusOK)
		helpers.WriteToResponseBodyJson(writer, json)
	})
	router.HandleFunc("/login", loginController.Login).Methods("POST")
	router.HandleFunc("/logout", loginController.Logout).Methods("POST")
}
