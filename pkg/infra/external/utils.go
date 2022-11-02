package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	application "tempest-gateway-service/pkg/application/entities"
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
