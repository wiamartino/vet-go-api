package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-vet/application"
	"go-vet/controllers"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVeterinarianController(t *testing.T) {
	t.Run("FindVeterinarians - should return all veterinarians", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.GET("/veterinarians", veterinarianController.FindVeterinarians)

		expectedVeterinarians := []domain.Veterinarian{
			{VeterinarianID: 1, FirstName: "Dr. John", LastName: "Smith", Specialty: "General Practice", Phone: "123-456-7890", Email: "dr.smith@vet.com"},
			{VeterinarianID: 2, FirstName: "Dr. Jane", LastName: "Johnson", Specialty: "Surgery", Phone: "123-456-7891", Email: "dr.johnson@vet.com"},
		}

		mockRepo.On("FindAll").Return(expectedVeterinarians, nil)

		// Act
		req, _ := http.NewRequest("GET", "/veterinarians", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		data := response["data"].([]interface{})
		assert.Len(t, data, 2)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindVeterinarians - should return error when service fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.GET("/veterinarians", veterinarianController.FindVeterinarians)

		mockRepo.On("FindAll").Return([]domain.Veterinarian{}, errors.New("database error"))

		// Act
		req, _ := http.NewRequest("GET", "/veterinarians", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateVeterinarian - should create veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.POST("/veterinarians", veterinarianController.CreateVeterinarian)

		veterinarian := domain.Veterinarian{
			FirstName: "Dr. Robert",
			LastName:  "Brown",
			Specialty: "Orthopedic Surgery",
			Phone:     "+1234567890",
			Email:     "dr.brown@vetclinic.com",
		}

		mockRepo.On("Create", &veterinarian).Return(nil)

		jsonData, _ := json.Marshal(veterinarian)

		// Act
		req, _ := http.NewRequest("POST", "/veterinarians", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateVeterinarian - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.POST("/veterinarians", veterinarianController.CreateVeterinarian)

		invalidJSON := `{"name": "Dr. Test", "specialty": }`

		// Act
		req, _ := http.NewRequest("POST", "/veterinarians", bytes.NewBufferString(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})

	t.Run("FindVeterinarian - should return veterinarian when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.GET("/veterinarians/:id", veterinarianController.FindVeterinarian)

		expectedVeterinarian := domain.Veterinarian{
			VeterinarianID: 1,
			FirstName:      "Dr. Susan",
			LastName:       "Wilson",
			Specialty:      "Cardiology",
			Phone:          "+1987654321",
			Email:          "dr.wilson@vetclinic.com",
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedVeterinarian, nil)

		// Act
		req, _ := http.NewRequest("GET", "/veterinarians/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindVeterinarian - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.GET("/veterinarians/:id", veterinarianController.FindVeterinarian)

		// Act
		req, _ := http.NewRequest("GET", "/veterinarians/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})

	t.Run("FindVeterinarian - should return error when veterinarian not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.GET("/veterinarians/:id", veterinarianController.FindVeterinarian)

		mockRepo.On("FindByID", uint(999)).Return(domain.Veterinarian{}, errors.New("veterinarian not found"))

		// Act
		req, _ := http.NewRequest("GET", "/veterinarians/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateVeterinarian - should update veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.PUT("/veterinarians/:id", veterinarianController.UpdateVeterinarian)

		veterinarian := domain.Veterinarian{
			VeterinarianID: 1,
			FirstName:      "Dr. Updated",
			LastName:       "Name",
			Specialty:      "Updated Specialty",
			Phone:          "+1111111111",
			Email:          "updated@vetclinic.com",
		}

		mockRepo.On("Update", &veterinarian).Return(nil)

		jsonData, _ := json.Marshal(veterinarian)

		// Act
		req, _ := http.NewRequest("PUT", "/veterinarians/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteVeterinarian - should delete veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.DELETE("/veterinarians/:id", veterinarianController.DeleteVeterinarian)

		// Mock the GetVeterinarianByID call that happens before deletion
		mockRepo.On("FindByID", uint(1)).Return(domain.Veterinarian{VeterinarianID: 1}, nil)
		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/veterinarians/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Veterinarian deleted", response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteVeterinarian - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		veterinarianService := application.NewVeterinarianService(mockRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)

		router := setupTestRouter()
		router.DELETE("/veterinarians/:id", veterinarianController.DeleteVeterinarian)

		// Act
		req, _ := http.NewRequest("DELETE", "/veterinarians/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})
}
