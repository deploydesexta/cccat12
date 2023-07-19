package nooprepo

import (
	"context"
	"github.com/deploydesexta/cccat12/src/domain/driver"
)

type InMemoryDriverRepository struct {
	drivers map[string]driver.Driver
}

func NewDriverRepository() *InMemoryDriverRepository {
	return &InMemoryDriverRepository{
		drivers: make(map[string]driver.Driver),
	}
}

func (r InMemoryDriverRepository) Save(ctx context.Context, d driver.Driver) error {
	r.drivers[d.DriverId()] = d
	return nil
}

func (r InMemoryDriverRepository) Get(ctx context.Context, driverId string) (driver.Driver, error) {
	return r.drivers[driverId], nil
}
