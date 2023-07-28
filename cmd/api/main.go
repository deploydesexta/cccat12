package main

import (
	"github.com/deploydesexta/cccat12/internal/application/usecase/calculateride"
	"github.com/deploydesexta/cccat12/internal/application/usecase/createdriver"
	"github.com/deploydesexta/cccat12/internal/application/usecase/createpassenger"
	"github.com/deploydesexta/cccat12/internal/application/usecase/getdriver"
	"github.com/deploydesexta/cccat12/internal/application/usecase/getpassenger"
	"github.com/deploydesexta/cccat12/internal/infrastructure/database/pgdb"
	"github.com/deploydesexta/cccat12/internal/infrastructure/repository/driverpg"
	"github.com/deploydesexta/cccat12/internal/infrastructure/repository/passengerpg"
	"github.com/deploydesexta/cccat12/internal/infrastructure/web"
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
	router := web.NewRootRouter(calculateRideUseCase, createDriverUseCase, getDriverUseCase, createPassengerUseCase, getPassengerUseCase)
	//server := NewFiberRouterAdapter()
	server := web.NewEchoRouterAdapter()
	server.Router(router)
	log.Fatal(server.Start(":3000"))
}
