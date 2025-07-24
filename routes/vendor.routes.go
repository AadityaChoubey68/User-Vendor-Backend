package routes

import (
	"github.com/AadityaChoubey68/user-vendor-dashboard/handler"
	"github.com/gorilla/mux"
)

func RegisterVendorRoutes(router *mux.Router, handler *handler.Vendorhandler) {
	router.HandleFunc("/vendor/signup", handler.VendorSignUpHandler)
	router.HandleFunc("/vendor/login", handler.VendorLogin)
}
