package routes

import (
	"github.com/dkurucz/go-lightweight-microservice-template/src/controllers"
	"github.com/gorilla/mux"
)

var RegisterBackendRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/account/create", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/account/login", contro llers.Login).Methods("POST")
	router.HandleFunc("/api/account/{id}", controllers.DeleteAccount).Methods("DELETE")
}
