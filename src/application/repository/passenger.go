package repository

import (
	"context"
	"github.com/deploydesexta/cccat12/src/domain/passenger"
)

type PassengerRepository interface {
	Save(ctx context.Context, d passenger.Passenger) error
	Get(ctx context.Context, driverId string) (passenger.Passenger, error)
}
