package ride_test

import (
	"github.com/deploydesexta/cccat12/src/domain/ride"
	"testing"
	"time"
)

func TestCalculateRide_DuringDay(t *testing.T) {

	r := ride.New()
	err := r.AddSegment(10, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 21.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}

func TestCalculateRide_DuringNight(t *testing.T) {

	r := ride.New()
	err := r.AddSegment(10, time.Date(2021, 3, 1, 23, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 39.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}

func TestCalculateRide_SundayDay(t *testing.T) {

	r := ride.New()
	err := r.AddSegment(10, time.Date(2021, 3, 7, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 29.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}

func TestCalculateRide_SundayNight(t *testing.T) {

	r := ride.New()
	err := r.AddSegment(10, time.Date(2021, 3, 7, 23, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 50.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}

func TestCalculateRide_InvalidDistance(t *testing.T) {
	expectedErrorMessage := "invalid distance"

	r := ride.New()
	err := r.AddSegment(-10, time.Date(2023, 3, 1, 10, 0, 0, 0, time.UTC))
	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error '%s', got '%v'", expectedErrorMessage, err)
	}
}

func TestCalculateRide_DuringDayWithMinimumPrice(t *testing.T) {

	r := ride.New()
	err := r.AddSegment(3, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()

	expectedPrice := 10.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}

func TestCalculateRide_MultipleSegments(t *testing.T) {
	r := ride.New()

	err := r.AddSegment(10, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	err = r.AddSegment(10, time.Date(2021, 3, 1, 10, 0, 0, 0, time.UTC))
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	price := r.Calculate()
	expectedPrice := 42.0
	if price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, price)
	}
}
