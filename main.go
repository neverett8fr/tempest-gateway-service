package main

import (
	"log"
	"tempest-gateway-service/cmd"
	application "tempest-gateway-service/pkg/application/service"
	"tempest-gateway-service/pkg/config"

	"github.com/gorilla/mux"
)

// Route declaration
func getRoutes() *mux.Router {
	r := mux.NewRouter()
	application.NewIngressRoutes(r)

	return r
}

// Initiate web server
func main() {
	conf, err := config.Initialise()
	if err != nil {
		log.Fatalf("error initialising config, err %v", err)
		return
	}
	log.Println("config initialised")

	router := getRoutes()
	log.Println("API routes retrieved")

	err = cmd.StartServer(&conf.Service, router)
	if err != nil {
		log.Fatalf("error starting server, %v", err)
		return
	}
	log.Println("server started")

}
