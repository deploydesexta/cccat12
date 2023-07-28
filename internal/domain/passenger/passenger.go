package passenger

import (
	"github.com/deploydesexta/cccat12/internal/domain/identifier"
	"github.com/deploydesexta/cccat12/internal/domain/person/cpf"
	"github.com/deploydesexta/cccat12/internal/domain/person/email"
)

type Passenger struct {
	passengerId string
	name        string
	document    cpf.Cpf
	email       email.Email
}

func New(document, mail, name string) (Passenger, error) {
	e, err := email.New(mail)
	if err != nil {
		return Passenger{}, err
	}

	c, err := cpf.New(document)
	if err != nil {
		return Passenger{}, err
	}

	id := identifier.Unique()

	return Passenger{id, name, c, e}, nil
}

func From(passengerId, document, mail, name string) (Passenger, error) {
	e, err := email.New(mail)
	if err != nil {
		return Passenger{}, err
	}

	c, err := cpf.New(document)
	if err != nil {
		return Passenger{}, err
	}

	return Passenger{passengerId, name, c, e}, nil
}

func (p Passenger) PassengerId() string {
	return p.passengerId
}

func (p Passenger) Name() string {
	return p.name
}

func (p Passenger) Document() cpf.Cpf {
	return p.document
}

func (p Passenger) Email() email.Email {
	return p.email
}
