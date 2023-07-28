package carplate_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/driver/carplate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCarPlates(t *testing.T) {

	asserts := assert.New(t)

	t.Run("Deve testar placas válidas", func(t *testing.T) {
		plates := []string{
			"ABC1235",
			"AAA1111",
		}

		for _, c := range plates {
			_, err := carplate.New(c)
			asserts.Nil(err, "Car plate %s deveria ser válido", c)
		}
	})

	t.Run("Deve testar placas invalidas", func(t *testing.T) {
		plates := []string{
			"A1235",
			"ABC123",
			"ABCA123",
			"ABC11123",
			"",
		}

		expectedError := "invalid car plate"
		for _, c := range plates {
			_, err := carplate.New(c)
			asserts.EqualError(err, expectedError, "A mensagem deveria ser {}", c)
		}
	})
}
