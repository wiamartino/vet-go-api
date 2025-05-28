package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-vet/application"
	"go-vet/controllers"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"go-vet/tests/testhelpers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthController(t *testing.T) {
	t.Run("Register - should register user successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		userService := application.NewUserService(mockRepo)
		authController := controllers.NewAuthController(userService)

		router := testhelpers.SetupTestRouter()
		router.POST("/auth/register", authController.Register)

		user := domain.User{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "John Doe",
			Role:     "user",
		}

		jsonData, _ := json.Marshal(user)

		// Mock the repository calls
		mockRepo.On("FindByEmail", "test@example.com").Return(domain.User{}, errors.New("user not found"))
		mockRepo.On("Create", mock.AnythingOfType("domain.User")).Return(nil)

		// Act
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("Register - should return error for duplicate email", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		userService := application.NewUserService(mockRepo)
		authController := controllers.NewAuthController(userService)

		router := testhelpers.SetupTestRouter()
		router.POST("/auth/register", authController.Register)

		user := domain.User{
			Email:    "test@example.com",
			Password: "password123",
			Name:     "John Doe",
		}

		jsonData, _ := json.Marshal(user)

		// Mock finding existing user
		existingUser := domain.User{ID: 1, Email: "test@example.com"}
		mockRepo.On("FindByEmail", "test@example.com").Return(existingUser, nil)

		// Act
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusConflict, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
		assert.Contains(t, response["error"], "Email is already registered")

		mockRepo.AssertExpectations(t)
	})

	t.Run("Register - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockUserRepository)
		userService := application.NewUserService(mockRepo)
		authController := controllers.NewAuthController(userService)

		router := testhelpers.SetupTestRouter()
		router.POST("/auth/register", authController.Register)

		invalidJSON := `{"email": "test@example.com", "password": }`

		// Act
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid input data", response["error"])
	})
}
