package runflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"opertizen/internal/command"
	"opertizen/internal/config"
	"opertizen/internal/network"

	"gopkg.in/yaml.v2"
)

type Step struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

type RunflowDetails struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []Step `yaml:"steps"`
}

type Runflow struct {
	Runflow RunflowDetails `yaml:"runflow"`
}

// LoadRunflow loads a runflow file and structures it for execution
func LoadRunflow(file string) (Runflow, error) {
	var runflow Runflow
	data, err := os.ReadFile(file)
	if err != nil {
		return runflow, err
	}
	err = yaml.Unmarshal(data, &runflow)
	return runflow, err
}

func (r *Runflow) ProcessRunflow() {
	reqs, err := r.parseRunflowFile()
	if err != nil {
		log.Fatalf("Error during runflow parsing: %v", err)
	}
	url := fmt.Sprintf("https://api.smartthings.com/v1/devices/%s/commands", config.Cfg.Properties.SmartThingsDeviceID)
	for _, req := range reqs {

		jsonData, err := json.Marshal(req)
		if err != nil {
			log.Fatal(err)
		}
		if err := network.CallSmartThingsAPI("POST", url, bytes.NewBuffer(jsonData)); err != nil {
			log.Fatalf("Failed to call SmartThings API: %v", err)
		}

		// wait 3 seconds before sending another request
		time.Sleep(3 * time.Second)
	}
}

func (r *Runflow) parseRunflowFile() ([]command.SmartThingsCommandRequest, error) {
	var requests []command.SmartThingsCommandRequest
	for _, step := range r.Runflow.Steps {
		cap, cmd, args, err := command.CheckCommandFromRunflow(step.Command)
		if err != nil {
			log.Fatal(err)
		}
		smartThingsCommand := command.SmartThingsCommand{
			Capability: cap,
			Command:    cmd,
			Arguments:  args,
		}
		request := command.SmartThingsCommandRequest{
			Commands: []command.SmartThingsCommand{smartThingsCommand},
		}
		requests = append(requests, request)
	}
	log.Println("Successfully parsed runflow file.")
	return requests, nil
}
