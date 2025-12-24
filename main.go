package main

import (
	"os"

	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"

	_ "go-vet/docs" // Import generated docs

	"github.com/sirupsen/logrus"
)

// @title Veterinary Management System API
// @version 1.0
// @description API for managing veterinary clinic operations including appointments, medical records, treatments, and more.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.veterinary-system.io/support
// @contact.email support@veterinary-system.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	// Explicit safeguard: only run in development if explicitly set
	// Production/empty defaults to safe mode
	if (env == "development" || env == "dev") && env != "production" && env != "prod" {
		logrus.Warning("⚠️  Running in DEVELOPMENT mode - dropping and seeding database")
		logrus.Warning("⚠️  This will DELETE ALL DATA in the database")
		if err := database.DropAllTablesAndSeed(); err != nil {
			panic("Failed to drop tables and seed database: " + err.Error())
		}
		logrus.Info("✓ Database dropped and seeded successfully")
	} else {
		logrus.Info("Production/Safe mode - skipping database drop and seed")
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
