package main_test

import (
	"encoding/json"
	"fmt"
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

func TestCreatePassenger(t *testing.T) {
	input := `{
		"name": "Ciclano",
    	"document": "674.338.630-87",
    	"email": "ciclano@gmail.com"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/passengers", "application/json", strings.NewReader(input))
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
		PassengerId *string `json:"PassengerId"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	// Verify
	if response.PassengerId == nil {
		t.Errorf("Expected id, got nil")
	}
}

func TestCreatePassenger_InvalidDocument(t *testing.T) {
	input := `{
		"name": "Ciclano",
    	"document": "11133863087",
    	"email": "ciclano@gmail.com"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/passengers", "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}

	// Parse response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Verify
	expectedError := "invalid cpf"
	if string(body) != expectedError {
		t.Errorf("Expected %s, got %s", expectedError, string(body))
	}
}

func TestGetPassenger(t *testing.T) {
	input := `{
		"name": "Ciclano",
    	"document": "674.338.630-87",
    	"email": "ciclano@gmail.com"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/passengers", "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Parse response body
	var createResponse struct {
		PassengerId string `json:"passengerId"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Simulate Get passenger request
	resp, err = http.Get(fmt.Sprintf("http://localhost:3000/passengers/%s", createResponse.PassengerId))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Parse response body
	var response struct {
		Id       string `json:"id"`
		Document string `json:"document"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	// Verify response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	expectedName := "Ciclano"
	if response.Name != expectedName {
		t.Errorf("Expected name '%s', got '%s'", expectedName, response.Name)
	}

	expectedDoc := "674.338.630-87"
	if response.Document != expectedDoc {
		t.Errorf("Expected document '%s', got '%s'", expectedDoc, response.Document)
	}

	expectedEmail := "ciclano@gmail.com"
	if response.Email != expectedEmail {
		t.Errorf("Expected name '%s', got '%s'", expectedEmail, response.Email)
	}
}

func TestCreateDriver(t *testing.T) {
	input := `{
		"name": "Fulano",
		"document": "776.794.210-48",
		"email": "fulano@gmail.com",
		"carplate": "ABC1234"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/drivers", "application/json", strings.NewReader(input))
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
		DriverId *string `json:"driverId"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Verify
	if response.DriverId == nil {
		t.Errorf("Expected id, got nil")
	}
}

func TestCreateDriver_InvalidDocument(t *testing.T) {
	input := `{
		"name": "Fulano",
		"document": "11179421048",
		"email": "fulano@gmail.com",
		"carplate": "ABC1234"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/drivers", "application/json", strings.NewReader(input))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("Expected status %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}

	// Parse response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	// Verify
	expectedError := "invalid cpf"
	if string(body) != expectedError {
		t.Errorf("Expected %s, got %s", expectedError, string(body))
	}
}

func TestGetDriver(t *testing.T) {
	input := `{
		"name": "Fulano",
		"document": "776.794.210-48",
		"email": "fulano@gmail.com",
		"carplate": "ABC1234"
	}`

	// Simulate server handling the request
	resp, err := http.Post("http://localhost:3000/drivers", "application/json", strings.NewReader(input))
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
	var createResponse struct {
		DriverId string `json:"driverId"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &createResponse)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Simulate Get passenger request
	resp, err = http.Get(fmt.Sprintf("http://localhost:3000/drivers/%s", createResponse.DriverId))
	if err != nil {
		t.Errorf("Http connection has failed")
		return
	}
	defer resp.Body.Close()

	// Parse response body
	var response struct {
		Id       string `json:"id"`
		Document string `json:"document"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		CarPlate string `json:"carplate"`
	}

	// Verify response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	expectedName := "Fulano"
	if response.Name != expectedName {
		t.Errorf("Expected name '%s', got '%s'", expectedName, response.Name)
	}

	expectedDoc := "776.794.210-48"
	if response.Document != expectedDoc {
		t.Errorf("Expected document '%s', got '%s'", expectedDoc, response.Document)
	}

	expectedEmail := "fulano@gmail.com"
	if response.Email != expectedEmail {
		t.Errorf("Expected name '%s', got '%s'", expectedEmail, response.Email)
	}

	expectedPlate := "ABC1234"
	if response.CarPlate != expectedPlate {
		t.Errorf("Expected carplate '%s', got '%s'", expectedPlate, response.CarPlate)
	}
}
