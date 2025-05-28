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

func TestTreatmentController(t *testing.T) {
	t.Run("FindTreatments - should return all treatments", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.GET("/treatments", treatmentController.FindTreatments)

		expectedTreatments := []domain.Treatment{
			{TreatmentID: 1, Name: "Vaccination", Description: "Annual vaccination", Cost: 75.00},
			{TreatmentID: 2, Name: "Dental Cleaning", Description: "Professional dental cleaning", Cost: 150.00},
		}

		mockRepo.On("FindAll").Return(expectedTreatments, nil)

		// Act
		req, _ := http.NewRequest("GET", "/treatments", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var treatments []domain.Treatment
		err := json.Unmarshal(w.Body.Bytes(), &treatments)
		assert.NoError(t, err)
		assert.Len(t, treatments, 2)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindTreatments - should return error when service fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.GET("/treatments", treatmentController.FindTreatments)

		mockRepo.On("FindAll").Return([]domain.Treatment{}, errors.New("database error"))

		// Act
		req, _ := http.NewRequest("GET", "/treatments", nil)
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

	t.Run("CreateTreatment - should create treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.POST("/treatments", treatmentController.CreateTreatment)

		treatment := domain.Treatment{
			Name:        "Surgery",
			Description: "Minor surgical procedure",
			Cost:        500.00,
		}

		mockRepo.On("Create", mock.AnythingOfType("*domain.Treatment")).Return(nil)

		jsonData, _ := json.Marshal(treatment)

		// Act
		req, _ := http.NewRequest("POST", "/treatments", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var responseTreatment domain.Treatment
		err := json.Unmarshal(w.Body.Bytes(), &responseTreatment)
		assert.NoError(t, err)
		assert.Equal(t, treatment.Name, responseTreatment.Name)
		assert.Equal(t, treatment.Description, responseTreatment.Description)
		assert.Equal(t, treatment.Cost, responseTreatment.Cost)

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateTreatment - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.POST("/treatments", treatmentController.CreateTreatment)

		invalidJSON := `{"name": "Test Treatment", "cost": }`

		// Act
		req, _ := http.NewRequest("POST", "/treatments", bytes.NewBufferString(invalidJSON))
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

	t.Run("FindTreatment - should return treatment when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.GET("/treatments/:id", treatmentController.FindTreatment)

		expectedTreatment := domain.Treatment{
			TreatmentID: 1,
			Name:        "X-Ray",
			Description: "Diagnostic imaging",
			Cost:        125.00,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedTreatment, nil)

		// Act
		req, _ := http.NewRequest("GET", "/treatments/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var responseTreatment domain.Treatment
		err := json.Unmarshal(w.Body.Bytes(), &responseTreatment)
		assert.NoError(t, err)
		assert.Equal(t, expectedTreatment.TreatmentID, responseTreatment.TreatmentID)
		assert.Equal(t, expectedTreatment.Name, responseTreatment.Name)
		assert.Equal(t, expectedTreatment.Cost, responseTreatment.Cost)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindTreatment - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.GET("/treatments/:id", treatmentController.FindTreatment)

		// Act
		req, _ := http.NewRequest("GET", "/treatments/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})

	t.Run("FindTreatment - should return error when treatment not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.GET("/treatments/:id", treatmentController.FindTreatment)

		mockRepo.On("FindByID", uint(999)).Return(domain.Treatment{}, errors.New("treatment not found"))

		// Act
		req, _ := http.NewRequest("GET", "/treatments/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateTreatment - should update treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.PUT("/treatments/:id", treatmentController.UpdateTreatment)

		treatment := domain.Treatment{
			TreatmentID: 1,
			Name:        "Updated Treatment",
			Description: "Updated description",
			Cost:        250.00,
		}

		mockRepo.On("Update", mock.AnythingOfType("*domain.Treatment")).Return(nil)

		jsonData, _ := json.Marshal(treatment)

		// Act
		req, _ := http.NewRequest("PUT", "/treatments/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var responseTreatment domain.Treatment
		err := json.Unmarshal(w.Body.Bytes(), &responseTreatment)
		assert.NoError(t, err)
		assert.Equal(t, treatment.Name, responseTreatment.Name)
		assert.Equal(t, treatment.Description, responseTreatment.Description)
		assert.Equal(t, treatment.Cost, responseTreatment.Cost)

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteTreatment - should delete treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/treatments/:id", treatmentController.DeleteTreatment)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/treatments/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Treatment deleted successfully", response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteTreatment - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		treatmentService := application.NewTreatmentService(mockRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/treatments/:id", treatmentController.DeleteTreatment)

		// Act
		req, _ := http.NewRequest("DELETE", "/treatments/invalid", nil)
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
