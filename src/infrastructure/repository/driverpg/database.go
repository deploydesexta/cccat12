package driverpg

import (
	"context"
	"database/sql"
	"errors"
	"github.com/deploydesexta/cccat12/src/domain/driver"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/pgdb"
)

const (
	selectStatement = "SELECT driver_id, car_plate, document, email, name FROM cccat12.driver WHERE driver_id = $1"
	insertStatement = "INSERT INTO cccat12.driver (driver_id, car_plate, document, email, name) VALUES ($1, $2, $3, $4, $5)"
)

var (
	driverNotFound = errors.New("driver not found")
)

type (
	DriverRepositoryDatabase struct {
		db *sql.DB
	}

	DriverEntity struct {
		Id       string
		Document string
		CarPlate string
		Email    string
		Name     string
	}
)

func New() DriverRepositoryDatabase {
	return DriverRepositoryDatabase{db: pgdb.New()}
}

func (r DriverRepositoryDatabase) Save(ctx context.Context, d driver.Driver) error {
	_, err := r.db.ExecContext(ctx, insertStatement, d.DriverId(), d.CarPlate().Value(), d.Document().Value(), d.Email().Value(), d.Name())
	if err != nil {
		return err
	}
	return nil
}

func (r DriverRepositoryDatabase) Get(ctx context.Context, driverId string) (driver.Driver, error) {
	row := r.db.QueryRow(selectStatement, driverId)
	if row == nil {
		return driver.Driver{}, driverNotFound
	}

	var e DriverEntity
	err := row.Scan(&e.Id, &e.CarPlate, &e.Document, &e.Email, &e.Name)
	if err != nil {
		return driver.Driver{}, err
	}

	return driver.From(e.Id, e.CarPlate, e.Document, e.Email, e.Name)
}
