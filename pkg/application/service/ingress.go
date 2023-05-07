package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	acceptIn := r.Header.Get(headerAccept)

	// make this a function to handle different formats in future, no if else stuff
	var body interface{}
	if contentType == contentTypeJSON {
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Printf("err %v", err)
		}
	} else {
		bodyByte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("err %v", err)
		}
		body = bodyByte
	}

	log.Printf("the body is %v", body)

	log.Printf("request received for service %v, route %v", service, route)

	request := entities.Request{
		Host:        conf.Endpoints[service].Host,
		Port:        conf.Endpoints[service].Port,
		Route:       route,
		ContentType: contentType,
		Accept:      acceptIn,
		Method:      r.Method,
		Auth:        authIn,
		Body:        body,
	}

	log.Printf("requestTy: %v", request.ContentType)
	log.Printf("requestAc: %v", request.Accept)
	log.Printf("requestBo: %v", request.Body)

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
