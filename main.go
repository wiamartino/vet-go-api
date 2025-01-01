package main

import (
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	r := routes.SetupRouter()
	database.ConnectDatabase()
	// logging
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	r.Run(":8080")
}
