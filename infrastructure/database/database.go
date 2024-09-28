package database

import (
	"fmt"
	"go-vet/domain"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDatabase() (*DB, error) {

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

	database.AutoMigrate(&domain.User{})
	database.AutoMigrate(&domain.Pet{})
	database.AutoMigrate(&domain.Client{})
	database.AutoMigrate(&domain.Appointment{})
	database.AutoMigrate(&domain.Veterinarian{}, &domain.Treatment{}, &domain.Invoice{})
	database.AutoMigrate(&domain.Medication{})
	return &DB{database}, nil
}
