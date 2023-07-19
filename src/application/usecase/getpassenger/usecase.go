package getpassenger

import (
	"context"
	"github.com/deploydesexta/cccat12/src/application/repository"
)

type (
	UseCase struct {
		repository repository.PassengerRepository
	}

	Input struct {
		PassengerId string
	}

	Output struct {
		PassengerId string
		Document    string
		Email       string
		Name        string
	}
)

func New(repository repository.PassengerRepository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (c UseCase) Execute(ctx context.Context, input Input) (Output, error) {

	d, err := c.repository.Get(ctx, input.PassengerId)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		PassengerId: d.PassengerId(),
		Document:    d.Document().Value(),
		Email:       d.Email().Value(),
		Name:        d.Name(),
	}

	return output, nil
}
