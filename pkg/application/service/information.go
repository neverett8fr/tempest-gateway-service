package service

import (
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"

	"github.com/gorilla/mux"
)

func newInformationRoutes(r *mux.Router) {
	r.HandleFunc("/routes", getRoutes).Methods("GET")
}

func getRoutes(w http.ResponseWriter, r *http.Request) {

	endpoints := conf.Endpoints
	body := application.NewResponse(endpoints)

	writeReponse(w, r, body)
}
