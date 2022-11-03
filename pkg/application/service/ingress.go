package service

import (
	"fmt"
	"log"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/config"
	"tempest-gateway-service/pkg/infra/entities"
	"tempest-gateway-service/pkg/infra/external"

	"github.com/gorilla/mux"
)

func newIngressRoutes(r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/{%s}/{%s:.*}", service, route), getRequest).Methods(http.MethodGet)
}

func getRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	service := params[service]
	route := params[route]

	log.Printf("request for service %v, route %v", service, route)
	if conf.Endpoints[service] == (config.Endpoints{}) {
		body := application.NewResponse(nil, fmt.Errorf("endpoint/service %v not found", service))
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// check permission

	// call service
	request := entities.GetRequest{
		Host:        conf.Endpoints[service].Host,
		Port:        conf.Endpoints[service].Port,
		Route:       route,
		ContentType: contentTypeJSON,
	}

	res, err := external.Get(request)
	if err != nil {
		body := application.NewResponse(nil, err)
		w.WriteHeader(http.StatusBadRequest)
		writeReponse(w, body)
		return
	}

	// respond
	body := application.NewResponse(res, nil)
	writeReponse(w, body)
}
