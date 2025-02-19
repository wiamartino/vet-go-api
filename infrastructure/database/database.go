package database

import (
	"fmt"
	"go-vet/domain"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

type AuditLog struct {
	ID        uint      `gorm:"primaryKey"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	User      string    `json:"user"`
}

var (
	dbInstance *DB
	once       sync.Once
)

func ConnectDatabase() (*DB, error) {
	once.Do(func() {
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

		database.AutoMigrate(&domain.User{}, &domain.Pet{}, &domain.Client{}, &domain.Appointment{}, &domain.Veterinarian{}, &domain.Treatment{}, &domain.Invoice{}, &domain.Medication{}, &AuditLog{})
		dbInstance = &DB{database}
	})

	return dbInstance, nil
}
