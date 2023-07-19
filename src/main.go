package main

import (
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/src/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/driverpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/passengerpg"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/calculate_ride", CalculateRide)
	e.POST("/passengers", CreatePassenger)
	e.GET("/passengers/:passengerId", GetPassenger)
	e.POST("/drivers", CreateDriver)
	e.GET("/drivers/:driverId", GetDriver)
	log.Fatal(e.Start(":3000"))
}

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
		Segments *[]struct {
			Distance float64 `json:"distance"`
			Date     string  `json:"date"`
		}
	}

	rideResponse struct {
		Price float64 `json:"price"`
	}
)

func CalculateRide(c echo.Context) error {
	var body rideBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	useCase := calculateride.NewUseCase()
	output, err := useCase.Execute(calculateride.Input(body))
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, rideResponse{output.Price})
}

func CreatePassenger(c echo.Context) error {
	var body passengerBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	useCase := createpassenger.New(passengerpg.New())
	output, err := useCase.Execute(c.Request().Context(), createpassenger.Input(body))
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, passengerResponse{output.PassengerId})
}

func GetPassenger(c echo.Context) error {
	passengerId := c.Param("passengerId")

	useCase := getpassenger.New(passengerpg.New())
	output, err := useCase.Execute(c.Request().Context(), getpassenger.Input{PassengerId: passengerId})
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, passengerDTO(output))
}

func CreateDriver(c echo.Context) error {
	var body driverBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	useCase := createdriver.New(driverpg.New())
	output, err := useCase.Execute(c.Request().Context(), createdriver.Input(body))
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, driverResponse(output))
}

func GetDriver(c echo.Context) error {
	driverId := c.Param("driverId")

	useCase := getdriver.New(driverpg.New())
	d, err := useCase.Execute(c.Request().Context(), getdriver.Input{DriverId: driverId})
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, driverDTO(d))
}
