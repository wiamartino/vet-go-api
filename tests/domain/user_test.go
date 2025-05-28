package domain_test

import (
	"go-vet/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	t.Run("should create user with valid fields", func(t *testing.T) {
		user := domain.User{
			ID:        1,
			Email:     "test@example.com",
			Password:  "hashedpassword123",
			Name:      "John Doe",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		assert.Equal(t, uint(1), user.ID)
		assert.Equal(t, "test@example.com", user.Email)
		assert.Equal(t, "hashedpassword123", user.Password)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "user", user.Role)
		assert.False(t, user.CreatedAt.IsZero())
		assert.False(t, user.UpdatedAt.IsZero())
	})

	t.Run("should handle different user roles", func(t *testing.T) {
		roles := []string{"admin", "user", "veterinarian"}

		for _, role := range roles {
			user := domain.User{
				Email: "test@example.com",
				Name:  "Test User",
				Role:  role,
			}
			assert.Equal(t, role, user.Role)
		}
	})
}
