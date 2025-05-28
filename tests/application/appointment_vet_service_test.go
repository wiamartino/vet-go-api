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

func TestAppointmentService(t *testing.T) {
	t.Run("GetAllAppointments - should return all appointments", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		service := application.NewAppointmentService(mockRepo)

		expectedAppointments := []domain.Appointment{
			{
				AppointmentID:        1,
				PetID:                1,
				VeterinarianID:       1,
				ReasonForAppointment: "Annual checkup",
				Date:                 time.Now(),
			},
			{
				AppointmentID:        2,
				PetID:                2,
				VeterinarianID:       1,
				ReasonForAppointment: "Vaccination",
				Date:                 time.Now().AddDate(0, 0, 1),
			},
		}

		mockRepo.On("FindAll").Return(expectedAppointments, nil)

		// Act
		appointments, err := service.GetAllAppointments()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedAppointments, appointments)
		assert.Len(t, appointments, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAppointmentByID - should return appointment when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		service := application.NewAppointmentService(mockRepo)

		expectedAppointment := domain.Appointment{
			AppointmentID:        1,
			PetID:                1,
			VeterinarianID:       1,
			ReasonForAppointment: "Emergency visit",
			Date:                 time.Now(),
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedAppointment, nil)

		// Act
		appointment, err := service.GetAppointmentByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedAppointment, appointment)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateAppointment - should create appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		service := application.NewAppointmentService(mockRepo)

		appointment := &domain.Appointment{
			PetID:                1,
			VeterinarianID:       1,
			ReasonForAppointment: "Surgery consultation",
			Date:                 time.Now().AddDate(0, 0, 7),
		}

		mockRepo.On("Create", appointment).Return(nil)

		// Act
		err := service.CreateAppointment(appointment)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateAppointment - should update appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		service := application.NewAppointmentService(mockRepo)

		appointment := &domain.Appointment{
			AppointmentID:        1,
			PetID:                1,
			VeterinarianID:       2,
			ReasonForAppointment: "Follow-up visit",
			Date:                 time.Now().AddDate(0, 0, 14),
		}

		mockRepo.On("Update", appointment).Return(nil)

		// Act
		err := service.UpdateAppointment(appointment)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteAppointment - should delete appointment successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockAppointmentRepository)
		service := application.NewAppointmentService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteAppointment(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestVeterinarianService(t *testing.T) {
	t.Run("GetAllVeterinarians - should return all veterinarians", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		expectedVets := []domain.Veterinarian{
			{VeterinarianID: 1, FirstName: "Dr. Sarah", LastName: "Johnson", Specialty: "Small Animals"},
			{VeterinarianID: 2, FirstName: "Dr. Mike", LastName: "Wilson", Specialty: "Surgery"},
		}

		mockRepo.On("FindAll").Return(expectedVets, nil)

		// Act
		vets, err := service.GetAllVeterinarians()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedVets, vets)
		assert.Len(t, vets, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetVeterinarianByID - should return veterinarian when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		expectedVet := domain.Veterinarian{
			VeterinarianID: 1,
			FirstName:      "Dr. Sarah",
			LastName:       "Johnson",
			Specialty:      "Small Animals",
			Phone:          "+1234567890",
			Email:          "dr.johnson@vetclinic.com",
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedVet, nil)

		// Act
		vet, err := service.GetVeterinarianByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedVet, vet)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateVeterinarian - should create veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		vet := &domain.Veterinarian{
			FirstName: "Dr. Emily",
			LastName:  "Davis",
			Specialty: "Exotic Animals",
			Phone:     "+1987654321",
			Email:     "dr.davis@vetclinic.com",
		}

		mockRepo.On("Create", vet).Return(nil)

		// Act
		err := service.CreateVeterinarian(vet)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateVeterinarian - should update veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		vet := &domain.Veterinarian{
			VeterinarianID: 1,
			FirstName:      "Dr. Sarah Updated",
			LastName:       "Johnson",
			Specialty:      "Small Animals & Surgery",
		}

		mockRepo.On("Update", vet).Return(nil)

		// Act
		err := service.UpdateVeterinarian(vet)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteVeterinarian - should delete veterinarian successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteVeterinarian(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllVeterinarians - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockVeterinarianRepository)
		service := application.NewVeterinarianService(mockRepo)

		mockRepo.On("FindAll").Return([]domain.Veterinarian{}, errors.New("database connection failed"))

		// Act
		vets, err := service.GetAllVeterinarians()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, vets)
		assert.Contains(t, err.Error(), "database connection failed")
		mockRepo.AssertExpectations(t)
	})
}
