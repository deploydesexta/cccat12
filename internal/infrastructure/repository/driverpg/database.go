package driverpg

import (
	"context"
	"errors"
	"github.com/deploydesexta/cccat12/internal/domain/driver"
	"github.com/deploydesexta/cccat12/internal/infrastructure/database"
)

const (
	selectStatement = "SELECT driver_id, car_plate, document, email, name FROM cccat12.driver WHERE driver_id = $1"
	insertStatement = "INSERT INTO cccat12.driver (driver_id, car_plate, document, email, name) VALUES ($1, $2, $3, $4, $5)"
)

var (
	driverNotFound = errors.New("driver not found")
)

type (
	// DriverRepositoryDatabase Interface Adapter
	DriverRepositoryDatabase struct {
		conn database.DbConn
	}

	DriverEntity struct {
		Id       string
		Document string
		CarPlate string
		Email    string
		Name     string
	}
)

func New(conn database.DbConn) DriverRepositoryDatabase {
	return DriverRepositoryDatabase{conn}
}

func (r DriverRepositoryDatabase) Save(ctx context.Context, d driver.Driver) error {
	_, err := r.conn.ExecContext(ctx, insertStatement, d.DriverId(), d.CarPlate().Value(), d.Document().Value(), d.Email().Value(), d.Name())
	if err != nil {
		return err
	}
	return nil
}

func (r DriverRepositoryDatabase) Get(ctx context.Context, driverId string) (driver.Driver, error) {
	row := r.conn.QueryRow(selectStatement, driverId)
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
