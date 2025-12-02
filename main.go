package main

import (
	"os"

	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logging first
	utils.SetupLogging()

	// Initialize database connection once
	db, err := database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Drop tables and seed database (ONLY in development mode)
	env := os.Getenv("GO_ENV")
	if env == "development" || env == "dev" {
		logrus.Warning("Running in development mode - dropping and seeding database")
		if err := database.DropAllTablesAndSeed(); err != nil {
			panic("Failed to drop tables and seed database: " + err.Error())
		}
	} else {
		logrus.Info("Production mode - skipping database drop and seed")
	}

	// Setup router with existing database connection
	r := routes.SetupRouter(db)

	// Get port from environment or use default
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
