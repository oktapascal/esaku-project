package routes

import (
	controllers2 "esaku-project/app/auths/controllers"
	"esaku-project/app/settings/controllers"
	"esaku-project/middlewares"
	"github.com/gorilla/mux"
)

func NewRouter(
	middlewareAuth *middlewares.MiddlewareAuthImpl,
	kelompokMenuController controllers.KelompokMenuController,
	formController controllers.FormController,
	unitController controllers.UnitController,
	karyawanController controllers.KaryawanController,
	aksesController controllers.HakAksesController,
	loginController controllers2.LoginController,
	menuController controllers.MenuController,
	userController controllers.UserController,
) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.Use(middlewares.CustomError)

	auth := router.PathPrefix("/api/esaku-auth").Subrouter()
	InitializeAuthRoutes(
		auth,
		loginController,
	)

	setting := router.PathPrefix("/api/esaku-setting").Subrouter()
	setting.Use(middlewareAuth.MiddlewareCookie)
	setting.Use(middlewareAuth.MiddlewareRefreshToken)
	setting.Use(middlewareAuth.MiddlewareBearerToken)

	InitializeSettingsRoutes(
		setting,
		kelompokMenuController,
		formController,
		unitController,
		karyawanController,
		aksesController,
		menuController,
		userController,
	)

	master := router.PathPrefix("/api/esaku-master").Subrouter()
	InitializeMasterRoutes(master)

	transaction := router.PathPrefix("/api/esaku-trans").Subrouter()
	InitializeTransactionRoutes(transaction)

	report := router.PathPrefix("/api/esaku-report").Subrouter()
	InitializeReportRoutes(report)

	filter := router.PathPrefix("/api/esaku-filter").Subrouter()
	InitializeFilterRoutes(
		filter,
		unitController,
		karyawanController,
		kelompokMenuController,
	)

	return router
}
