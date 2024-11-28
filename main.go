package main

import (
	"log"
	"oghenekparobor/market-lens/api"
	"oghenekparobor/market-lens/config"
)

func main() {
	// Loads the .env file
	config.LoadEnv()

	// runs all DB operation (i.e. Connection and AutoMigration)
	config.InitDB(config.GetPostgresConfig())

	// entry point for the user's server
	server := api.MarketLensApiServer()

	err := server.Run()

	// if there's an error log it
	if err != nil {
		log.Fatal(err)
	}
}
