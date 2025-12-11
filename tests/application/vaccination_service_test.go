package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestVaccinationService(t *testing.T) {
	t.Run("GetAllVaccinations - should return all vaccinations successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		expectedVaccinations := []domain.Vaccination{
			{VaccinationID: 1, PetID: 1, VaccineName: "Rabies"},
			{VaccinationID: 2, PetID: 1, VaccineName: "Parvovirus"},
		}

		mockRepo.On("FindAll").Return(expectedVaccinations, nil)

		vaccinations, err := service.GetAllVaccinations()

		assert.NoError(t, err)
		assert.Equal(t, expectedVaccinations, vaccinations)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetVaccinationByID - should return vaccination when found", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		expectedVaccination := domain.Vaccination{VaccinationID: 1, PetID: 1, VaccineName: "Rabies"}
		mockRepo.On("FindByID", uint(1)).Return(expectedVaccination, nil)

		vaccination, err := service.GetVaccinationByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedVaccination, vaccination)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetVaccinationByID - should return error for invalid ID", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		_, err := service.GetVaccinationByID(0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid vaccination ID")
		mockRepo.AssertNotCalled(t, "FindByID")
	})

	t.Run("CreateVaccination - should create vaccination successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		vaccination := &domain.Vaccination{
			PetID:       1,
			VaccineName: "Rabies",
		}

		mockRepo.On("Create", mock.AnythingOfType("*domain.Vaccination")).Return(nil)

		err := service.CreateVaccination(vaccination)

		assert.NoError(t, err)
		assert.Equal(t, domain.VaccinationStatusScheduled, vaccination.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateVaccination - should return error for missing pet ID", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		vaccination := &domain.Vaccination{
			VaccineName: "Rabies",
		}

		err := service.CreateVaccination(vaccination)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "pet ID is required")
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("UpdateVaccination - should update vaccination successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		existingVaccination := domain.Vaccination{VaccinationID: 1, PetID: 1, VaccineName: "Rabies"}
		vaccination := &domain.Vaccination{
			VaccinationID: 1,
			PetID:         1,
			VaccineName:   "Updated Vaccine",
		}

		mockRepo.On("FindByID", uint(1)).Return(existingVaccination, nil)
		mockRepo.On("Update", mock.AnythingOfType("*domain.Vaccination")).Return(nil)

		err := service.UpdateVaccination(vaccination)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteVaccination - should delete vaccination successfully", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		vaccination := domain.Vaccination{VaccinationID: 1}
		mockRepo.On("FindByID", uint(1)).Return(vaccination, nil)
		mockRepo.On("Delete", uint(1)).Return(nil)

		err := service.DeleteVaccination(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteVaccination - should return error when vaccination not found", func(t *testing.T) {
		mockRepo := new(mocks.MockVaccinationRepository)
		service := application.NewVaccinationService(mockRepo)

		mockRepo.On("FindByID", uint(999)).Return(domain.Vaccination{}, errors.New("not found"))

		err := service.DeleteVaccination(999)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "vaccination not found")
		mockRepo.AssertExpectations(t)
	})
}
