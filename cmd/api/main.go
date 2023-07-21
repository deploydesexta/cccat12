package main

import (
	"github.com/deploydesexta/cccat12/src/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/src/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/src/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/src/application/usecase/getpassenger"
	"github.com/deploydesexta/cccat12/src/infrastructure/http"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/driverpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/passengerpg"
	"github.com/deploydesexta/cccat12/src/infrastructure/repository/pgdb"
	_ "github.com/lib/pq"
	"log"
)

// main composition root
func main() {
	// Drivers
	conn := pgdb.New()
	// Interface Adapters
	driverRepository := driverpg.New(conn)
	passengerRepository := passengerpg.New(conn)
	// UseCases
	calculateRideUseCase := calculateride.New()
	createDriverUseCase := createdriver.New(driverRepository)
	getDriverUseCase := getdriver.New(driverRepository)
	createPassengerUseCase := createpassenger.New(passengerRepository)
	getPassengerUseCase := getpassenger.New(passengerRepository)
	// Frameworks
	router := http.NewRootRouter(calculateRideUseCase, createDriverUseCase, getDriverUseCase, createPassengerUseCase, getPassengerUseCase)
	//server := NewFiberRouterAdapter()
	server := NewEchoRouterAdapter()
	server.Router(router)
	log.Fatal(server.Start(":3000"))
}
