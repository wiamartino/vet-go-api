package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"
	"log"
)

func main() {
	// Initialize logging first
	utils.SetupLogging()

	// Initialize database connection
	_, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Drop tables and seed (only for development)
	if err := database.DropAllTablesAndSeed(); err != nil {
		log.Fatalf("Failed to drop tables and seed database: %v", err)
	}

	// Setup router (this will reuse the existing database connection)
	r := routes.SetupRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
