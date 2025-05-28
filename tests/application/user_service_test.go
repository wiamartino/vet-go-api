package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	t.Run("Register - should successfully register a new user", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		service := application.NewUserService(mockRepo)

		user := domain.User{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "John Doe",
			Role:     "user",
		}

		mockRepo.On("Create", mock.AnythingOfType("domain.User")).Return(nil)

		// Act
		err := service.Register(user)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Register - should return error for invalid email", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		service := application.NewUserService(mockRepo)

		user := domain.User{
			Email:    "",
			Password: "password123",
			Name:     "John Doe",
		}

		// Act
		err := service.Register(user)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "email and password are required")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("Register - should return error for short password", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		service := application.NewUserService(mockRepo)

		user := domain.User{
			Email:    "test@example.com",
			Password: "123",
			Name:     "John Doe",
		}

		// Act
		err := service.Register(user)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "password must be at least 8 characters long")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("FindByEmail - should return user when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		service := application.NewUserService(mockRepo)

		expectedUser := domain.User{
			ID:    1,
			Email: "test@example.com",
			Name:  "John Doe",
		}

		mockRepo.On("FindByEmail", "test@example.com").Return(expectedUser, nil)

		// Act
		user, err := service.FindByEmail("test@example.com")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByEmail - should return error when user not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		service := application.NewUserService(mockRepo)

		mockRepo.On("FindByEmail", "notfound@example.com").Return(domain.User{}, errors.New("user not found"))

		// Act
		user, err := service.FindByEmail("notfound@example.com")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, domain.User{}, user)
		mockRepo.AssertExpectations(t)
	})
}
