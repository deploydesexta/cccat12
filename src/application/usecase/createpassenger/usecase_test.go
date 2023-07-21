package createpassenger_test

import (
	"context"
	"github.com/deploydesexta/cccat12/mocks/nooprepo"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/passengerpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/pgdb"
	"testing"
)

type (
	CreateInput struct {
		Document string
		Email    string
		Name     string
	}

	GetInput struct {
		PassengerId string
	}
)

// narrow integration test
func TestDeveriaCadastrarPassageiro_InMemoryNarrowTest(t *testing.T) {
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	// fake test pattern
	repo := nooprepo.NewPassengerRepository()
	useCase := createpassenger.New(repo)
	_, err := useCase.Execute(context.Background(), createpassenger.Input(input))
	if err != nil {
		t.Errorf("Error executing use case: %v", err)
		return
	}
}

// narrow integration test
func TestDeveriaCadastrarPassageiro_IntegratedNarrowTest(t *testing.T) {
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	createUseCase := createpassenger.New(passengerpg.New(pgdb.New()))
	_, err := createUseCase.Execute(context.Background(), createpassenger.Input(input))
	if err != nil {
		t.Errorf("Error executing use case: %v", err)
		return
	}
}

// broad integration test
func TestDeveriaCadastrarPassageiro_IntegratedBroadTest(t *testing.T) {
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	createUseCase := createpassenger.New(passengerpg.New(pgdb.New()))
	d, err := createUseCase.Execute(context.Background(), createpassenger.Input(input))
	if err != nil {
		t.Errorf("Error executing create use case: %v", err)
		return
	}

	getUseCase := getpassenger.New(passengerpg.New(pgdb.New()))
	r, err := getUseCase.Execute(context.Background(), getpassenger.Input(GetInput{d.PassengerId}))
	if err != nil {
		t.Errorf("Error executing get use case: %v", err)
		return
	}

	if r.Email != expectedEmail {
		t.Errorf("Email deveria ser %s", expectedEmail)
	}

	if r.Name != expectedName {
		t.Errorf("Nome deveria ser %s", expectedName)
	}

	if r.Document != expectedDocument {
		t.Errorf("CPF deveria ser %s", expectedDocument)
	}
}
