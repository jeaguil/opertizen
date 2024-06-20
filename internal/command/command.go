package command

import (
	"errors"
	"log"
	"regexp"
	"slices"
	"strings"
)

// Capabilities is the current list of supported capabilities from SmartThings API
var capabilities = map[string][]string{
	"switch":                  {"on", "off"},
	"audioVolume":             {"volumeDown", "volumnUp"},
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
	Capability string   `json:"capability"`
	Command    string   `json:"command"`
	Arguments  []string `json:"arguments"`
}

type SmartThingsCommandRequest struct {
	Commands []SmartThingsCommand `json:"commands"`
}

// CheckCommandFromRunflow checks if a command is supported by Opertizen.
// If a capability or command is not found, returns error
// indicating a command cannot be processed.
func CheckCommandFromRunflow(commandStr string) (string, string, []string, error) {
	var cap string
	var cmd string
	var args []string
	splitCmd := strings.Split(commandStr, ";")
	if len(splitCmd) == 2 {
		capa := splitCmd[0]
		comm := splitCmd[1]
		if _, ok := capabilities[capa]; ok {
			cap = capa

			openParenIndex := strings.Index(comm, "(")
			if openParenIndex > 0 {
				args = parseArguments(cap, comm)
				comm = comm[:openParenIndex]
			}

			if slices.Contains(capabilities[cap], comm) {
				cmd = comm
			}
		} else {
			return "", "", nil, errors.New("Cannot find supported command in runflow: " + commandStr)
		}
	} else {
		return "", "", nil, errors.New("Invalid format or unsupported command in runflow file:\ncommand:<capability>;<command>(<arguments>): " + commandStr)
	}
	log.Println(cap, cmd, args)
	return cap, cmd, args, nil
}

func parseArguments(capability, command string) []string {
	var args []string
	var err error
	if strings.Compare(capability, "samsungvd.remoteControl") == 0 {
		args, err = constructSamsungvdRemoteControlRequest(command)
		if err != nil {
			log.Fatalf("Error during parsing arguments for a command: %v", command)
		}
	}
	return args
}

func constructSamsungvdRemoteControlRequest(command string) ([]string, error) {
	re := regexp.MustCompile(`\w+\(([^)]+)\)`)
	matches := re.FindStringSubmatch(command)
	if len(matches) < 2 {
		return nil, errors.New("Invalid arguments for command: " + command)
	}

	args := strings.Split(matches[1], ",")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	return args, nil
}
