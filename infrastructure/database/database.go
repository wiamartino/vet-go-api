package database

import (
	"fmt"
	"go-vet/config"
	"go-vet/domain"
	"log"
	"os"
	"sync"
	"time"

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
		&domain.Client{},
		&domain.Veterinarian{},
		&domain.Pet{},
		&domain.Allergy{},
		&domain.Appointment{},
		&domain.Treatment{},
		&domain.Medication{},
		&domain.Invoice{},
		&domain.MedicalRecord{},
		&domain.Vaccination{},
		&domain.Surgery{},
	}
}

// getDSN builds the database connection string from config.
// Falls back to os.Getenv for backward compatibility (e.g. integration tests
// that set env vars directly without calling config.LoadConfig).
func getDSN() string {
	if config.AppConfig != nil {
		dbCfg := config.AppConfig.Database
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.Name, dbCfg.Port, dbCfg.SSLMode, dbCfg.TimeZone,
		)
	}

	// Fallback for tests that don't load config
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

	// Apply connection pool settings from config (fix #24)
	if config.AppConfig != nil {
		sqlDB, err := database.DB()
		if err == nil {
			sqlDB.SetMaxOpenConns(config.AppConfig.Database.MaxOpenConns)
			sqlDB.SetMaxIdleConns(config.AppConfig.Database.MaxIdleConns)
			sqlDB.SetConnMaxLifetime(config.AppConfig.Database.ConnMaxLifetime)
		}
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

	// Drop tables in reverse order (child tables first, then parent tables)
	// Drop each table individually to ensure proper foreign key constraint handling
	tablesToDrop := []interface{}{
		&domain.Surgery{},       // references Pet, Veterinarian
		&domain.Vaccination{},   // references Pet, Veterinarian
		&domain.MedicalRecord{}, // references Pet, Veterinarian, Appointment
		&domain.Invoice{},       // references Client, Appointment
		&domain.Medication{},    // independent
		&domain.Treatment{},     // independent
		&domain.Appointment{},   // references Pet, Veterinarian
		&domain.Allergy{},       // references Pet, Veterinarian
		&domain.Pet{},           // references Client
		&domain.Veterinarian{},  // independent
		&domain.Client{},        // independent
		&domain.User{},          // independent
		&AuditLog{},             // independent
	}

	for _, table := range tablesToDrop {
		if err := dbInstance.Migrator().DropTable(table); err != nil {
			log.Printf("Warning: failed to drop table %T: %v", table, err)
		}
	}

	// Recreate tables in correct order
	if err := dbInstance.AutoMigrate(models()...); err != nil {
		return fmt.Errorf("failed to recreate tables: %w", err)
	}

	// Seed database
	seedFile, err := os.ReadFile("database/seed.sql")
	if err != nil {
		return fmt.Errorf("failed to read seed file: %w", err)
	}

	// Disable foreign key checks temporarily for seeding
	if err := dbInstance.Exec("SET session_replication_role = 'replica';").Error; err != nil {
		return fmt.Errorf("failed to disable foreign key checks: %w", err)
	}

	// Execute seed file
	if err := dbInstance.Exec(string(seedFile)).Error; err != nil {
		// Re-enable foreign key checks before returning error
		dbInstance.Exec("SET session_replication_role = 'origin';")
		return fmt.Errorf("failed to execute seed file: %w", err)
	}

	// Re-enable foreign key checks
	if err := dbInstance.Exec("SET session_replication_role = 'origin';").Error; err != nil {
		return fmt.Errorf("failed to re-enable foreign key checks: %w", err)
	}

	return nil
}
