package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	route = "config/config.yaml"
)

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Endpoints struct {
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	AllowedMethods []string `yaml:"methods"`
	AllowedRoutes  []string `yaml:"allowed-routes"`
	LimitRoutes    bool     `yaml:"limit-routes"`
	Auth           bool     `yaml:"auth"`
}

type Auth struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Service   Service              `yaml:"gateway-service-config"`
	Endpoints map[string]Endpoints `yaml:"endpoints"`
	Auth      Auth                 `yaml:"auth"`
}

func Initialise() (*Config, error) {

	conf := Config{}

	yamlFile, err := ioutil.ReadFile(route)
	if err != nil {
		return &Config{}, fmt.Errorf("issue finding config yaml, err %v", err)
	}
	yamlFile = []byte(os.ExpandEnv(string(yamlFile)))

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return &Config{}, fmt.Errorf("issue unmarshalling config yaml, err %v", err)
	}

	return &conf, nil
}
