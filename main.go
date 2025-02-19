package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"go-vet/utils"
)

func main() {
	r := routes.SetupRouter()
	database.ConnectDatabase()
	utils.SetupLogging()
	r.Run(":8080")
}
