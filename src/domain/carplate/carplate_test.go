package carplate_test

import (
	"github.com/deploydesexta/cccat12/src/domain/carplate"
	"testing"
)

func TestPlacasValidas(t *testing.T) {

	plates := []string{
		"ABC1235",
		"AAA1111",
	}

	for _, c := range plates {
		if _, err := carplate.New(c); err != nil {
			t.Errorf("Car plate %s deveria ser válido", c)
		}
	}
}

func TestPlacasInvalidas(t *testing.T) {

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
		if err == nil {
			t.Errorf("Car plate %s deveria ser inválido", c)
		} else if err.Error() != expectedError {
			t.Errorf("Mensagem deveria ser \"invalid car plate\": %s", err.Error())
		}
	}
}
