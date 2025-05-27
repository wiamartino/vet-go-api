package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"
)

func main() {
	// Initialize logging first
	utils.SetupLogging()

	// Initialize database connection once
	db, err := database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Drop tables and seed database (only for development)
	if err := database.DropAllTablesAndSeed(); err != nil {
		panic("Failed to drop tables and seed database: " + err.Error())
	}

	// Setup router with existing database connection
	r := routes.SetupRouter(db)

	r.Run(":8080")
}
