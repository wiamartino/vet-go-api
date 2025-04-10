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

// Models returns all database models that need to be migrated
func models() []interface{} {
	return []interface{}{
		&AuditLog{},
		&domain.User{},
		&domain.Pet{},
		&domain.Client{},
		&domain.Appointment{},
		&domain.Veterinarian{},
		&domain.Treatment{},
		&domain.Invoice{},
		&domain.Medication{},
	}
}

// getDSN builds the database connection string from environment variables
func getDSN() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
}

// ConnectDatabase initializes the database connection
func ConnectDatabase() (*DB, error) {
	var err error
	once.Do(func() {
		database, dbErr := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
		if dbErr != nil {
			err = fmt.Errorf("failed to connect to database: %w", dbErr)
			return
		}

		if migErr := database.AutoMigrate(models()...); migErr != nil {
			err = fmt.Errorf("failed to migrate database: %w", migErr)
			return
		}

		dbInstance = &DB{database}
	})

	if err != nil {
		return nil, err
	}
	return dbInstance, nil
}

// DropAllTablesAndSeed drops all tables and reseeds the database
func DropAllTablesAndSeed() error {
	if dbInstance == nil {
		return fmt.Errorf("database instance not initialized")
	}

	// Drop all tables
	if err := dbInstance.Migrator().DropTable(models()...); err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	// Recreate tables
	if err := dbInstance.AutoMigrate(models()...); err != nil {
		return fmt.Errorf("failed to recreate tables: %w", err)
	}

	// Seed database
	seedFile, err := os.ReadFile("database/seed.sql")
	if err != nil {
		return fmt.Errorf("failed to read seed file: %w", err)
	}

	if err := dbInstance.Exec(string(seedFile)).Error; err != nil {
		return fmt.Errorf("failed to execute seed file: %w", err)
	}

	return nil
}
