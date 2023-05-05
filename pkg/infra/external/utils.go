package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
)

const (
	headerAuth             = "Authorization"
	headerAccept           = "Accept"
	headerContentType      = "Content-Type"
	headerTransferEncoding = "Transfer-Encoding"
)

func readBody(resp http.Response) (*application.Response, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body, err %v", err)
	}

	applicationResponse := application.Response{}
	err = json.Unmarshal(body, &applicationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling body, err %v", err)
	}

	return &applicationResponse, nil
}

func constructRoute(req entities.Request) (string, error) {

	if req.Host == "" {
		return "", fmt.Errorf("error, no host specified")
	}

	route := fmt.Sprintf("%s/%s", req.Host, req.Route)
	if req.Port != 0 {
		route = fmt.Sprintf("%s:%v/%s", req.Host, req.Port, req.Route)
	}

	return route, nil
}

func NewRequest(req entities.Request) (*application.Response, error) {

	switch req.Method {
	case http.MethodGet:
		return Get(req)
	case http.MethodPost:
		return Post(req)
	}

	return nil, fmt.Errorf("error, method not supported yet")
}
