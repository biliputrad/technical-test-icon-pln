package main

import (
	"fmt"
	"log"
	"technical-test-icon-pln/practical-test/common/constants"
	"technical-test-icon-pln/practical-test/config/database/postgres"
	"technical-test-icon-pln/practical-test/config/env"
	"technical-test-icon-pln/practical-test/config/route"
	routeRegister "technical-test-icon-pln/practical-test/route"
)

func main() {
	//Init config app.env
	config, err := env.LoadConfig(".")
	if err != nil {
		message := fmt.Sprintf("%s can't load configuration file", constants.Configuration)
		log.Fatal(message)
	}

	//Init database
	db := postgres.ConfigurationPostgres(config)

	//Init router
	router := route.InitRouter(config)

	//Register routes
	apiV1 := router.Group("api/v1")
	routeRegister.RouteRegister(db, apiV1)

	//Run route
	route.RunRoute(config, router)
}
