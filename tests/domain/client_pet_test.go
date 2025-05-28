package domain_test

import (
	"go-vet/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientModel(t *testing.T) {
	t.Run("should create client with valid fields", func(t *testing.T) {
		client := domain.Client{
			ClientID:  1,
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 Main St",
			Phone:     "+1234567890",
			Email:     "john.doe@example.com",
		}

		assert.Equal(t, uint(1), client.ClientID)
		assert.Equal(t, "John", client.FirstName)
		assert.Equal(t, "Doe", client.LastName)
		assert.Equal(t, "123 Main St", client.Address)
		assert.Equal(t, "+1234567890", client.Phone)
		assert.Equal(t, "john.doe@example.com", client.Email)
	})

	t.Run("should handle client with pets", func(t *testing.T) {
		pets := []domain.Pet{
			{
				PetID:       1,
				Name:        "Fluffy",
				Species:     "Cat",
				Breed:       "Persian",
				DateOfBirth: time.Now().AddDate(-2, 0, 0),
				ClientID:    1,
			},
		}

		client := domain.Client{
			ClientID:  1,
			FirstName: "John",
			LastName:  "Doe",
			Pets:      pets,
		}

		assert.Equal(t, 1, len(client.Pets))
		assert.Equal(t, "Fluffy", client.Pets[0].Name)
		assert.Equal(t, uint(1), client.Pets[0].ClientID)
	})
}

func TestPetModel(t *testing.T) {
	t.Run("should create pet with valid fields", func(t *testing.T) {
		birthDate := time.Now().AddDate(-3, 0, 0)
		pet := domain.Pet{
			PetID:       1,
			Name:        "Buddy",
			Species:     "Dog",
			Breed:       "Golden Retriever",
			DateOfBirth: birthDate,
			ClientID:    1,
		}

		assert.Equal(t, uint(1), pet.PetID)
		assert.Equal(t, "Buddy", pet.Name)
		assert.Equal(t, "Dog", pet.Species)
		assert.Equal(t, "Golden Retriever", pet.Breed)
		assert.Equal(t, birthDate, pet.DateOfBirth)
		assert.Equal(t, uint(1), pet.ClientID)
	})

	t.Run("should handle different pet species", func(t *testing.T) {
		species := []string{"Dog", "Cat", "Bird", "Rabbit", "Hamster"}

		for i, speciesType := range species {
			pet := domain.Pet{
				PetID:   uint(i + 1),
				Name:    "Pet" + speciesType,
				Species: speciesType,
				Breed:   "Mixed",
			}
			assert.Equal(t, speciesType, pet.Species)
		}
	})
}
