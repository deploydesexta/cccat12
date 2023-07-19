package createdriver

import (
	"context"
	"github.com/deploydesexta/cccat12/src/application/repository"
	"github.com/deploydesexta/cccat12/src/domain/driver"
)

type (
	UseCase struct {
		repository repository.DriverRepository
	}

	Input struct {
		Document string
		CarPlate string
		Email    string
		Name     string
	}

	Output struct {
		DriverId string
	}
)

func New(repository repository.DriverRepository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (c UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	d, err := driver.New(input.CarPlate, input.Document, input.Email, input.Name)
	if err != nil {
		return Output{}, err
	}

	err = c.repository.Save(ctx, d)
	if err != nil {
		return Output{}, err
	}

	return Output{d.DriverId()}, nil
}
