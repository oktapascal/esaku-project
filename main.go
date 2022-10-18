package main

import (
	controllers2 "esaku-project/app/auths/controllers"
	repository2 "esaku-project/app/auths/repository"
	services2 "esaku-project/app/auths/services"
	"esaku-project/app/settings/controllers"
	"esaku-project/app/settings/repository"
	"esaku-project/app/settings/services"
	"esaku-project/configs"
	"esaku-project/configs/databases"
	"esaku-project/configs/routes"
	"esaku-project/configs/storages"
	"esaku-project/helpers"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"net/http"
)

func main() {
	appConfig := configs.New(".env.dev")
	validate := validator.New()
	sqlServer := databases.NewSqlServer(appConfig)
	awsS3 := storages.NewSessionAws(appConfig)

	kelompokMenuRepository := repository.NewKelompokMenuRepositoryImpl()
	kelompokMenuService := services.NewKelompokMenuServiceImpl(kelompokMenuRepository, sqlServer, validate)
	kelompokMenuController := controllers.NewKelompokMenuControllerImpl(kelompokMenuService)

	formMenuRepository := repository.NewFormRepositoryImpl()
	formService := services.NewFormServiceImpl(formMenuRepository, sqlServer, validate)
	formController := controllers.NewFormControllerImpl(formService)

	unitRepository := repository.NewUnitRepositoryImpl()
	unitService := services.NewUnitServiceImpl(unitRepository, sqlServer, validate)
	unitController := controllers.NewUnitControllerImpl(unitService)

	karyawanRepository := repository.NewKaryawanRepositoryImpl()
	karyawanService := services.NewKaryawanServiceImpl(karyawanRepository, sqlServer, validate, awsS3)
	karyawanController := controllers.NewKaryawanControllerImpl(karyawanService)

	aksesRepository := repository.NewHakAksesRepositoryImpl()
	aksesService := services.NewHakAksesServiceImpl(aksesRepository, sqlServer, validate)
	aksesController := controllers.NewHakAksesControllerImpl(aksesService)

	loginRepository := repository2.NewLoginRepositoryImpl()
	loginService := services2.NewLoginServiceImpl(loginRepository, sqlServer, validate)
	loginController := controllers2.NewLoginControllerImpl(loginService)

	router := routes.NewRouter(
		kelompokMenuController,
		formController,
		unitController,
		karyawanController,
		aksesController,
		loginController,
	)

	// CORS
	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: handlers.CORS(credentials, methods, origins)(router),
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
