package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-vet/config"
	"go-vet/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupVeterinarianTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.Veterinarian{})
	config.DB = db
	return db
}

func setupVeterinarianRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestUpdateVeterinarianSuccess(t *testing.T) {
	db := setupVeterinarianTestDB()
	defer db.Migrator().DropTable(&models.Veterinarian{})

	veterinarian := models.Veterinarian{FirstName: "Dr. John", Specialty: "Surgery"}
	db.Create(&veterinarian)

	router := setupVeterinarianRouter()

	updatedVeterinarian := models.Veterinarian{FirstName: "Dr. John Updated", Specialty: "General Practice"}
	jsonValue, _ := json.Marshal(updatedVeterinarian)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/veterinarians/%d", veterinarian.VeterinarianID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseVeterinarian models.Veterinarian
	json.Unmarshal(w.Body.Bytes(), &responseVeterinarian)
	assert.Equal(t, updatedVeterinarian.FirstName, responseVeterinarian.FirstName)
	assert.Equal(t, updatedVeterinarian.Specialty, responseVeterinarian.Specialty)
}

func TestUpdateVeterinarianNotFound(t *testing.T) {
	setupVeterinarianTestDB()
	defer config.DB.Migrator().DropTable(&models.Veterinarian{})

	router := setupVeterinarianRouter()

	req, _ := http.NewRequest("PUT", "/veterinarians/999", bytes.NewBuffer([]byte("{}")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Veterinarian not found")
}

func TestUpdateVeterinarianBadRequest(t *testing.T) {
	db := setupVeterinarianTestDB()
	defer db.Migrator().DropTable(&models.Veterinarian{})

	veterinarian := models.Veterinarian{FirstName: "Dr. John", Specialty: "Surgery"}
	db.Create(&veterinarian)

	router := setupVeterinarianRouter()

	req, _ := http.NewRequest("PUT", fmt.Sprintf("/veterinarians/%d", veterinarian.VeterinarianID), bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
