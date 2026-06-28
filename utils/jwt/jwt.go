package jwt

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"go-vet/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var (
	jwtKey        []byte
	jwtIssuer     string
	jwtTimeout    time.Duration
	mu            sync.RWMutex
	isInitialized bool
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// InitFromConfig initializes JWT settings from the centralized config.
// Called automatically the first time a token is generated or validated.
func initFromConfig() {
	mu.Lock()
	defer mu.Unlock()

	if isInitialized {
		return
	}

	// Prefer config.AppConfig if it has been loaded (normal app startup)
	if config.AppConfig != nil && config.AppConfig.JWT.SecretKey != "" {
		jwtKey = []byte(config.AppConfig.JWT.SecretKey)
		jwtIssuer = config.AppConfig.JWT.Issuer
		jwtTimeout = config.AppConfig.JWT.Timeout
		isInitialized = true
		return
	}

	// Fallback: read env vars directly (for tests that don't load config)
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return // Will fail later with a clear error when GenerateToken is called
	}

	jwtKey = []byte(secretKey)

	jwtIssuer = os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		jwtIssuer = "vet-go-api"
		logrus.Info("JWT_ISSUER not set, using default: ", jwtIssuer)
	}

	jwtTimeoutStr := os.Getenv("JWT_TIMEOUT_HOURS")
	if jwtTimeoutStr == "" {
		jwtTimeout = 24 * time.Hour
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

	isInitialized = true
}

// ensureInitialized makes sure JWT is configured before use.
// Returns the current key, issuer, and timeout under a read lock.
func ensureInitialized() (key []byte, issuer string, timeout time.Duration, err error) {
	// Fast path: already initialized
	mu.RLock()
	if isInitialized {
		key, issuer, timeout = jwtKey, jwtIssuer, jwtTimeout
		mu.RUnlock()
		return
	}
	mu.RUnlock()

	// Slow path: initialize
	initFromConfig()

	mu.RLock()
	defer mu.RUnlock()
	if !isInitialized || len(jwtKey) == 0 {
		err = errors.New("JWT_SECRET_KEY must be set in environment or config")
		return
	}
	key, issuer, timeout = jwtKey, jwtIssuer, jwtTimeout
	return
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID uint, email string, role string) (string, error) {
	key, issuer, timeout, err := ensureInitialized()
	if err != nil {
		return "", err
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
	key, _, _, err := ensureInitialized()
	if err != nil {
		return nil, err
	}

	claims := &Claims{}
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

// ResetForTesting allows tests to re-initialize JWT config.
// This should only be called in test code.
func ResetForTesting() {
	mu.Lock()
	defer mu.Unlock()
	isInitialized = false
	jwtKey = nil
	jwtIssuer = ""
	jwtTimeout = 0
}
