package passenger_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/passenger"
	"testing"
)

func TestCriarPassageiro(t *testing.T) {
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	passageiro, err := passenger.New(expectedDocument, expectedEmail, expectedName)
	if err != nil {
		t.Errorf("Passageiro deveria ser válido")
	}

	if passageiro.PassengerId() == "" {
		t.Errorf("Id deveria ser gerado")
	}

	if passageiro.Email().Value() != expectedEmail {
		t.Errorf("Email deveria ser %s", expectedEmail)
	}

	if passageiro.Name() != expectedName {
		t.Errorf("Nome deveria ser %s", expectedName)
	}

	if passageiro.Document().Value() != expectedDocument {
		t.Errorf("CPF deveria ser %s", expectedDocument)
	}
}

func TestNaoDeveriaCriarPassageiroComCPFInvalido(t *testing.T) {
	expectedError := "invalid cpf"

	_, err := passenger.New("1", "john.doe@gmail.com", "John Doe")
	if err == nil {
		t.Errorf("Um erro deveria ter sido retornado")
	}

	if err.Error() != expectedError {
		t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
	}
}

func TestNaoDeveriaCriarPassageiroComEmailInvalido(t *testing.T) {
	expectedError := "invalid email"

	_, err := passenger.New("83432616074", "john.doe@gmail", "John Doe")
	if err == nil {
		t.Errorf("Um erro deveria ter sido retornado")
	}

	if err.Error() != expectedError {
		t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
	}
}
