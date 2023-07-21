package http

import (
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/src/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
	"net/http"
)

type (
	driverBody struct {
		Document string `jsonutil:"document"`
		CarPlate string `jsonutil:"carPlate"`
		Email    string `jsonutil:"email"`
		Name     string `jsonutil:"name"`
	}

	driverResponse struct {
		DriverId string `jsonutil:"driverId"`
	}

	driverDTO struct {
		DriverId string `jsonutil:"driverId"`
		Document string `jsonutil:"document"`
		CarPlate string `jsonutil:"carPlate"`
		Email    string `jsonutil:"email"`
		Name     string `jsonutil:"name"`
	}

	passengerBody struct {
		Document string `jsonutil:"document"`
		Email    string `jsonutil:"email"`
		Name     string `jsonutil:"name"`
	}

	passengerResponse struct {
		PassengerId string `jsonutil:"passengerId"`
	}

	passengerDTO struct {
		PassengerId string `jsonutil:"passengerId"`
		Document    string `jsonutil:"document"`
		Email       string `jsonutil:"email"`
		Name        string `jsonutil:"name"`
	}

	rideBody struct {
		Segments *[]struct {
			Distance float64 `jsonutil:"distance"`
			Date     string  `jsonutil:"date"`
		}
	}

	rideResponse struct {
		Price float64 `jsonutil:"price"`
	}
)

func (c *MainRouter) CalculateRide(req Request) error {
	var body rideBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusBadRequest, err.Error())
	}

	output, err := c.calculateRide.Execute(calculateride.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, rideResponse{output.Price})
}

func (c *MainRouter) CreatePassenger(req Request) error {
	var body passengerBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	output, err := c.createPassenger.Execute(req.Context(), createpassenger.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, passengerResponse{output.PassengerId})
}

func (c *MainRouter) GetPassenger(req Request) error {
	passengerId := req.Param("passengerId")

	output, err := c.getPassenger.Execute(req.Context(), getpassenger.Input{PassengerId: passengerId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, passengerDTO(output))
}

func (c *MainRouter) CreateDriver(req Request) error {
	var body driverBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	output, err := c.createDriver.Execute(req.Context(), createdriver.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, driverResponse(output))
}

func (c *MainRouter) GetDriver(req Request) error {
	driverId := req.Param("driverId")

	d, err := c.getDriver.Execute(req.Context(), getdriver.Input{DriverId: driverId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, driverDTO(d))
}
