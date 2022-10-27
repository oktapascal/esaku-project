package routes

import (
	"esaku-project/app/settings/controllers"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeFilterRoutes(
	router *mux.Router,
	unitController controllers.UnitController,
	karyawanController controllers.KaryawanController,
	kelompokMenuController controllers.KelompokMenuController,
) {
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		json := map[string]string{
			"status": "OK",
		}

		writer.WriteHeader(http.StatusOK)
		helpers.WriteToResponseBodyJson(writer, json)
	})
	//	Unit
	router.HandleFunc("v1/unit", unitController.Filter).Methods("GET")
	//	Karyawan
	router.HandleFunc("/v1/karyawan", karyawanController.Filter).Methods("GET")
	//	Kelompok Menu
	router.HandleFunc("/v1/kelompok-menu", kelompokMenuController.Filter).Methods("GET")
}
