package config

import (
	"go-vet/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=name_db port=5432 sslmode=disable TimeZone=America/Buenos_Aires"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Client{}, &models.Pet{})
	database.AutoMigrate(&models.Appointment{})
	database.AutoMigrate(&models.Veterinarian{}, &models.Treatment{}, &models.Invoice{})
	DB = database
}
