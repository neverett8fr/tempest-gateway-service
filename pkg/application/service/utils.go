package service

import (
	"encoding/json"
	"log"
	"net/http"
	"tempest-gateway-service/pkg/config"

	"github.com/gorilla/mux"
)

const (
	service = "service"
	route   = "route"

	headerAuth        = "Authorization"
	headerContentType = "Content-Type"
	headerAccept      = "Accept"

	// contentTypeJSON   = "application/json"
	// contentTypeStream = "application/octet-stream"
	// transferChunked   = "chunked"
)

var (
	conf config.Config
)

func NewServiceRoutes(r *mux.Router, confIn *config.Config) {
	conf = *confIn

	newIngressRoutes(r)
	newInformationRoutes(r)
}

func writeReponse(w http.ResponseWriter, body interface{}) {

	reponseBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error converting reponse to bytes, err %v", err)
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(reponseBody)
	if err != nil {
		log.Printf("error writing response, err %v", err)
	}
}
