package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	jwtKey      []byte
	jwtIssuer   string
	jwtTimeout  time.Duration
	initialized bool
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func init() {
	// Try to initialize, but don't fail if JWT_SECRET_KEY is not set
	// This allows tests to run without having to set environment variables globally
	initializeJWT()
}

func initializeJWT() {
	if initialized {
		return
	}

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warning("Error loading .env file, using default values")
	}

	// Get JWT secret key
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		// Don't fail immediately - this allows tests to set environment variables
		return
	}
	jwtKey = []byte(jwtSecretKey)

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

	initialized = true
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID uint, email string, role string) (string, error) {
	// Ensure JWT is initialized
	initializeJWT()

	if len(jwtKey) == 0 {
		jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
		if jwtSecretKey == "" {
			return "", errors.New("JWT_SECRET_KEY must be set in environment")
		}
		jwtKey = []byte(jwtSecretKey)

		// Also set other defaults if not initialized
		if jwtIssuer == "" {
			jwtIssuer = os.Getenv("JWT_ISSUER")
			if jwtIssuer == "" {
				jwtIssuer = "vet-go-api"
			}
		}

		if jwtTimeout == 0 {
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
		}
	}

	expirationTime := time.Now().Add(jwtTimeout)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwtIssuer,
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check expiration
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return nil, errors.New("token expired")
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
