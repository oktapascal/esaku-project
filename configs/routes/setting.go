package routes

import (
	"esaku-project/app/settings/controllers"
	"github.com/gorilla/mux"
)

func InitializeSettingsRoutes(
	router *mux.Router,
	kelompokMenuController controllers.KelompokMenuController,
	formController controllers.FormController,
	unitController controllers.UnitController,
) {
	// Kelompok Menu
	router.HandleFunc("/kelompok-menu", kelompokMenuController.FindAll).Methods("GET")
	router.HandleFunc("/kelompok-menu/{kodeKlp}", kelompokMenuController.FindById).Methods("GET")
	router.HandleFunc("/kelompok-menu", kelompokMenuController.Save).Methods("POST")
	router.HandleFunc("/kelompok-menu/{kodeKlp}", kelompokMenuController.Update).Methods("PUT")
	router.HandleFunc("/kelompok-menu/{kodeKlp}", kelompokMenuController.Delete).Methods("DELETE")

	// Form
	router.HandleFunc("/form", formController.FindAll).Methods("GET")
	router.HandleFunc("/form/{kodeForm}", formController.FindById).Methods("GET")
	router.HandleFunc("/form", formController.Save).Methods("POST")
	router.HandleFunc("/form/{kodeForm}", formController.Update).Methods("PUT")
	router.HandleFunc("/form/{kodeForm}", formController.Delete).Methods("DELETE")

	// Unit
	router.HandleFunc("/unit", unitController.FindAll).Methods("GET")
	router.HandleFunc("/unit/{kodeUnit}", unitController.FindById).Methods("GET")
	router.HandleFunc("/unit", unitController.Save).Methods("POST")
	router.HandleFunc("/unit/{kodeUnit}", unitController.Update).Methods("PUT")
	router.HandleFunc("/unit/{kodeUnit}", unitController.Delete).Methods("DELETE")
}
