package auth

import (
	"fmt"
	"log"
	"tempest-gateway-service/pkg/config"
	"tempest-gateway-service/pkg/infra/entities"
)

// call "auth" service
// get token from auth
// use auth to check permissions for user & resources

func CheckValidRequest(conf config.Config, service string, request entities.Request) error {
	if !contains(conf.Endpoints[service].AllowedMethods, request.Method) {
		return fmt.Errorf("error method %v not allowed", request.Method)
	}

	if conf.Endpoints[service].Host == "" {
		return fmt.Errorf("endpoint/service %v not found", service)
	}

	if conf.Endpoints[service].Auth {
		log.Printf("calling auth service")
		// need to create auth service
	}

	// call service - return token to use for post requests, etc. if "basic", generate new token for first time use, if jwt, use that - and make sure it's cached on auth

	return nil
}
