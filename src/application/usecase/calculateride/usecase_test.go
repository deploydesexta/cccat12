package calculateride_test

import (
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"testing"
)

type (
	Input struct {
		Segments *[]struct {
			Distance float64
			Date     string
		}
	}
)

func TestCalculateRide_DuringDay(t *testing.T) {
	expectedPrice := 21.0

	input := Input{
		Segments: &[]struct {
			Distance float64
			Date     string
		}{
			{Distance: 10, Date: "2021-03-01T10:00:00"},
		},
	}

	useCase := calculateride.New()
	r, err := useCase.Execute(calculateride.Input(input))
	if err != nil {
		t.Errorf("Error executing use case: %v", err)
		return
	}

	// Verify price value
	if r.Price != expectedPrice {
		t.Errorf("Expected price %.2f, got %.2f", expectedPrice, r.Price)
	}
}

func TestCalculateRide_InvalidDistance(t *testing.T) {
	expectedErrorMessage := "invalid distance"

	input := Input{
		Segments: &[]struct {
			Distance float64
			Date     string
		}{
			{Distance: -10, Date: "2021-03-01T10:00:00"},
		},
	}

	useCase := calculateride.New()
	_, err := useCase.Execute(calculateride.Input(input))

	if err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMessage, err.Error())
		return
	}
}
