package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/auth"
	"tempest-gateway-service/pkg/infra/entities"
	"tempest-gateway-service/pkg/infra/external"

	"github.com/gorilla/mux"
)

func newIngressRoutes(r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/{%s}/{%s:.*}", service, route), forward).Methods(http.MethodGet, http.MethodPost)

}

func forward(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	service := params[service]
	route := params[route]

	authIn := r.Header.Get(headerAuth)
	contentType := r.Header.Get(headerContentType)

	var body interface{}
	_ = json.NewDecoder(r.Body).Decode(&body)

	log.Printf("request received for service %v, route %v", service, route)

	request := entities.Request{
		Host:        conf.Endpoints[service].Host,
		Port:        conf.Endpoints[service].Port,
		Route:       route,
		ContentType: contentType,
		Method:      r.Method,
		Auth:        authIn,
		Body:        body,
	}

	// check permission
	err := auth.CheckValidRequest(conf, service, request)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		writeReponse(w, body)
		return
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
