package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"
)

func main() {
	r := routes.SetupRouter()
	utils.SetupLogging()
	if err := database.DropAllTablesAndSeed(); err != nil {
		panic("Failed to drop tables and seed database: " + err.Error())
	}
	r.Run(":8080")
}
