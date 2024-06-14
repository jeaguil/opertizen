package config

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

var Cfg *Configuration

func init() {
	goPath := os.Getenv("GOPATH")
	absolutePath := path.Join(goPath, "src/opertizen/")
	var err error
	Cfg, err = LoadConfig(absolutePath + "/configs/tvconf.yaml")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
}

type Configuration struct {
	Properties struct {
		OperatingSystem        string `yaml:"operating_system"`
		Model                  string `yaml:"model"`
		Description            string `yaml:"description"`
		SmartThingsDeviceID    string `yaml:"smartthings_devide_id"`
		SmartThingsAccessToken string `yaml:"smartthings_access_token"`
	} `yaml:"properties"`

	Network struct {
		IPAddress string `yaml:"ip_address"`
		Port      int    `yaml:"port"`
	} `yaml:"network"`
}

func LoadConfig(filepath string) (*Configuration, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
