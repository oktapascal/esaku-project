package main

import (
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

	router := routes.NewRouter(
		kelompokMenuController,
		formController,
		unitController,
		karyawanController,
	)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
