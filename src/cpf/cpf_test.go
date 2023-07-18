package cpf_test

import (
	"github.com/deploydesexta/cccat12/src/cpf"
	"testing"
)

func TestCPFsValidos(t *testing.T) {

	cpfs := []string{
		"83432616074",
		"74587887803",
		"87175659520",
	}

	for _, c := range cpfs {
		if _, err := cpf.NewCpf(c); err != nil {
			t.Errorf("CPF %s deveria ser válido", c)
		}
	}
}

func TestCPFsInvalidos(t *testing.T) {

	cpfs := []string{
		"83432616076",
		"99999999999",
		"834326160",
		"",
	}

	expectedError := "invalid cpf"
	for _, c := range cpfs {
		_, err := cpf.NewCpf(c)
		if err == nil {
			t.Errorf("CPF %s deveria ser inválido", c)
		} else if err.Error() != expectedError {
			t.Errorf("Mensagem deveria ser \"invalid cpf\": %s", err.Error())
		}
	}
}
