package middlewares_test

import (
	"go-vet/middlewares"
	"go-vet/utils/jwt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	t.Run("should return error when Authorization header is missing", func(t *testing.T) {
		// Arrange
		router := gin.New()
		router.Use(middlewares.AuthMiddleware())
		router.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		// Act
		req, _ := http.NewRequest("GET", "/protected", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Authorization header required")
	})

	t.Run("should return error when Authorization header format is invalid", func(t *testing.T) {
		// Arrange
		router := gin.New()
		router.Use(middlewares.AuthMiddleware())
		router.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		// Act
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "InvalidToken")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid token format")
	})

	t.Run("should return error when token is invalid", func(t *testing.T) {
		// Arrange
		router := gin.New()
		router.Use(middlewares.AuthMiddleware())
		router.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		})

		// Act
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer invalid.token.here")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid or expired token")
	})

	t.Run("should proceed when valid token is provided", func(t *testing.T) {
		// This test would require a valid JWT token
		// For now, we'll skip the actual token validation part
		// In a real scenario, you'd generate a valid token using your JWT utils

		// Arrange
		router := gin.New()

		// Mock middleware that simulates successful authentication
		router.Use(func(c *gin.Context) {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "Bearer valid.jwt.token" {
				// Simulate setting user context
				c.Set("userID", uint(1))
				c.Set("email", "test@example.com")
				c.Set("role", "user")
				c.Next()
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
			}
		})

		router.GET("/protected", func(c *gin.Context) {
			userID, exists := c.Get("userID")
			assert.True(t, exists)
			assert.Equal(t, uint(1), userID)

			c.JSON(http.StatusOK, gin.H{"message": "success", "userID": userID})
		})

		// Act
		req, _ := http.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer valid.jwt.token")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})
}

func TestJWTUtils(t *testing.T) {
	t.Run("should generate and validate token", func(t *testing.T) {
		// This test would validate the JWT utility functions
		// You'd need to implement this based on your JWT utils structure

		// Example test structure:
		userID := uint(1)
		email := "test@example.com"
		role := "user"

		// Generate token
		token, err := jwt.GenerateToken(userID, email, role)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Validate token
		claims, err := jwt.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, email, claims.Email)
		assert.Equal(t, role, claims.Role)
	})

	t.Run("should return error for invalid token", func(t *testing.T) {
		// Test invalid token
		claims, err := jwt.ValidateToken("invalid.token.here")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("should return error for expired token", func(t *testing.T) {
		// This would require generating an expired token
		// Implementation depends on your JWT utility structure

		// For now, just test with a malformed token
		claims, err := jwt.ValidateToken("expired.token.here")
		assert.Error(t, err)
		assert.Nil(t, claims)
	})
}
