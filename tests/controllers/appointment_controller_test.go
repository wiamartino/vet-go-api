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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAppointmentController(t *testing.T) {
	t.Run("FindAppointments - should return all appointments", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.GET("/appointments", appointmentController.FindAppointments)

		expectedAppointments := []domain.Appointment{
			{AppointmentID: 1, PetID: 1, VeterinarianID: 1, Date: time.Now().Add(24 * time.Hour), Time: time.Now().Add(24 * time.Hour), ReasonForAppointment: "Checkup"},
			{AppointmentID: 2, PetID: 2, VeterinarianID: 1, Date: time.Now().Add(48 * time.Hour), Time: time.Now().Add(48 * time.Hour), ReasonForAppointment: "Vaccination"},
		}

		mockRepo.On("FindAll").Return(expectedAppointments, nil)

		// Act
		req, _ := http.NewRequest("GET", "/appointments", nil)
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

	t.Run("FindAppointments - should return error when service fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.GET("/appointments", appointmentController.FindAppointments)

		mockRepo.On("FindAll").Return([]domain.Appointment{}, errors.New("database error"))

		// Act
		req, _ := http.NewRequest("GET", "/appointments", nil)
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

	t.Run("CreateAppointment - should create appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.POST("/appointments", appointmentController.CreateAppointment)

		appointment := domain.Appointment{
			PetID:                1,
			VeterinarianID:       1,
			Date:                 time.Now().Add(72 * time.Hour),
			Time:                 time.Now().Add(72 * time.Hour),
			ReasonForAppointment: "Emergency checkup",
		}

		mockRepo.On("Create", mock.AnythingOfType("*domain.Appointment")).Return(nil)

		jsonData, _ := json.Marshal(appointment)

		// Act
		req, _ := http.NewRequest("POST", "/appointments", bytes.NewBuffer(jsonData))
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

	t.Run("CreateAppointment - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.POST("/appointments", appointmentController.CreateAppointment)

		invalidJSON := `{"pet_id": 1, "veterinarian_id": }`

		// Act
		req, _ := http.NewRequest("POST", "/appointments", bytes.NewBufferString(invalidJSON))
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

	t.Run("FindAppointment - should return appointment when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.GET("/appointments/:id", appointmentController.FindAppointment)

		expectedAppointment := domain.Appointment{
			AppointmentID:        1,
			PetID:                1,
			VeterinarianID:       1,
			Date:                 time.Now().Add(24 * time.Hour),
			Time:                 time.Now().Add(24 * time.Hour),
			ReasonForAppointment: "Regular checkup",
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedAppointment, nil)

		// Act
		req, _ := http.NewRequest("GET", "/appointments/1", nil)
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

	t.Run("FindAppointment - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.GET("/appointments/:id", appointmentController.FindAppointment)

		// Act
		req, _ := http.NewRequest("GET", "/appointments/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})

	t.Run("FindAppointment - should return error when appointment not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.GET("/appointments/:id", appointmentController.FindAppointment)

		mockRepo.On("FindByID", uint(999)).Return(domain.Appointment{}, errors.New("appointment not found"))

		// Act
		req, _ := http.NewRequest("GET", "/appointments/999", nil)
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

	t.Run("UpdateAppointment - should update appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.PUT("/appointments/:id", appointmentController.UpdateAppointment)

		appointment := domain.Appointment{
			AppointmentID:        1,
			PetID:                1,
			VeterinarianID:       2,
			Date:                 time.Now().Add(96 * time.Hour),
			Time:                 time.Now().Add(96 * time.Hour),
			ReasonForAppointment: "Updated appointment",
		}

		mockRepo.On("Update", mock.AnythingOfType("*domain.Appointment")).Return(nil)

		jsonData, _ := json.Marshal(appointment)

		// Act
		req, _ := http.NewRequest("PUT", "/appointments/1", bytes.NewBuffer(jsonData))
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

	t.Run("DeleteAppointment - should delete appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.DELETE("/appointments/:id", appointmentController.DeleteAppointment)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/appointments/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Appointment deleted", response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteAppointment - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		appointmentService := application.NewAppointmentService(mockRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)

		router := setupTestRouter()
		router.DELETE("/appointments/:id", appointmentController.DeleteAppointment)

		// Act
		req, _ := http.NewRequest("DELETE", "/appointments/invalid", nil)
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
