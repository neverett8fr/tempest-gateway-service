package external

import (
	"fmt"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
)

func Get(req entities.Request) (*application.Response, error) {

	route, err := constructRoute(req)
	if err != nil {
		return nil, fmt.Errorf("error constructing route, err %v", err)
	}

	request, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return nil, fmt.Errorf("error building request, err %v", err)
	}

	request.Header.Add(headerAuth, req.Auth)
	request.Header.Add(headerAccept, req.Accept)
	request.Header.Add(headerContentType, req.ContentType)
	request.Header.Add(headerTransferEncoding, req.Transfer)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readBody(*resp)
}
