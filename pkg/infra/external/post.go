package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
)

func Post(req entities.Request) (*application.Response, error) {

	route, err := constructRoute(req)
	if err != nil {
		return nil, fmt.Errorf("error constructing route, err %v", err)
	}

	body := bytes.Buffer{}
	err = json.NewEncoder(&body).Encode(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error encoding body, err %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, route, &body)
	if err != nil {
		return nil, fmt.Errorf("error building request, err %v", err)
	}

	request.Header.Add(headerAuth, req.Auth)
	request.Header.Add(headerAccept, req.ContentType)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readBody(*resp)
}
