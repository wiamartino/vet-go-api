package utils_test

import (
	"os"
	"testing"
	"time"

	jwtUtils "go-vet/utils/jwt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	// Set required environment variables before any JWT package imports
	os.Setenv("JWT_SECRET_KEY", "test-secret-key-for-testing-only")
	os.Setenv("JWT_ISSUER", "vet-go-api")
	os.Setenv("JWT_TIMEOUT_HOURS", "24")
}

func setupJWTTestEnv(t *testing.T) {
	// Environment variables are already set in init()
}

func TestJWTUtils(t *testing.T) {
	setupJWTTestEnv(t)

	t.Run("GenerateToken - should create valid token", func(t *testing.T) {
		// Arrange
		userID := uint(123)
		email := "test@example.com"
		role := "veterinarian"

		// Act
		token, err := jwtUtils.GenerateToken(userID, email, role)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		assert.Contains(t, token, ".")

		// JWT tokens should have 3 parts separated by dots
		parts := splitToken(token)
		assert.Len(t, parts, 3)
	})

	t.Run("GenerateToken - should create different tokens for different users", func(t *testing.T) {
		// Arrange
		user1Token, err1 := jwtUtils.GenerateToken(1, "user1@test.com", "user")
		user2Token, err2 := jwtUtils.GenerateToken(2, "user2@test.com", "admin")

		// Assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, user1Token, user2Token)
	})

	t.Run("ValidateToken - should validate correct token", func(t *testing.T) {
		// Arrange
		userID := uint(456)
		email := "validate@example.com"
		role := "admin"

		token, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		// Act
		claims, err := jwtUtils.ValidateToken(token)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, claims)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, email, claims.Email)
		assert.Equal(t, role, claims.Role)
		assert.Equal(t, "vet-go-api", claims.Issuer)
		assert.Equal(t, email, claims.Subject)
	})

	t.Run("ValidateToken - should reject malformed token", func(t *testing.T) {
		// Act
		claims, err := jwtUtils.ValidateToken("invalid.token.format")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, claims)
		assert.Contains(t, err.Error(), "invalid token")
	})

	t.Run("ValidateToken - should reject empty token", func(t *testing.T) {
		// Act
		claims, err := jwtUtils.ValidateToken("")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("ValidateToken - should reject token with invalid signature", func(t *testing.T) {
		// This is a token signed with a different secret
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20iLCJyb2xlIjoidXNlciIsImV4cCI6OTk5OTk5OTk5OSwiaWF0IjoxNjAwMDAwMDAwLCJpc3MiOiJ2ZXQtZ28tdGVzdCIsInN1YiI6InRlc3RAZXhhbXBsZS5jb20ifQ.invalid_signature"

		// Act
		claims, err := jwtUtils.ValidateToken(invalidToken)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, claims)
	})

	t.Run("RefreshToken - should create new token with same claims", func(t *testing.T) {
		// Arrange
		userID := uint(789)
		email := "refresh@example.com"
		role := "user"

		originalToken, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		// Wait a short time to ensure different timestamps
		time.Sleep(10 * time.Millisecond)

		// Act
		time.Sleep(time.Second) // Ensure different timestamp
		refreshedToken, err := jwtUtils.RefreshToken(originalToken)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, refreshedToken)
		// Note: Tokens might be the same if generated within the same second
		// The important part is that the refresh process works without error

		// Validate the refreshed token has the same user information
		claims, err := jwtUtils.ValidateToken(refreshedToken)
		require.NoError(t, err)
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, email, claims.Email)
		assert.Equal(t, role, claims.Role)
	})

	t.Run("RefreshToken - should reject invalid token", func(t *testing.T) {
		// Act
		refreshedToken, err := jwtUtils.RefreshToken("invalid.token.here")

		// Assert
		assert.Error(t, err)
		assert.Empty(t, refreshedToken)
	})

	t.Run("Token expiration - should handle token expiration correctly", func(t *testing.T) {
		// Arrange
		userID := uint(999)
		email := "expiry@example.com"
		role := "test"

		token, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		// Validate immediately (should work)
		claims, err := jwtUtils.ValidateToken(token)
		assert.NoError(t, err)
		assert.NotNil(t, claims)

		// Check that the expiration time is in the future
		assert.True(t, time.Unix(claims.ExpiresAt, 0).After(time.Now()))
		assert.True(t, time.Unix(claims.IssuedAt, 0).Before(time.Now().Add(time.Second)))
	})

	t.Run("Token claims - should include all required fields", func(t *testing.T) {
		// Arrange
		userID := uint(555)
		email := "claims@example.com"
		role := "moderator"

		token, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		// Act
		claims, err := jwtUtils.ValidateToken(token)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, claims)

		// Check all custom claims
		assert.Equal(t, userID, claims.UserID)
		assert.Equal(t, email, claims.Email)
		assert.Equal(t, role, claims.Role)

		// Check standard claims
		assert.NotZero(t, claims.ExpiresAt)
		assert.NotZero(t, claims.IssuedAt)
		assert.Equal(t, "vet-go-api", claims.Issuer)
		assert.Equal(t, email, claims.Subject)
	})

	t.Run("Edge cases - should handle special characters in email", func(t *testing.T) {
		// Arrange
		userID := uint(777)
		email := "test+special@example-domain.co.uk"
		role := "user"

		// Act
		token, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		claims, err := jwtUtils.ValidateToken(token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, email, claims.Email)
		assert.Equal(t, email, claims.Subject)
	})

	t.Run("Edge cases - should handle maximum uint user ID", func(t *testing.T) {
		// Arrange
		userID := uint(^uint(0) >> 1) // Maximum uint value
		email := "maxuser@example.com"
		role := "admin"

		// Act
		token, err := jwtUtils.GenerateToken(userID, email, role)
		require.NoError(t, err)

		claims, err := jwtUtils.ValidateToken(token)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, userID, claims.UserID)
	})
}

// Helper function to split JWT token into parts
func splitToken(token string) []string {
	parts := []string{}
	current := ""

	for _, char := range token {
		if char == '.' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(char)
		}
	}

	if current != "" {
		parts = append(parts, current)
	}

	return parts
}
