package routes

import (
	"esaku-project/app/settings/controllers"
	"github.com/gorilla/mux"
)

func InitializeFilterRoutes(
	router *mux.Router,
	unitController controllers.UnitController,
	karyawanController controllers.KaryawanController,
	kelompokMenuController controllers.KelompokMenuController,
) {
	//	Unit
	router.HandleFunc("v1/unit", unitController.Filter).Methods("GET")
	//	Karyawan
	router.HandleFunc("/v1/karyawan", karyawanController.Filter).Methods("GET")
	//	Kelompok Menu
	router.HandleFunc("/v1/kelompok-menu", kelompokMenuController.Filter).Methods("GET")
}
