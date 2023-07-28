package web

import (
	"github.com/deploydesexta/cccat12/internal/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/internal/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/internal/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/internal/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/internal/application/usecase/getpassenger"
	"net/http"
)

type (
	driverBody struct {
		Document string `json:"document"`
		CarPlate string `json:"carPlate"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	driverResponse struct {
		DriverId string `json:"driverId"`
	}

	driverDTO struct {
		DriverId string `json:"driverId"`
		Document string `json:"document"`
		CarPlate string `json:"carPlate"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	passengerBody struct {
		Document string `json:"document"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	passengerResponse struct {
		PassengerId string `json:"passengerId"`
	}

	passengerDTO struct {
		PassengerId string `json:"passengerId"`
		Document    string `json:"document"`
		Email       string `json:"email"`
		Name        string `json:"name"`
	}

	rideBody struct {
		Positions *[]struct {
			Lat  float64 `json:"lat"`
			Long float64 `json:"long"`
			Date string  `json:"date"`
		}
	}

	rideResponse struct {
		Price float64 `json:"price"`
	}
)

func (c *RootRouter) CalculateRide(req Request) error {
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

func (c *RootRouter) CreatePassenger(req Request) error {
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

func (c *RootRouter) GetPassenger(req Request) error {
	passengerId := req.Param("passengerId")

	output, err := c.getPassenger.Execute(req.Context(), getpassenger.Input{PassengerId: passengerId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, passengerDTO(output))
}

func (c *RootRouter) CreateDriver(req Request) error {
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

func (c *RootRouter) GetDriver(req Request) error {
	driverId := req.Param("driverId")

	d, err := c.getDriver.Execute(req.Context(), getdriver.Input{DriverId: driverId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, driverDTO(d))
}
