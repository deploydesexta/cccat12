package driver

import (
	"github.com/deploydesexta/cccat12/src/domain/carplate"
	"github.com/deploydesexta/cccat12/src/domain/cpf"
	"github.com/deploydesexta/cccat12/src/domain/email"
	"github.com/deploydesexta/cccat12/src/domain/identifier"
)

type Driver struct {
	driverId string
	name     string
	document cpf.Cpf
	email    email.Email
	carPlate carplate.CarPlate
}

func New(carPlate, document, mail, name string) (Driver, error) {
	e, err := email.New(mail)
	if err != nil {
		return Driver{}, err
	}

	p, err := carplate.New(carPlate)
	if err != nil {
		return Driver{}, err
	}

	c, err := cpf.New(document)
	if err != nil {
		return Driver{}, err
	}

	id := identifier.Unique()

	return Driver{id, name, c, e, p}, nil
}

func From(id, carPlate, document, mail, name string) (Driver, error) {
	e, err := email.New(mail)
	if err != nil {
		return Driver{}, err
	}

	p, err := carplate.New(carPlate)
	if err != nil {
		return Driver{}, err
	}

	c, err := cpf.New(document)
	if err != nil {
		return Driver{}, err
	}

	return Driver{id, name, c, e, p}, nil
}

func (d Driver) DriverId() string {
	return d.driverId
}

func (d Driver) Name() string {
	return d.name
}

func (d Driver) Document() cpf.Cpf {
	return d.document
}

func (d Driver) Email() email.Email {
	return d.email
}

func (d Driver) CarPlate() carplate.CarPlate {
	return d.carPlate
}
