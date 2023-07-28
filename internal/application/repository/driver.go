package repository

import (
	"context"
	"github.com/deploydesexta/cccat12/internal/domain/driver"
)

type DriverRepository interface {
	Save(ctx context.Context, d driver.Driver) error
	Get(ctx context.Context, driverId string) (driver.Driver, error)
}
