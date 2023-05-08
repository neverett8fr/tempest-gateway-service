package auth

import (
	"encoding/json"
	"fmt"
	application "tempest-gateway-service/pkg/application/entities"
)

func contains(arr []string, search string) bool {

	for _, val := range arr {
		if val == search {
			return true
		}
	}

	return false
}

func interfaceResponseToApplicationResponse(in interface{}) (*application.Response, error) {

	res := application.Response{}
	bodyBytes, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("error converting response %v", err)
	}

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling respone %v", err)
	}

	return &res, nil

}
