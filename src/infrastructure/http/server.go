package http

import (
	"context"
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/src/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
)

type (
	Request interface {
		Bind(i interface{}) error
		JSON(code int, i interface{}) error
		String(code int, s string) error
		Param(name string) string
		Context() context.Context
	}

	HandlerFunc func(c Request) error

	Server interface {
		Post(path string, handler HandlerFunc)
		Get(path string, handler HandlerFunc)
		Router(ctrl Router)
		Start(port string) error
	}

	Router interface {
		Bind(r Server)
	}

	MainRouter struct {
		calculateRide   *calculateride.UseCase
		createDriver    *createdriver.UseCase
		getDriver       *getdriver.UseCase
		createPassenger *createpassenger.UseCase
		getPassenger    *getpassenger.UseCase
	}
)

func NewRootRouter(
	calculateRide *calculateride.UseCase,
	createDriver *createdriver.UseCase,
	getDriver *getdriver.UseCase,
	createPassenger *createpassenger.UseCase,
	getPassenger *getpassenger.UseCase,
) Router {
	return &MainRouter{
		calculateRide,
		createDriver,
		getDriver,
		createPassenger,
		getPassenger,
	}
}

func (c *MainRouter) Bind(s Server) {
	s.Post("/calculate_ride", c.CalculateRide)
	s.Post("/passengers", c.CreatePassenger)
	s.Get("/passengers/:passengerId", c.GetPassenger)
	s.Post("/drivers", c.CreateDriver)
	s.Get("/drivers/:driverId", c.GetDriver)
}
