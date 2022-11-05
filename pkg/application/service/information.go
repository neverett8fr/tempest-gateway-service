package service

import (
	"encoding/json"
	"log"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
	"tempest-gateway-service/pkg/infra/external"

	"github.com/gorilla/mux"
)

func newInformationRoutes(r *mux.Router) {
	r.HandleFunc("/routes", getRoutes).Methods(http.MethodGet)
	r.HandleFunc("/token", createToken).Methods(http.MethodPost)
}

func getRoutes(w http.ResponseWriter, r *http.Request) {

	endpoints := conf.Endpoints
	body := application.NewResponse(endpoints)

	writeReponse(w, body)
}

func createToken(w http.ResponseWriter, r *http.Request) {

	var body interface{}
	_ = json.NewDecoder(r.Body).Decode(&body)

	log.Printf("request received for service %v, route %v", service, route)

	request := entities.Request{
		Host:        conf.Auth.Host,
		Port:        conf.Auth.Port,
		Route:       "/token",
		ContentType: contentTypeJSON,
		Method:      r.Method,
		Body:        body,
	}

	// call service
	res, err := external.NewRequest(request)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// respond
	responseBody := application.NewResponse(res, nil)
	writeReponse(w, responseBody)
}
