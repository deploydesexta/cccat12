package nooprepo

import (
	"context"
	"github.com/deploydesexta/cccat12/internal/domain/passenger"
)

type InMemoryPassengerRepository struct {
	drivers map[string]passenger.Passenger
}

func NewPassengerRepository() *InMemoryPassengerRepository {
	return &InMemoryPassengerRepository{
		drivers: make(map[string]passenger.Passenger),
	}
}

func (r InMemoryPassengerRepository) Save(_ context.Context, d passenger.Passenger) error {
	r.drivers[d.PassengerId()] = d
	return nil
}

func (r InMemoryPassengerRepository) Get(_ context.Context, driverId string) (passenger.Passenger, error) {
	return r.drivers[driverId], nil
}
