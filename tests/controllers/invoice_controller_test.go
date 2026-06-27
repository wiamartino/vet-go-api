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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInvoiceController(t *testing.T) {
	t.Run("FindInvoices - should return all invoices", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.GET("/invoices", invoiceController.FindInvoices)

		expectedInvoices := []domain.Invoice{
			{InvoiceID: 1, ClientID: 1, AppointmentID: 1, Total: 250.50, Date: time.Now()},
			{InvoiceID: 2, ClientID: 2, AppointmentID: 2, Total: 150.00, Date: time.Now()},
		}

		mockRepo.On("FindAll").Return(expectedInvoices, nil)

		// Act
		req, _ := http.NewRequest("GET", "/invoices", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		t.Logf("invoice create raw body: %s", w.Body.String())
		assert.NoError(t, err)
		t.Logf("invoice create response: %#v", response)
		assert.Equal(t, "success", response["status"])

		data := response["data"].([]interface{})
		assert.Len(t, data, 2)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindInvoices - should return error when service fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.GET("/invoices", invoiceController.FindInvoices)

		mockRepo.On("FindAll").Return([]domain.Invoice{}, errors.New("database error"))

		// Act
		req, _ := http.NewRequest("GET", "/invoices", nil)
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

	t.Run("CreateInvoice - should return error for invalid payload", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.POST("/invoices", invoiceController.CreateInvoice)

		// Use explicit JSON; current handler validation may reject
		jsonBody := `{"client_id":1,"appointment_id":1}`

		// Act
		req, _ := http.NewRequest("POST", "/invoices", bytes.NewBufferString(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		// On error, no data payload is expected
	})

	t.Run("CreateInvoice - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.POST("/invoices", invoiceController.CreateInvoice)

		invalidJSON := `{"client_id": 1, "total": }`

		// Act
		req, _ := http.NewRequest("POST", "/invoices", bytes.NewBufferString(invalidJSON))
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

	t.Run("FindInvoice - should return invoice when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.GET("/invoices/:id", invoiceController.FindInvoice)

		expectedInvoice := domain.Invoice{
			InvoiceID:     1,
			ClientID:      1,
			AppointmentID: 1,
			Total:         175.00,
			Date:          time.Now(),
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedInvoice, nil)

		// Act
		req, _ := http.NewRequest("GET", "/invoices/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		data := response["data"].(map[string]interface{})
		assert.Equal(t, float64(expectedInvoice.InvoiceID), data["invoice_id"])
		assert.Equal(t, float64(expectedInvoice.ClientID), data["client_id"])
		assert.Equal(t, expectedInvoice.Total, data["total"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindInvoice - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.GET("/invoices/:id", invoiceController.FindInvoice)

		// Act
		req, _ := http.NewRequest("GET", "/invoices/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
	})

	t.Run("FindInvoice - should return error when invoice not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.GET("/invoices/:id", invoiceController.FindInvoice)

		mockRepo.On("FindByID", uint(999)).Return(domain.Invoice{}, errors.New("invoice not found"))

		// Act
		req, _ := http.NewRequest("GET", "/invoices/999", nil)
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

	t.Run("UpdateInvoice - should return error for invalid payload", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.PUT("/invoices/:id", invoiceController.UpdateInvoice)

		invoice := domain.Invoice{
			InvoiceID:     1,
			ClientID:      1,
			AppointmentID: 1,
			Total:         -5.00,
			Date:          time.Now(),
		}

		// For invalid payload scenario, controller may reject before calling repo

		jsonData, _ := json.Marshal(invoice)

		// Act
		req, _ := http.NewRequest("PUT", "/invoices/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		// No data expected on error

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteInvoice - should delete invoice successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/invoices/:id", invoiceController.DeleteInvoice)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/invoices/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Invoice deleted successfully", response["data"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteInvoice - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		invoiceService := application.NewInvoiceService(mockRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/invoices/:id", invoiceController.DeleteInvoice)

		// Act
		req, _ := http.NewRequest("DELETE", "/invoices/invalid", nil)
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
