package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)


func InvokeAlertNotificationAPI(message types.UpdateActiveStatus ) error {
	// Prepare the API payload
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal API payload: %v", err)
	}

	// Make the HTTP POST request
	resp, err := http.Post("http://localhost:8080/api/alert/alert-notification", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to call alert notification API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API call failed with status %d: %s", resp.StatusCode, string(body))
	}

	log.Printf("Alert notification sent successfully for alert ID: %s", message.ID)
	return nil
}