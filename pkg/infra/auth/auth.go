package auth

import (
	"fmt"
	"log"
	"net/http"
	"tempest-gateway-service/pkg/config"
	"tempest-gateway-service/pkg/infra/entities"
	"tempest-gateway-service/pkg/infra/external"
)

// call "auth" service
// get token from auth
// use auth to check permissions for user & resources

func CheckValidRequest(conf config.Config, service string, request entities.Request) error {
	if conf.Endpoints[service].Host == "" {
		return fmt.Errorf("endpoint/service %v not found", service)
	}

	if !contains(conf.Endpoints[service].AllowedMethods, request.Method) {
		return fmt.Errorf("error method %v not allowed", request.Method)
	}

	if conf.Endpoints[service].LimitRoutes && !contains(conf.Endpoints[service].AllowedRoutes, request.Route) {
		return fmt.Errorf("error route %v not allowed", request.Route)
	}

	if conf.Endpoints[service].Auth {
		log.Printf("calling auth service")

		res, err := external.NewRequest(entities.Request{
			Host:   conf.Auth.Host,
			Port:   conf.Auth.Port,
			Route:  "token",
			Method: http.MethodGet,
			Auth:   request.Auth,
		})
		if err != nil {
			return fmt.Errorf("error calling auth service, err %v", err)
		}
		if res.Data != "token is valid" {
			return fmt.Errorf("error auth service returned error, err %v", res.Errors)
		}

	}

	return nil
}
