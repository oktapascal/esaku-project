package routes

import (
	controllers2 "esaku-project/app/auths/controllers"
	"github.com/gorilla/mux"
)

func InitializeAuthRoutes(
	router *mux.Router,
	loginController controllers2.LoginController,
) {
	router.HandleFunc("/login", loginController.Login).Methods("POST")
	router.HandleFunc("/logout", loginController.Logout).Methods("POST")
}
