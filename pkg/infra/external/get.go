package external

import (
	"fmt"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
)

func Get(req entities.GetRequest) (*application.Response, error) {

	route := fmt.Sprintf("%s/%s", req.Host, req.Route)
	if req.Port != 0 {
		route = fmt.Sprintf("%s:%v/%s", req.Host, req.Port, req.Route)
	}

	resp, err := http.Get(route)
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readBody(*resp)
}
