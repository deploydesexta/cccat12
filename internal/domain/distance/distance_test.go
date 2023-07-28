package distance_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/distance"
	"github.com/deploydesexta/cccat12/internal/domain/distance/coord"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceCalculator(t *testing.T) {
	assertions := assert.New(t)

	t.Run("Deve calcular a distância entre duas coordenadas", func(t *testing.T) {
		expectedPrice := 10.0

		from := coord.New(-27.584905257808835, -48.545022195325124)
		to := coord.New(-27.496887588317275, -48.522234807851476)

		actualPrice := distance.Calculate(from, to)

		assertions.Equal(expectedPrice, actualPrice)
	})
}
