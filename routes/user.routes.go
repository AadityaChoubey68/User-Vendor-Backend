package routes

import (
	"github.com/AadityaChoubey68/user-vendor-dashboard/handler"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, handler *handler.Userhandler) {
	router.HandleFunc("/user/signup", handler.UserSignUpHandler).Methods("POST")

	router.HandleFunc("/user/login", handler.UserLogin).Methods("POST")
}
