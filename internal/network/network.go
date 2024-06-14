package network

import (
	"io"
	"log"
	"net/http"

	"opertizen/internal/config"
)

// CallSmartThingsAPI performs a request to the SmartThings API
func CallSmartThingsAPI(method string, url string, body io.Reader) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+config.Cfg.Properties.SmartThingsAccessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to send command: %s", responseBody)
	}

	log.Println("Command sent successfully:", string(responseBody))
	return nil
}
