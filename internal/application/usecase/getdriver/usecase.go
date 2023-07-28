package getdriver

import (
	"context"
	"github.com/deploydesexta/cccat12/internal/application/repository"
)

type (
	UseCase struct {
		repository repository.DriverRepository
	}

	Input struct {
		DriverId string
	}

	Output struct {
		DriverId string
		Document string
		CarPlate string
		Email    string
		Name     string
	}
)

func New(repository repository.DriverRepository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (c UseCase) Execute(ctx context.Context, input Input) (Output, error) {

	d, err := c.repository.Get(ctx, input.DriverId)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		DriverId: d.DriverId(),
		Document: d.Document().Value(),
		CarPlate: d.CarPlate().Value(),
		Email:    d.Email().Value(),
		Name:     d.Name(),
	}

	return output, nil
}
