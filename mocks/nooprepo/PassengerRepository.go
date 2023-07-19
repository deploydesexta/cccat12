package nooprepo

import (
	"context"
	"github.com/deploydesexta/cccat12/src/domain/passenger"
)

type InMemoryPassengerRepository struct {
	drivers map[string]passenger.Passenger
}

func NewPassengerRepository() *InMemoryPassengerRepository {
	return &InMemoryPassengerRepository{
		drivers: make(map[string]passenger.Passenger),
	}
}

func (r InMemoryPassengerRepository) Save(ctx context.Context, d passenger.Passenger) error {
	r.drivers[d.PassengerId()] = d
	return nil
}

func (r InMemoryPassengerRepository) Get(ctx context.Context, driverId string) (passenger.Passenger, error) {
	return r.drivers[driverId], nil
}
