package main_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestCalculateRide_DuringDay(t *testing.T) {
	input := `{
		"segments": [
			{ "distance": 10, "date": "2021-03-01T10:00:00" }
		]
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/calculate_ride", "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Parse response body
	var response struct {
		Price float64 `json:"price"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Verify price value
	expectedPrice := 21.0
	if response.Price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, response.Price)
	}
}

func TestCalculateRide_InvalidDistance(t *testing.T) {
	input := `{
		"segments": [
			{ "distance": -10, "date": "2021-03-01T10:00:00" }
		]
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/calculate_ride", "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}

	// Verify response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	expectedErrorMessage := "invalid distance"
	if string(body) != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMessage, string(body))
	}
}
