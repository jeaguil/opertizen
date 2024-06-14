package network

import (
	"fmt"
	"opertizen/internal/config"
	"testing"
)

// TestSmartThingsAPIConnectivity calls getDevice.
// Get a device's given description.
func TestSmartThingsAPIConnectivity(t *testing.T) {
	url := fmt.Sprintf("https://api.smartthings.com/v1/devices/%s", config.Cfg.Properties.SmartThingsDeviceID)
	if err := CallSmartThingsAPI("GET", url, nil); err != nil {
		t.Logf("Failed to call SmartThings API: %v", err)
	}
}
