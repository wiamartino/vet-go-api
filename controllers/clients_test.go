package controllers_test

import (
	"bytes"
	"encoding/json"
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

func setupClientTestDB() {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Client{})
	config.DB = db
}

func TestFindClients(t *testing.T) {
	setupClientTestDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controllers.FindClients(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]")
}

func TestCreateClient(t *testing.T) {
	setupClientTestDB()
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	client := models.Client{FirstName: "John Doe", Email: "john@example.com"}
	jsonValue, _ := json.Marshal(client)
	c.Request, _ = http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	controllers.CreateClient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseClient models.Client
	json.Unmarshal(w.Body.Bytes(), &responseClient)
	assert.Equal(t, client.FirstName, responseClient.FirstName)
	assert.Equal(t, client.Email, responseClient.Email)
}

func TestFindClient(t *testing.T) {
	setupClientTestDB()
	client := models.Client{FirstName: "John Doe", Email: "john@example.com"}
	config.DB.Create(&client)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "id", Value: "1"}}

	controllers.FindClient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseClient models.Client
	json.Unmarshal(w.Body.Bytes(), &responseClient)
	assert.Equal(t, client.FirstName, responseClient.FirstName)
	assert.Equal(t, client.Email, responseClient.Email)
}

func TestUpdateClient(t *testing.T) {
	setupClientTestDB()
	client := models.Client{FirstName: "John Doe", Email: "john@example.com"}
	config.DB.Create(&client)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	updatedClient := models.Client{FirstName: "Jane Doe", Email: "jane@example.com"}
	jsonValue, _ := json.Marshal(updatedClient)
	c.Request, _ = http.NewRequest("PUT", "/clients/1", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = gin.Params{{Key: "id", Value: "1"}}

	controllers.UpdateClient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseClient models.Client
	json.Unmarshal(w.Body.Bytes(), &responseClient)
	assert.Equal(t, updatedClient.FirstName, responseClient.FirstName)
	assert.Equal(t, updatedClient.Email, responseClient.Email)
}

func TestDeleteClient(t *testing.T) {
	setupClientTestDB()
	client := models.Client{FirstName: "John Doe", Email: "john@example.com"}
	config.DB.Create(&client)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{{Key: "id", Value: "1"}}

	controllers.DeleteClient(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Client deleted")
}
