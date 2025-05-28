package controllers_test

import (
	"bytes"
	"encoding/json"
	"go-vet/application"
	"go-vet/controllers"
	"go-vet/domain"
	"go-vet/tests/mocks"
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

		router := setupTestRouter()
		router.GET("/medications", medicationController.FindMedications)

		medications := []domain.Medication{
			{ID: 1, Name: "Aspirin", Description: "Pain reliever", Price: 10.50},
			{ID: 2, Name: "Penicillin", Description: "Antibiotic", Price: 25.00},
		}

		mockRepo.On("FindAll").Return(medications, nil)

		// Act
		req, _ := http.NewRequest("GET", "/medications", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMedication - should create medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		medicationService := application.NewMedicationService(mockRepo)
		medicationController := controllers.NewMedicationController(medicationService)

		router := setupTestRouter()
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
		mockRepo.AssertExpectations(t)
	})
}
