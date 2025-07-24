package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AadityaChoubey68/user-vendor-dashboard/driver"
	"github.com/AadityaChoubey68/user-vendor-dashboard/handler"
	"github.com/AadityaChoubey68/user-vendor-dashboard/routes"
	"github.com/AadityaChoubey68/user-vendor-dashboard/service"
	"github.com/AadityaChoubey68/user-vendor-dashboard/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	defer driver.CloseDB()
	driver.ConnectDB()
	db := driver.GetDB()

	userStore := storage.New(db)
	userService := service.NewUserService(userStore)
	userHandler := handler.NewHandler(userService)

	vendorStore := storage.VdNew(db)
	vendorService := service.NewVendorService(vendorStore)
	vendorHandler := handler.NewVendorHandler(vendorService)

	router := mux.NewRouter()
	routes.RegisterUserRoutes(router, userHandler)
	routes.RegisterVendorRoutes(router, vendorHandler)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "✅ Server is working! Hello from the root route.")
	})

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: router,
	}
	fmt.Println("🚀 Server started on Port : localhost:8082")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("❌ Failer to start server")
	}
}
