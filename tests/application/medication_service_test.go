package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedicationService(t *testing.T) {
	t.Run("GetAllMedications - should return all medications", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		expectedMedications := []domain.Medication{
			{ID: 1, Name: "Amoxicillin", Description: "Antibiotic for bacterial infections", Price: 25.99},
			{ID: 2, Name: "Carprofen", Description: "Anti-inflammatory pain reliever", Price: 32.50},
		}

		mockRepo.On("FindAll").Return(expectedMedications, nil)

		// Act
		medications, err := service.GetAllMedications()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedMedications, medications)
		assert.Len(t, medications, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllMedications - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		mockRepo.On("FindAll").Return([]domain.Medication{}, errors.New("database error"))

		// Act
		medications, err := service.GetAllMedications()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, medications)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetMedicationByID - should return medication when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		expectedMedication := domain.Medication{
			ID:          1,
			Name:        "Doxycycline",
			Description: "Broad-spectrum antibiotic",
			Price:       18.75,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedMedication, nil)

		// Act
		medication, err := service.GetMedicationByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedMedication, medication)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetMedicationByID - should return error when medication not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		mockRepo.On("FindByID", uint(999)).Return(domain.Medication{}, errors.New("medication not found"))

		// Act
		medication, err := service.GetMedicationByID(999)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, domain.Medication{}, medication)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMedication - should create medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		medication := &domain.Medication{
			Name:        "Prednisone",
			Description: "Corticosteroid for inflammation",
			Price:       42.25,
		}

		mockRepo.On("Create", medication).Return(nil)

		// Act
		err := service.CreateMedication(medication)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateMedication - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		medication := &domain.Medication{
			Name:        "Test Medication",
			Description: "Test description",
			Price:       10.00,
		}

		mockRepo.On("Create", medication).Return(errors.New("database constraint violation"))

		// Act
		err := service.CreateMedication(medication)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database constraint violation")
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateMedication - should update medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		medication := &domain.Medication{
			ID:          1,
			Name:        "Amoxicillin Updated",
			Description: "Updated antibiotic description",
			Price:       28.99,
		}

		mockRepo.On("Update", medication).Return(nil)

		// Act
		err := service.UpdateMedication(medication)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateMedication - should return error when medication not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		medication := &domain.Medication{
			ID:          999,
			Name:        "Non-existent Medication",
			Description: "This should fail",
			Price:       0.00,
		}

		mockRepo.On("Update", medication).Return(errors.New("medication not found"))

		// Act
		err := service.UpdateMedication(medication)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "medication not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteMedication - should delete medication successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteMedication(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteMedication - should return error when medication not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockMedicationRepository)
		service := application.NewMedicationService(mockRepo)

		mockRepo.On("Delete", uint(999)).Return(errors.New("medication not found"))

		// Act
		err := service.DeleteMedication(999)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "medication not found")
		mockRepo.AssertExpectations(t)
	})
}
