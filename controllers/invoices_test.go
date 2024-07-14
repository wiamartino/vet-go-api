package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-vet/config"
	"go-vet/controllers"
	"go-vet/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInvoiceTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Invoice{}, &models.Client{}, &models.Appointment{})
	config.DB = db
	return db
}

func setupInvoiceRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestUpdateInvoiceSuccess(t *testing.T) {
	db := setupInvoiceTestDB()
	defer db.Migrator().DropTable(&models.Invoice{}, &models.Client{}, &models.Appointment{})

	client := models.Client{FirstName: "John", LastName: "Doe", Email: "john@example.com"}
	db.Create(&client)
	invoice := models.Invoice{ClientID: client.ClientID, Total: 100}
	db.Create(&invoice)

	router := setupInvoiceRouter()
	router.PUT("/invoices/:id", controllers.UpdateInvoice)

	updatedInvoice := models.Invoice{Total: 200}
	jsonValue, _ := json.Marshal(updatedInvoice)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/invoices/%d", invoice.ClientID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseInvoice models.Invoice
	json.Unmarshal(w.Body.Bytes(), &responseInvoice)
	assert.Equal(t, updatedInvoice.Total, responseInvoice.Total)
}

func TestUpdateInvoiceNotFound(t *testing.T) {
	setupInvoiceTestDB()
	defer config.DB.Migrator().DropTable(&models.Invoice{}, &models.Client{}, &models.Appointment{})

	router := setupInvoiceRouter()
	router.PUT("/invoices/:id", controllers.UpdateInvoice)

	req, _ := http.NewRequest("PUT", "/invoices/999", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Invoice not found")
}

func TestUpdateInvoiceBadRequest(t *testing.T) {
	db := setupInvoiceTestDB()
	defer db.Migrator().DropTable(&models.Invoice{}, &models.Client{}, &models.Appointment{})

	client := models.Client{FirstName: "John", LastName: "Doe", Email: "john@example.com"}
	db.Create(&client)
	invoice := models.Invoice{ClientID: client.ClientID, Total: 100}
	db.Create(&invoice)

	router := setupInvoiceRouter()
	router.PUT("/invoices/:id", controllers.UpdateInvoice)

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/invoices/%d", invoice.ClientID), bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
