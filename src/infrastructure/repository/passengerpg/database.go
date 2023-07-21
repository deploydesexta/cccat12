package passengerpg

import (
	"context"
	"errors"
	"github.com/deploydesexta/cccat12/src/domain/passenger"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository"
)

const (
	selectStatement = "SELECT passenger_id, document, email, name FROM cccat12.passenger WHERE passenger_id = $1"
	insertStatement = "INSERT INTO cccat12.passenger (passenger_id, document, email, name) VALUES ($1, $2, $3, $4)"
)

var (
	passengerNotFound = errors.New("passenger not found")
)

type (
	// PassengerRepositoryDatabase Interface Adapter
	PassengerRepositoryDatabase struct {
		conn repository.DbConn
	}

	PassengerEntity struct {
		Id       string
		Document string
		Email    string
		Name     string
	}
)

func New(conn repository.DbConn) PassengerRepositoryDatabase {
	return PassengerRepositoryDatabase{conn}
}

func (r PassengerRepositoryDatabase) Save(ctx context.Context, p passenger.Passenger) error {
	_, err := r.conn.ExecContext(ctx, insertStatement, p.PassengerId(), p.Document().Value(), p.Email().Value(), p.Name())
	if err != nil {
		return err
	}
	return nil
}

func (r PassengerRepositoryDatabase) Get(ctx context.Context, driverId string) (passenger.Passenger, error) {
	row := r.conn.QueryRow(selectStatement, driverId)
	if row == nil {
		return passenger.Passenger{}, passengerNotFound
	}

	var e PassengerEntity
	err := row.Scan(&e.Id, &e.Document, &e.Email, &e.Name)
	if err != nil {
		return passenger.Passenger{}, err
	}

	return passenger.From(e.Id, e.Document, e.Email, e.Name)
}
