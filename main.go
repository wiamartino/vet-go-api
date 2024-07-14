package main

import (
	"go-vet/config"
	"go-vet/routes"
)

func main() {
	r := routes.SetupRouter()
	config.ConnectDatabase()
	r.Run(":8080")
}
