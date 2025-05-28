package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreatmentService(t *testing.T) {
	t.Run("GetAllTreatments - should return all treatments", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		expectedTreatments := []domain.Treatment{
			{TreatmentID: 1, Name: "Annual Vaccination", Description: "Complete vaccination package", Cost: 150.00},
			{TreatmentID: 2, Name: "Dental Cleaning", Description: "Professional dental cleaning", Cost: 275.00},
		}

		mockRepo.On("FindAll").Return(expectedTreatments, nil)

		// Act
		treatments, err := service.GetAllTreatments()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTreatments, treatments)
		assert.Len(t, treatments, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllTreatments - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		mockRepo.On("FindAll").Return([]domain.Treatment{}, errors.New("database connection failed"))

		// Act
		treatments, err := service.GetAllTreatments()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, treatments)
		assert.Contains(t, err.Error(), "database connection failed")
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetTreatmentByID - should return treatment when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		expectedTreatment := domain.Treatment{
			TreatmentID: 1,
			Name:        "Spay Surgery",
			Description: "Spaying surgery for female pets",
			Cost:        450.00,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedTreatment, nil)

		// Act
		treatment, err := service.GetTreatmentByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTreatment, treatment)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetTreatmentByID - should return error when treatment not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		mockRepo.On("FindByID", uint(999)).Return(domain.Treatment{}, errors.New("treatment not found"))

		// Act
		treatment, err := service.GetTreatmentByID(999)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, domain.Treatment{}, treatment)
		assert.Contains(t, err.Error(), "treatment not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateTreatment - should create treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		treatment := &domain.Treatment{
			Name:        "Microchip Implantation",
			Description: "Pet identification microchip insertion",
			Cost:        45.00,
		}

		mockRepo.On("Create", treatment).Return(nil)

		// Act
		err := service.CreateTreatment(treatment)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateTreatment - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		treatment := &domain.Treatment{
			Name:        "Invalid Treatment",
			Description: "",
			Cost:        -10.00,
		}

		mockRepo.On("Create", treatment).Return(errors.New("invalid treatment data"))

		// Act
		err := service.CreateTreatment(treatment)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid treatment data")
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateTreatment - should update treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		treatment := &domain.Treatment{
			TreatmentID: 1,
			Name:        "Annual Vaccination Updated",
			Description: "Updated vaccination package with new vaccines",
			Cost:        175.00,
		}

		mockRepo.On("Update", treatment).Return(nil)

		// Act
		err := service.UpdateTreatment(treatment)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateTreatment - should return error when treatment not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		treatment := &domain.Treatment{
			TreatmentID: 999,
			Name:        "Non-existent Treatment",
			Description: "This should fail",
			Cost:        100.00,
		}

		mockRepo.On("Update", treatment).Return(errors.New("treatment not found"))

		// Act
		err := service.UpdateTreatment(treatment)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "treatment not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteTreatment - should delete treatment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteTreatment(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteTreatment - should return error when treatment not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		mockRepo.On("Delete", uint(999)).Return(errors.New("treatment not found"))

		// Act
		err := service.DeleteTreatment(999)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "treatment not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetTreatmentByID - should handle valid treatment with zero cost", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockTreatmentRepository)
		service := application.NewTreatmentService(mockRepo)

		expectedTreatment := domain.Treatment{
			TreatmentID: 2,
			Name:        "Consultation",
			Description: "Basic consultation - no charge",
			Cost:        0.00,
		}

		mockRepo.On("FindByID", uint(2)).Return(expectedTreatment, nil)

		// Act
		treatment, err := service.GetTreatmentByID(2)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedTreatment, treatment)
		assert.Equal(t, 0.00, treatment.Cost)
		mockRepo.AssertExpectations(t)
	})
}
