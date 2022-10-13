package routes

import (
	"esaku-project/app/settings/controllers"
	"esaku-project/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter(
	kelompokMenuController controllers.KelompokMenuController,
	formController controllers.FormController,
	unitController controllers.UnitController,
	karyawanController controllers.KaryawanController,
	aksesController controllers.HakAksesController,
) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Use(middlewares.CustomError)

	auth := router.PathPrefix("/api/esaku-auth").Subrouter()
	InitializeAuthRoutes(auth)

	setting := router.PathPrefix("/api/esaku-setting").Subrouter()
	setting.Use(middlewares.MiddlewareAuth)
	InitializeSettingsRoutes(
		setting,
		kelompokMenuController,
		formController,
		unitController,
		karyawanController,
		aksesController,
	)

	master := router.PathPrefix("/api/esaku-master").Subrouter()
	InitializeMasterRoutes(master)

	transaction := router.PathPrefix("/api/esaku-trans").Subrouter()
	InitializeTransactionRoutes(transaction)

	report := router.PathPrefix("/api/esaku-report").Subrouter()
	InitializeReportRoutes(report)

	return router
}
