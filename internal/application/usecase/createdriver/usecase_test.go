package createdriver_test

import (
	"context"
	"github.com/deploydesexta/cccat12/internal/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/internal/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/internal/infrastructure/database/pgdb"
	"github.com/deploydesexta/cccat12/internal/infrastructure/repository/driverpg"
	"github.com/deploydesexta/cccat12/mocks/nooprepo"
	"testing"
)

type (
	CreateInput struct {
		Document string
		CarPlate string
		Email    string
		Name     string
	}

	GetInput struct {
		DriverId string
	}
)

// narrow integration test
func TestDeveriaCadastrarMotorista_InMemoryNarrowTest(t *testing.T) {
	expectedCarPlate := "AAA9999"
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		CarPlate: expectedCarPlate,
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	// fake test pattern
	repo := nooprepo.NewDriverRepository()
	useCase := createdriver.New(repo)
	_, err := useCase.Execute(context.Background(), createdriver.Input(input))
	if err != nil {
		t.Errorf("Error executing use case: %v", err)
		return
	}
}

// narrow integration test
func TestDeveriaCadastrarMotorista_IntegratedNarrowTest(t *testing.T) {
	expectedCarPlate := "AAA9999"
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		CarPlate: expectedCarPlate,
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	createUseCase := createdriver.New(driverpg.New(pgdb.New()))
	_, err := createUseCase.Execute(context.Background(), createdriver.Input(input))
	if err != nil {
		t.Errorf("Error executing use case: %v", err)
		return
	}
}

// broad integration test
func TestDeveriaCadastrarMotorista_IntegratedBroadTest(t *testing.T) {
	expectedCarPlate := "AAA9999"
	expectedDocument := "83432616074"
	expectedEmail := "john.doe@gmail.com"
	expectedName := "John Doe"

	input := CreateInput{
		CarPlate: expectedCarPlate,
		Document: expectedDocument,
		Email:    expectedEmail,
		Name:     expectedName,
	}

	createUseCase := createdriver.New(driverpg.New(pgdb.New()))
	d, err := createUseCase.Execute(context.Background(), createdriver.Input(input))
	if err != nil {
		t.Errorf("Error executing create use case: %v", err)
		return
	}

	getUseCase := getdriver.New(driverpg.New(pgdb.New()))
	r, err := getUseCase.Execute(context.Background(), getdriver.Input(GetInput{d.DriverId}))
	if err != nil {
		t.Errorf("Error executing get use case: %v", err)
		return
	}

	if r.CarPlate != expectedCarPlate {
		t.Errorf("Placa deveria ser %s", expectedCarPlate)
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
