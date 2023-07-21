package createpassenger

import (
	"context"
	"github.com/deploydesexta/cccat12/src/application/repository"
	"github.com/deploydesexta/cccat12/src/domain/passenger"
)

type (
	UseCase struct {
		repository repository.PassengerRepository
	}

	Input struct {
		Document string
		Email    string
		Name     string
	}

	Output struct {
		PassengerId string
	}
)

func New(repository repository.PassengerRepository) *UseCase {
	return &UseCase{
		repository: repository,
	}
}

func (c UseCase) Execute(ctx context.Context, input Input) (Output, error) {
	p, err := passenger.New(input.Document, input.Email, input.Name)
	if err != nil {
		return Output{}, err
	}

	err = c.repository.Save(ctx, p)
	if err != nil {
		return Output{}, err
	}

	return Output{p.PassengerId()}, nil
}
