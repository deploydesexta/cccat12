package email_test

import (
	"github.com/deploydesexta/cccat12/internal/domain/person/email"
	"testing"
)

func TestEmailsValidos(t *testing.T) {

	emails := []string{
		"john.doe@gmail.com",
		"john-doe@gmail.com",
		"1@gmail.com",
	}

	for _, c := range emails {
		if _, err := email.New(c); err != nil {
			t.Errorf("Email %s deveria ser válido", c)
		}
	}
}

func TestEmailsInvalidos(t *testing.T) {

	emails := []string{
		"john.doe@gmail",
		"john.doe@gmail..com",
		"",
	}

	expectedError := "invalid email"
	for _, c := range emails {
		_, err := email.New(c)
		if err == nil {
			t.Errorf("Email %s deveria ser inválido", c)
		} else if err.Error() != expectedError {
			t.Errorf("Mensagem deveria ser \"%s\": %s", expectedError, err.Error())
		}
	}
}
