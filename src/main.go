package main

import (
	"database/sql"
	"github.com/deploydesexta/cccat12/src/cpf"
	"github.com/deploydesexta/cccat12/src/ride"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
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
		Id       string `json:"string"`
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
		Id       string `json:"id"`
		Document string `json:"document"`
		Email    string `json:"email"`
		Name     string `json:"name"`
	}

	rideBody struct {
		Segments []struct {
			Distance float64 `json:"distance"`
			Date     string  `json:"date"`
		} `json:"segments"`
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

	r := ride.New()
	for _, s := range body.Segments {
		t, err := time.Parse("2006-01-02T15:04:05", s.Date)
		if err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}

		err = r.AddSegment(s.Distance, t)
		if err != nil {
			return c.String(http.StatusUnprocessableEntity, err.Error())
		}
	}

	return c.JSON(http.StatusOK, rideResponse{r.Calculate()})
}

func CreatePassenger(c echo.Context) error {
	var body passengerBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	if _, err := cpf.NewCpf(body.Document); err != nil {
		return c.String(http.StatusUnprocessableEntity, "Invalid cpf")
	}

	db := pgp()
	defer db.Close()

	passengerId := uuid.New().String()
	_, err := db.ExecContext(c.Request().Context(), "INSERT INTO cccat12.passenger (passenger_id, name, email, document) VALUES ($1, $2, $3, $4)", passengerId, body.Name, body.Email, body.Document)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, passengerResponse{passengerId})
}

func GetPassenger(c echo.Context) error {
	passengerId := c.Param("passengerId")

	db := pgp()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM cccat12.passenger WHERE passenger_id = $1", passengerId)
	if row == nil {
		return c.String(http.StatusUnprocessableEntity, "Passenger not found")
	}

	var passenger passengerDTO
	err := row.Scan(&passenger.Id, &passenger.Name, &passenger.Email, &passenger.Document)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, passenger)
}

func CreateDriver(c echo.Context) error {
	var body driverBody
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	if _, err := cpf.NewCpf(body.Document); err != nil {
		return c.String(http.StatusUnprocessableEntity, "Invalid cpf")
	}

	db := pgp()
	defer db.Close()

	driverId := uuid.New().String()
	_, err := db.ExecContext(c.Request().Context(), "INSERT INTO cccat12.driver (driver_id, name, email, document, car_plate) VALUES ($1, $2, $3, $4, $5)", driverId, body.Name, body.Email, body.Document, body.CarPlate)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, driverResponse{driverId})
}

func GetDriver(c echo.Context) error {
	driverId := c.Param("driverId")

	db := pgp()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM cccat12.driver WHERE driver_id = $1", driverId)
	if row == nil {
		return c.String(http.StatusUnprocessableEntity, "Driver not found")
	}

	var driver driverDTO
	err := row.Scan(&driver.Id, &driver.Name, &driver.Email, &driver.Document, &driver.CarPlate)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, driver)
}

func pgp() *sql.DB {
	connStr := "postgresql://postgres:123456@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
