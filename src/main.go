package main

import (
	"github.com/deploydesexta/cccat12/src/ride"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.POST("/calculate_ride", CalculateRide)
	log.Fatal(e.Start(":3000"))
}

type (
	response struct {
		Price float64 `json:"price"`
	}

	requestBody struct {
		Segments []struct {
			Distance float64 `json:"distance"`
			Date     string  `json:"date"`
		} `json:"segments"`
	}
)

func CalculateRide(c echo.Context) error {
	var body requestBody

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

	res := response{
		Price: r.Calculate(),
	}

	return c.JSON(http.StatusOK, res)
}
