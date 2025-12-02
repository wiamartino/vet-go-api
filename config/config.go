package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Logging  LoggingConfig
	Security SecurityConfig
}

type ServerConfig struct {
	Port        string
	Environment string
	Host        string
}

type DatabaseConfig struct {
	Host            string
	User            string
	Password        string
	Name            string
	Port            string
	SSLMode         string
	TimeZone        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type JWTConfig struct {
	SecretKey string
	Issuer    string
	Timeout   time.Duration
}

type LoggingConfig struct {
	Level  string
	Format string
}

type SecurityConfig struct {
	PasswordMinLength  int
	BCryptCost         int
	RateLimitPerMinute int
	CORSOrigins        []string
}

var AppConfig *Config

func LoadConfig() error {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		logrus.Warning("No .env file found, using environment variables")
	}

	config := &Config{
		Server: ServerConfig{
			Port:        getEnv("SERVER_PORT", "8080"),
			Environment: getEnv("GO_ENV", "development"),
			Host:        getEnv("SERVER_HOST", "localhost"),
		},
		Database: DatabaseConfig{
			Host:            getEnv("DB_HOST", "localhost"),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", ""),
			Name:            getEnv("DB_NAME", "vet_go"),
			Port:            getEnv("DB_PORT", "5432"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			TimeZone:        getEnv("DB_TIMEZONE", "UTC"),
			MaxOpenConns:    getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetime: time.Duration(getEnvAsInt("DB_CONN_MAX_LIFETIME_MINUTES", 30)) * time.Minute,
		},
		JWT: JWTConfig{
			SecretKey: getEnv("JWT_SECRET_KEY", ""),
			Issuer:    getEnv("JWT_ISSUER", "vet-go-api"),
			Timeout:   time.Duration(getEnvAsInt("JWT_TIMEOUT_HOURS", 24)) * time.Hour,
		},
		Logging: LoggingConfig{
			Level:  getEnv("LOG_LEVEL", "info"),
			Format: getEnv("LOG_FORMAT", "text"),
		},
		Security: SecurityConfig{
			PasswordMinLength:  getEnvAsInt("PASSWORD_MIN_LENGTH", 8),
			BCryptCost:         getEnvAsInt("BCRYPT_COST", 12),
			RateLimitPerMinute: getEnvAsInt("RATE_LIMIT_PER_MINUTE", 60),
			CORSOrigins:        []string{getEnv("CORS_ORIGINS", "*")},
		},
	}

	// Validate required configuration
	if config.JWT.SecretKey == "" {
		return fmt.Errorf("JWT_SECRET_KEY is required")
	}

	if config.Database.Password == "" {
		logrus.Warning("DB_PASSWORD is empty - this may cause connection issues")
	}

	AppConfig = config
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		logrus.Warningf("Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

func IsDevelopment() bool {
	return AppConfig.Server.Environment == "development"
}

func IsProduction() bool {
	return AppConfig.Server.Environment == "production"
}
