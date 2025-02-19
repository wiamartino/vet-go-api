package main

import (
	"go-vet/routes"
	"go-vet/utils"
)

func main() {
	r := routes.SetupRouter()
	utils.SetupLogging()
	r.Run(":8080")
}
