package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
)

func main() {
	r := routes.SetupRouter()
	database.ConnectDatabase()
	r.Run(":8080")
}
