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
)

func TestClientController(t *testing.T) {
	t.Run("FindClients - should return all clients", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		clientService := application.NewClientService(mockRepo)
		clientController := controllers.NewClientController(clientService)

		router := testhelpers.SetupTestRouter()
		router.GET("/clients", clientController.FindClients)

		expectedClients := []domain.Client{
			{ClientID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
			{ClientID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		}

		mockRepo.On("FindAll").Return(expectedClients, nil)

		// Act
		req, _ := http.NewRequest("GET", "/clients", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		clients := response["data"].([]interface{})
		assert.Len(t, clients, 2)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindClients - should return error when service fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		clientService := application.NewClientService(mockRepo)
		clientController := controllers.NewClientController(clientService)

		router := testhelpers.SetupTestRouter()
		router.GET("/clients", clientController.FindClients)

		mockRepo.On("FindAll").Return([]domain.Client{}, errors.New("database error"))

		// Act
		req, _ := http.NewRequest("GET", "/clients", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
		assert.Contains(t, response["error"], "database error")

		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateClient - should create client successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		clientService := application.NewClientService(mockRepo)
		clientController := controllers.NewClientController(clientService)

		router := testhelpers.SetupTestRouter()
		router.POST("/clients", clientController.CreateClient)

		client := domain.Client{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Phone:     "+1234567890",
			Address:   "123 Main St",
		}

		jsonData, _ := json.Marshal(client)
		mockRepo.On("Create", &client).Return(nil)

		// Act
		req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(jsonData))
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

	t.Run("CreateClient - should return error for invalid JSON", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		clientService := application.NewClientService(mockRepo)
		clientController := controllers.NewClientController(clientService)

		router := testhelpers.SetupTestRouter()
		router.POST("/clients", clientController.CreateClient)

		invalidJSON := `{"first_name": "John", "last_name": }`

		// Act
		req, _ := http.NewRequest("POST", "/clients", bytes.NewBufferString(invalidJSON))
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
}

func TestPetController(t *testing.T) {
	t.Run("FindPets - should return all pets", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		petService := application.NewPetService(mockRepo)
		petController := controllers.NewPetController(petService)

		router := testhelpers.SetupTestRouter()
		router.GET("/pets", petController.FindPets)

		expectedPets := []domain.Pet{
			{PetID: 1, Name: "Fluffy", Species: "Cat", ClientID: 1},
			{PetID: 2, Name: "Buddy", Species: "Dog", ClientID: 2},
		}

		mockRepo.On("FindAll").Return(expectedPets, nil)

		// Act
		req, _ := http.NewRequest("GET", "/pets", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])

		pets := response["data"].([]interface{})
		assert.Len(t, pets, 2)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindPet - should return pet by ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		petService := application.NewPetService(mockRepo)
		petController := controllers.NewPetController(petService)

		router := testhelpers.SetupTestRouter()
		router.GET("/pets/:id", petController.FindPet)

		expectedPet := domain.Pet{
			PetID:    1,
			Name:     "Fluffy",
			Species:  "Cat",
			Breed:    "Persian",
			ClientID: 1,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedPet, nil)

		// Act
		req, _ := http.NewRequest("GET", "/pets/1", nil)
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

	t.Run("FindPet - should return error for invalid ID", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		petService := application.NewPetService(mockRepo)
		petController := controllers.NewPetController(petService)

		router := testhelpers.SetupTestRouter()
		router.GET("/pets/:id", petController.FindPet)

		// Act
		req, _ := http.NewRequest("GET", "/pets/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Invalid pet ID", response["error"])
	})

	t.Run("FindPet - should return error when pet not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		petService := application.NewPetService(mockRepo)
		petController := controllers.NewPetController(petService)

		router := testhelpers.SetupTestRouter()
		router.GET("/pets/:id", petController.FindPet)

		mockRepo.On("FindByID", uint(999)).Return(domain.Pet{}, errors.New("pet not found"))

		// Act
		req, _ := http.NewRequest("GET", "/pets/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])
		assert.Equal(t, "Pet not found", response["error"])

		mockRepo.AssertExpectations(t)
	})

	t.Run("DeletePet - should delete pet successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		petService := application.NewPetService(mockRepo)
		petController := controllers.NewPetController(petService)

		router := testhelpers.SetupTestRouter()
		router.DELETE("/pets/:id", petController.DeletePet)

		expectedPet := domain.Pet{PetID: 1}
		mockRepo.On("FindByID", uint(1)).Return(expectedPet, nil)
		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		req, _ := http.NewRequest("DELETE", "/pets/1", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "success", response["status"])
		assert.Equal(t, "Pet deleted", response["data"])

		mockRepo.AssertExpectations(t)
	})
}
