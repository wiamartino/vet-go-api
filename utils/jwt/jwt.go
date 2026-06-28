package jwt

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	jwtKey        []byte
	jwtIssuer     string
	jwtTimeout    time.Duration
	once          sync.Once
	mu            sync.RWMutex
	isInitialized bool
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func init() {
	// Try to initialize, but don't fail if JWT_SECRET_KEY is not set
	// This allows tests to run without having to set environment variables globally
	initializeJWT()
}

func initializeJWT() {
	once.Do(func() {
		loadJWTConfig()
	})
}

// loadJWTConfig loads JWT configuration from environment variables
// This is separated from initializeJWT to allow testing
func loadJWTConfig() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("Error loading .env file, using default values")
	}

	// Get JWT secret key
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		// Don't fail immediately - allow lazy initialization when key is set
		mu.Lock()
		isInitialized = false
		mu.Unlock()
		return
	}

	mu.Lock()
	defer mu.Unlock()

	jwtKey = []byte(jwtSecretKey)
	isInitialized = true

	// Get JWT issuer (optional)
	jwtIssuer = os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		jwtIssuer = "vet-go-api"
		logrus.Info("JWT_ISSUER not set, using default: ", jwtIssuer)
	}

	// Get JWT timeout (optional)
	jwtTimeoutStr := os.Getenv("JWT_TIMEOUT_HOURS")
	if jwtTimeoutStr == "" {
		jwtTimeout = 24 * time.Hour // Default to 24 hours
		logrus.Info("JWT_TIMEOUT_HOURS not set, using default: 24 hours")
	} else {
		var timeoutHours int
		_, err := fmt.Sscanf(jwtTimeoutStr, "%d", &timeoutHours)
		if err != nil {
			jwtTimeout = 24 * time.Hour
			logrus.Warning("Invalid JWT_TIMEOUT_HOURS, using default: 24 hours")
		} else {
			jwtTimeout = time.Duration(timeoutHours) * time.Hour
		}
	}
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID uint, email string, role string) (string, error) {
	// Ensure JWT is initialized
	initializeJWT()

	// Get all config values under a single read lock
	mu.RLock()
	key := jwtKey
	issuer := jwtIssuer
	timeout := jwtTimeout
	initialized := isInitialized
	mu.RUnlock()

	// If not initialized, try one more time with write lock
	if !initialized {
		mu.Lock()
		// Check again after acquiring write lock
		if !isInitialized {
			jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
			if jwtSecretKey != "" {
				jwtKey = []byte(jwtSecretKey)

				jwtIssuer = os.Getenv("JWT_ISSUER")
				if jwtIssuer == "" {
					jwtIssuer = "vet-go-api"
				}

				jwtTimeoutStr := os.Getenv("JWT_TIMEOUT_HOURS")
				if jwtTimeoutStr == "" {
					jwtTimeout = 24 * time.Hour
				} else {
					var timeoutHours int
					_, err := fmt.Sscanf(jwtTimeoutStr, "%d", &timeoutHours)
					if err != nil {
						jwtTimeout = 24 * time.Hour
					} else {
						jwtTimeout = time.Duration(timeoutHours) * time.Hour
					}
				}

				isInitialized = true
				key = jwtKey
				issuer = jwtIssuer
				timeout = jwtTimeout
			}
		} else {
			// Another goroutine initialized it
			key = jwtKey
			issuer = jwtIssuer
			timeout = jwtTimeout
		}
		mu.Unlock()
	}

	// Fail if JWT key is not initialized
	if len(key) == 0 {
		return "", errors.New("JWT_SECRET_KEY must be set in environment")
	}

	now := time.Now()
	expirationTime := now.Add(timeout)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    issuer,
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	mu.RLock()
	key := jwtKey
	mu.RUnlock()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Note: Expiration is already validated by jwt.ParseWithClaims
	// The library automatically checks ExpiresAt and returns an error if expired

	return claims, nil
}

// RefreshToken generates a new token with updated expiration time
func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Generate a new token with the same user information but updated expiry
	return GenerateToken(claims.UserID, claims.Email, claims.Role)
}
