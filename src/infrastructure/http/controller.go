package http

import (
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/src/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/driverpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/passengerpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/pgdb"
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
		Segments *[]struct {
			Distance float64 `json:"distance"`
			Date     string  `json:"date"`
		}
	}

	rideResponse struct {
		Price float64 `json:"price"`
	}
)

func CalculateRide(req Request) error {
	var body rideBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusBadRequest, err.Error())
	}

	useCase := calculateride.NewUseCase()
	output, err := useCase.Execute(calculateride.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, rideResponse{output.Price})
}

func CreatePassenger(req Request) error {
	var body passengerBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	conn := pgdb.New()
	useCase := createpassenger.New(passengerpg.New(conn))
	output, err := useCase.Execute(req.Context(), createpassenger.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, passengerResponse{output.PassengerId})
}

func GetPassenger(req Request) error {
	passengerId := req.Param("passengerId")

	conn := pgdb.New()
	useCase := getpassenger.New(passengerpg.New(conn))
	output, err := useCase.Execute(req.Context(), getpassenger.Input{PassengerId: passengerId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, passengerDTO(output))
}

func CreateDriver(req Request) error {
	var body driverBody
	if err := req.Bind(&body); err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	conn := pgdb.New()
	useCase := createdriver.New(driverpg.New(conn))
	output, err := useCase.Execute(req.Context(), createdriver.Input(body))
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, driverResponse(output))
}

func GetDriver(req Request) error {
	driverId := req.Param("driverId")

	conn := pgdb.New()
	useCase := getdriver.New(driverpg.New(conn))
	d, err := useCase.Execute(req.Context(), getdriver.Input{DriverId: driverId})
	if err != nil {
		return req.String(http.StatusUnprocessableEntity, err.Error())
	}

	return req.JSON(http.StatusOK, driverDTO(d))
}
