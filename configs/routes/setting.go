package routes

import (
	"esaku-project/app/settings/controllers"
	"esaku-project/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeSettingsRoutes(
	router *mux.Router,
	kelompokMenuController controllers.KelompokMenuController,
	formController controllers.FormController,
	unitController controllers.UnitController,
	karyawanController controllers.KaryawanController,
	aksesController controllers.HakAksesController,
	menuController controllers.MenuController,
	userController controllers.UserController,
) {
	router.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		json := map[string]string{
			"status": "OK",
		}

		writer.WriteHeader(http.StatusOK)
		helpers.WriteToResponseBodyJson(writer, json)
	})
	// Kelompok Menu
	router.HandleFunc("/v1/kelompok-menu", kelompokMenuController.FindAll).Methods("GET")
	router.HandleFunc("/v1/kelompok-menu/{kodeKlp}", kelompokMenuController.FindById).Methods("GET")
	router.HandleFunc("/v1/kelompok-menu", kelompokMenuController.Save).Methods("POST")
	router.HandleFunc("/v1/kelompok-menu/{kodeKlp}", kelompokMenuController.Update).Methods("PUT")
	router.HandleFunc("/v1/kelompok-menu/{kodeKlp}", kelompokMenuController.Delete).Methods("DELETE")

	// Form
	router.HandleFunc("/v1/form", formController.FindAll).Methods("GET")
	router.HandleFunc("/v1/form/{kodeForm}", formController.FindById).Methods("GET")
	router.HandleFunc("/v1/form", formController.Save).Methods("POST")
	router.HandleFunc("/v1/form/{kodeForm}", formController.Update).Methods("PUT")
	router.HandleFunc("/v1/form/{kodeForm}", formController.Delete).Methods("DELETE")

	// Unit
	router.HandleFunc("/v1/unit", unitController.FindAll).Methods("GET")
	router.HandleFunc("/v1/unit/{kodeUnit}", unitController.FindById).Methods("GET")
	router.HandleFunc("/v1/unit", unitController.Save).Methods("POST")
	router.HandleFunc("/v1/unit/{kodeUnit}", unitController.Update).Methods("PUT")
	router.HandleFunc("/v1/unit/{kodeUnit}", unitController.Delete).Methods("DELETE")

	//	Karyawan
	router.HandleFunc("/v1/karyawan", karyawanController.FindAll).Methods("GET")
	router.HandleFunc("v1/karyawan/{nik}", karyawanController.FindById).Methods("GET")
	router.HandleFunc("/v1/karyawan", karyawanController.Save).Methods("POST")
	router.HandleFunc("/v1/karyawan-upload", karyawanController.UploadImage).Methods("POST")
	router.HandleFunc("/v1/karyawan/{nik}", karyawanController.Update).Methods("PUT")
	router.HandleFunc("/v1/karyawan/{nik}", karyawanController.Delete).Methods("DELETE")

	//	Hak Akses
	router.HandleFunc("/v1/hak-akses", aksesController.FindAll).Methods("GET")
	router.HandleFunc("/v1/hak-akses/{nik}", aksesController.FindById).Methods("GET")
	router.HandleFunc("/v1/hak-akses", aksesController.Save).Methods("POST")
	router.HandleFunc("/v1/hak-akses/{nik}", aksesController.FindById).Methods("PUT")
	router.HandleFunc("/v1/hak-akses/{nik}", aksesController.Delete).Methods("DELETE")

	//	MENU
	router.HandleFunc("/v1/menu/{kodeKlp}", menuController.FindById).Methods("GET")
	router.HandleFunc("/v1/menu", menuController.Save).Methods("POST")

	// USER
	router.HandleFunc("/v1/user/update-data", userController.Update).Methods("POST")
	router.HandleFunc("/v1/user/update-password", userController.UpdatePassword).Methods("POST")
	router.HandleFunc("/v1/user/upload", userController.UploadImage).Methods("POST")
	router.HandleFunc("/v1/user/biodata", userController.FindById).Methods("GET")

}
