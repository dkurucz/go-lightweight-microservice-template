package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/dkurucz/go-lightweight-microservice-template/src/config"
	"github.com/dkurucz/go-lightweight-microservice-template/src/models"
	"github.com/dkurucz/go-lightweight-microservice-template/src/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"log"
	"net/http"
)

var db *gorm.DB

func main() {
	fmt.Println("Starting GLMT server")
	
	fmt.Println("Running GLMT Auto Migration")
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&models.Account{}, &models.Session{})

	fmt.Println("Go Microservice Template Gorilla server running on http://localhost:9010")
	r := mux.NewRouter()
	routes.RegisterBackendRoutes(r)
	http.Handle("/api/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
}
