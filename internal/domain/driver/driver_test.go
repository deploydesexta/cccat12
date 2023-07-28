package driver_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/driver"
	"testing"
)

func TestCriarMotorista(t *testing.T) {
	expectedCarPlate := "AAA9999"
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	motorista, err := driver.New(expectedCarPlate, expectedDocument, expectedEmail, expectedName)
	if err != nil {
		t.Errorf("Motorista deveria ser válido")
	}

	if motorista.DriverId() == "" {
		t.Errorf("Id deveria ser gerado")
	}

	if motorista.CarPlate().Value() != expectedCarPlate {
		t.Errorf("Placa deveria ser %s", expectedCarPlate)
	}

	if motorista.Email().Value() != expectedEmail {
		t.Errorf("Email deveria ser %s", expectedEmail)
	}

	if motorista.Name() != expectedName {
		t.Errorf("Nome deveria ser %s", expectedName)
	}

	if motorista.Document().Value() != expectedDocument {
		t.Errorf("CPF deveria ser %s", expectedDocument)
	}
}

func TestNaoDeveriaCriarMotoristaComCPFInvalido(t *testing.T) {
	expectedError := "invalid cpf"

	_, err := driver.New("AAA9999", "1", "john.doe@gmail.com", "John Doe")
	if err == nil {
		t.Errorf("Um erro deveria ter sido retornado")
	}

	if err.Error() != expectedError {
		t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
	}
}

func TestNaoDeveriaCriarMotoristaComEmailInvalido(t *testing.T) {
	expectedError := "invalid email"

	_, err := driver.New("AAA9999", "83432616074", "john.doe@gmail", "John Doe")
	if err == nil {
		t.Errorf("Um erro deveria ter sido retornado")
	}

	if err.Error() != expectedError {
		t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
	}
}

func TestNaoDeveriaCriarMotoristaComPlacaInvalido(t *testing.T) {
	expectedError := "invalid car plate"

	_, err := driver.New("AAA999", "83432616074", "john.doe@gmail.com", "John Doe")
	if err == nil {
		t.Errorf("Um erro deveria ter sido retornado")
	}

	if err.Error() != expectedError {
		t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
	}
}
