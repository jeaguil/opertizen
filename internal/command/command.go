package command

import (
	"errors"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// Capabilities is the current list of supported capabilities from SmartThings API
var capabilities = map[string][]string{
	"switch":                  {"on", "off"},
	"audioVolume":             {"volumeDown", "volumeUp", "setVolume"},
	"audioMute":               {"mute", "unmute"},
	"samsungvd.remoteControl": {"send"},
}

// SamsungvdRemoteControl is the cabilitity that only has the send command.
// Send commands takes two arguments: keyValue, and keyState.
// keyValue: UP, DOWN, LEFT, RIGHT, OK, BACK, MENU, HOME
// keyState: PRESSED, RELEASED, PRESSED_AND_RELEASED
var SamsungvdRemoteControl = map[string][]string{
	"keyValue": {"UP", "DOWN", "LEFT", "RIGHT", "OK", "BACK", "MENU", "HOME"},
	"keyState": {"PRESSED", "RELEASED", "PRESSED_AND_RELEASED"},
}

// SmartThingsCommand defines the structure for a command through the SmartThings API
type SmartThingsCommand struct {
	Capability string        `json:"capability"`
	Command    string        `json:"command"`
	Arguments  []interface{} `json:"arguments"`
}

type SmartThingsCommandRequest struct {
	Commands []SmartThingsCommand `json:"commands"`
}

// CheckCommandFromRunflow checks if a command is supported by Opertizen.
// If a capability or command is not found, returns error
// indicating a command cannot be processed.
func CheckCommandFromRunflow(commandStr string) bool {
	splitCmd := strings.Split(commandStr, ";")
	if len(splitCmd) == 2 {
		cap := splitCmd[0]
		cmd := splitCmd[1]
		if _, ok := capabilities[cap]; ok {
			openParenIndex := strings.Index(cmd, "(")
			if openParenIndex > 0 {
				cmd = cmd[:openParenIndex]
			}
			if !slices.Contains(capabilities[cap], cmd) {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

// ConstructSmartThingsRequest constructs a command suitable for the SmartThings API
func ConstructSmartThingsRequest(commandStr string) SmartThingsCommand {
	var request SmartThingsCommand
	splitCmd := strings.Split(commandStr, ";")
	capadability := splitCmd[0]
	command := splitCmd[1]
	openParenIndex := strings.Index(command, "(")
	if openParenIndex > 0 {
		var err error
		request.Arguments, err = parseArguments(command)
		if err != nil {
			log.Fatal(err)
		}
		command = command[:openParenIndex]
	}
	request.Capability = capadability
	request.Command = command
	return request
}

func parseArguments(command string) ([]interface{}, error) {
	re := regexp.MustCompile(`\w+\(([^)]+)\)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) < 2 {
		return nil, errors.New("Invalid arguments for command: " + command)
	}

	args := strings.Split(matches[1], ",")
	var processedArgs []interface{}
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if intArg, err := strconv.Atoi(arg); err == nil {
			processedArgs = append(processedArgs, intArg)
		} else {
			processedArgs = append(processedArgs, arg)
		}
	}

	return processedArgs, nil
}
