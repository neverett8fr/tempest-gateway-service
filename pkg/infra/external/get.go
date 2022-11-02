package external

import (
	"fmt"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
	"tempest-gateway-service/pkg/infra/entities"
)

func Get(req entities.GetRequest) (*application.Response, error) {

	resp, err := http.Get(fmt.Sprintf("%s://%s:%v/%s", req.Protocol, req.Host, req.Port, req.Route))
	if err != nil {
		return nil, fmt.Errorf("error calling service, err %v", err)
	}

	return readBody(*resp)
}
