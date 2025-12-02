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
	mu         sync.RWMutex
)

// GetDB returns the current database instance
func GetDB() *DB {
	mu.RLock()
	defer mu.RUnlock()
	return dbInstance
}

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
	// Check if already connected
	mu.RLock()
	if dbInstance != nil {
		mu.RUnlock()
		return dbInstance, nil
	}
	mu.RUnlock()

	// Acquire write lock to initialize
	mu.Lock()
	defer mu.Unlock()

	// Double-check after acquiring write lock
	if dbInstance != nil {
		return dbInstance, nil
	}

	database, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := database.AutoMigrate(models()...); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	dbInstance = &DB{database}
	return dbInstance, nil
}

// DropAllTablesAndSeed drops all tables and reseeds the database
func DropAllTablesAndSeed() error {
	// Ensure database connection is initialized
	if dbInstance == nil {
		_, err := ConnectDatabase()
		if err != nil {
			return fmt.Errorf("failed to initialize database connection: %w", err)
		}
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
