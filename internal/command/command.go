package command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"opertizen/internal/network"
)

// Capabilities is the current list of supported capabilities from SmartThings API
var Capabilities = map[string][]string{
	"switch":                  {"on", "off"},
	"audioVolume":             {"volumeDown", "volumnUp"},
	"audioMute":               {"mute", "unmute"},
	"samsungvd.remoteControl": {""},
}

// SmartThingsCommand defines the structure for a command through the SmartThings API
type SmartThingsCommand struct {
	Capability string `json:"capability"`
	Command    string `json:"command"`
}

type SmartThingsCommandRequest struct {
	Commands []SmartThingsCommand `json:"commands"`
}

// CallUpdateDevice calls SmartThings updateDevice
// Executes a specified command on a device through the SmartThings API
func CallUpdateDevice(deviceID string, commands SmartThingsCommandRequest) error {
	url := fmt.Sprintf("https://api.smartthings.com/v1/devices/%s/commands", deviceID)
	jsonData, err := json.Marshal(commands)
	if err != nil {
		return err
	}
	if err = network.CallSmartThingsAPI("POST", url, bytes.NewBuffer(jsonData)); err != nil {
		return err
	}
	return nil
}
