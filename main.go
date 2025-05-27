package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"
	"os"
)

func main() {
	// Setup logging first
	utils.SetupLogging()

	// Establish database connection once
	db, err := database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Drop and seed database (only for development environment)
	env := os.Getenv("GO_ENV")
	if env == "" || env == "development" {
		if err := database.DropAllTablesAndSeed(); err != nil {
			panic("Failed to drop tables and seed database: " + err.Error())
		}
	}

	// Setup router with the existing database connection
	r := routes.SetupRouter(db)

	r.Run(":8080")
}
