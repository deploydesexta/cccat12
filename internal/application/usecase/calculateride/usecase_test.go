package calculateride_test

import (
	"github.com/deploydesexta/cccat12/internal/application/usecase/calculateride"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	Input struct {
		Positions *[]struct {
			Lat  float64
			Long float64
			Date string
		}
	}
)

func TestCalculateRideUseCase(t *testing.T) {
	assertions := assert.New(t)

	t.Run("Deve fazer o cálculo do preço de uma corrida durante o dia", func(t *testing.T) {
		expectedPrice := 21.0

		input := Input{
			Positions: &[]struct {
				Lat  float64
				Long float64
				Date string
			}{
				{Lat: -27.584905257808835, Long: -48.545022195325124, Date: "2021-03-01T10:00:00"},
				{Lat: -27.496887588317275, Long: -48.522234807851476, Date: "2021-03-01T10:01:00"},
			},
		}

		useCase := calculateride.New()
		r, err := useCase.Execute(calculateride.Input(input))
		if err != nil {
			t.Errorf("Error executing use case: %v", err)
			return
		}

		assertions.Equal(expectedPrice, r.Price)
	})
}
