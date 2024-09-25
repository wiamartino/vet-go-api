package config

import (
	"fmt"
	"go-vet/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Pet{})
	database.AutoMigrate(&models.Client{})
	database.AutoMigrate(&models.Appointment{})
	database.AutoMigrate(&models.Veterinarian{}, &models.Treatment{}, &models.Invoice{})
	database.AutoMigrate(&models.Medication{})
	DB = database
}
