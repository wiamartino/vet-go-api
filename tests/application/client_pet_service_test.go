package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientService(t *testing.T) {
	t.Run("GetAllClients - should return all clients", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		expectedClients := []domain.Client{
			{ClientID: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
			{ClientID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane@example.com"},
		}

		mockRepo.On("FindAll").Return(expectedClients, nil)

		// Act
		clients, err := service.GetAllClients()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedClients, clients)
		assert.Len(t, clients, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllClients - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		mockRepo.On("FindAll").Return([]domain.Client{}, errors.New("database error"))

		// Act
		clients, err := service.GetAllClients()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, clients)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetClientByID - should return client when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		expectedClient := domain.Client{
			ClientID:  1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedClient, nil)

		// Act
		client, err := service.GetClientByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedClient, client)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateClient - should create client successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		client := &domain.Client{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@example.com",
			Phone:     "+1234567890",
		}

		mockRepo.On("Create", client).Return(nil)

		// Act
		err := service.CreateClient(client)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateClient - should update client successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		client := &domain.Client{
			ClientID:  1,
			FirstName: "John Updated",
			LastName:  "Doe",
			Email:     "john.updated@example.com",
		}

		mockRepo.On("Update", client).Return(nil)

		// Act
		err := service.UpdateClient(client)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteClient - should delete client successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockClientRepository)
		service := application.NewClientService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteClient(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestPetService(t *testing.T) {
	t.Run("GetAllPets - should return all pets", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		service := application.NewPetService(mockRepo)

		expectedPets := []domain.Pet{
			{PetID: 1, Name: "Fluffy", Species: "Cat", ClientID: 1},
			{PetID: 2, Name: "Buddy", Species: "Dog", ClientID: 2},
		}

		mockRepo.On("FindAll").Return(expectedPets, nil)

		// Act
		pets, err := service.GetAllPets()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedPets, pets)
		assert.Len(t, pets, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetPetByID - should return pet when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		service := application.NewPetService(mockRepo)

		expectedPet := domain.Pet{
			PetID:       1,
			Name:        "Fluffy",
			Species:     "Cat",
			Breed:       "Persian",
			DateOfBirth: time.Now().AddDate(-2, 0, 0),
			ClientID:    1,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedPet, nil)

		// Act
		pet, err := service.GetPetByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedPet, pet)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreatePet - should create pet successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		service := application.NewPetService(mockRepo)

		pet := &domain.Pet{
			Name:        "Buddy",
			Species:     "Dog",
			Breed:       "Golden Retriever",
			DateOfBirth: time.Now().AddDate(-1, 0, 0),
			ClientID:    1,
		}

		mockRepo.On("Create", pet).Return(nil)

		// Act
		err := service.CreatePet(pet)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdatePet - should update pet successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		service := application.NewPetService(mockRepo)

		pet := &domain.Pet{
			PetID:   1,
			Name:    "Buddy Updated",
			Species: "Dog",
			Breed:   "Golden Retriever",
		}

		mockRepo.On("Update", pet).Return(nil)

		// Act
		err := service.UpdatePet(pet)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeletePet - should delete pet successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockPetRepository)
		service := application.NewPetService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeletePet(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
