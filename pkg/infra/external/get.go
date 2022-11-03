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

	resp, err := http.Get(route)
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readBody(*resp)
}
