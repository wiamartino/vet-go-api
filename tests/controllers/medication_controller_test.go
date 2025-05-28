package controllers_test

import (
	"bytes"
	"encoding/json"
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

func TestMedicationController(t *testing.T) {
	t.Run("FindMedications - should return all medications", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.GET("/medications", medicationController.FindMedications)

		medications := []domain.Medication{
			{Name: "Aspirin", Description: "Pain reliever", Price: 10.50},
			{Name: "Penicillin", Description: "Antibiotic", Price: 25.00},
		}

		mockRepo.On("FindAll").Return(medications, nil)

		// Act
		req, _ := http.NewRequest("GET", "/medications", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.NotNil(t, response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMedication - should create medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.POST("/medications", medicationController.CreateMedication)

		medication := domain.Medication{
			Name:        "New Medication",
			Description: "Test medication",
			Price:       15.00,
		}

		jsonData, _ := json.Marshal(medication)
		mockRepo.On("Create", mock.AnythingOfType("*domain.Medication")).Return(nil)

		// Act
		req, _ := http.NewRequest("POST", "/medications", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMedication - should handle invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.POST("/medications", medicationController.CreateMedication)

		invalidJSON := []byte(`{"name": "Invalid"`)

		// Act
		req, _ := http.NewRequest("POST", "/medications", bytes.NewBuffer(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("FindMedication - should return medication by ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.GET("/medications/:id", medicationController.FindMedication)

		medication := domain.Medication{
			ID:          1,
			Name:        "Aspirin",
			Description: "Pain reliever",
			Price:       10.50,
		}

		mockRepo.On("FindByID", uint(1)).Return(medication, nil)

		// Act
		req, _ := http.NewRequest("GET", "/medications/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])
		assert.NotNil(t, response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindMedication - should handle invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.GET("/medications/:id", medicationController.FindMedication)

		// Act
		req, _ := http.NewRequest("GET", "/medications/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("UpdateMedication - should update medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.PUT("/medications/:id", medicationController.UpdateMedication)

		medication := domain.Medication{
			ID:          1,
			Name:        "Updated Medication",
			Description: "Updated description",
			Price:       20.00,
		}

		mockRepo.On("Update", &medication).Return(nil)

		jsonData, _ := json.Marshal(medication)

		// Act
		req, _ := http.NewRequest("PUT", "/medications/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateMedication - should handle invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.PUT("/medications/:id", medicationController.UpdateMedication)

		invalidJSON := []byte(`{"name": "Invalid"`)

		// Act
		req, _ := http.NewRequest("PUT", "/medications/1", bytes.NewBuffer(invalidJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("DeleteMedication - should delete medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/medications/:id", medicationController.DeleteMedication)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/medications/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "success", response["status"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteMedication - should handle invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/medications/:id", medicationController.DeleteMedication)

		// Act
		req, _ := http.NewRequest("DELETE", "/medications/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
