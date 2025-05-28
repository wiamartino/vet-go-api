package domain_test

import (
	"go-vet/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppointmentModel(t *testing.T) {
	t.Run("should create appointment with valid fields", func(t *testing.T) {
		appointmentDate := time.Now().AddDate(0, 0, 7) // Next week
		appointmentTime := time.Now().Add(2 * time.Hour)

		appointment := domain.Appointment{
			AppointmentID:        1,
			Date:                 appointmentDate,
			Time:                 appointmentTime,
			PetID:                1,
			VeterinarianID:       1,
			ReasonForAppointment: "Annual checkup",
		}

		assert.Equal(t, uint(1), appointment.AppointmentID)
		assert.Equal(t, appointmentDate, appointment.Date)
		assert.Equal(t, appointmentTime, appointment.Time)
		assert.Equal(t, uint(1), appointment.PetID)
		assert.Equal(t, uint(1), appointment.VeterinarianID)
		assert.Equal(t, "Annual checkup", appointment.ReasonForAppointment)
	})

	t.Run("should handle different appointment reasons", func(t *testing.T) {
		reasons := []string{
			"Annual checkup",
			"Vaccination",
			"Emergency visit",
			"Surgery consultation",
			"Follow-up",
		}

		for i, reason := range reasons {
			appointment := domain.Appointment{
				AppointmentID:        uint(i + 1),
				ReasonForAppointment: reason,
				PetID:                1,
				VeterinarianID:       1,
			}
			assert.Equal(t, reason, appointment.ReasonForAppointment)
		}
	})
}

func TestVeterinarianModel(t *testing.T) {
	t.Run("should create veterinarian with valid fields", func(t *testing.T) {
		vet := domain.Veterinarian{
			VeterinarianID: 1,
			FirstName:      "Dr. Sarah",
			LastName:       "Johnson",
			Specialty:      "Small Animals",
			Phone:          "+1234567890",
			Email:          "dr.johnson@vetclinic.com",
		}

		assert.Equal(t, uint(1), vet.VeterinarianID)
		assert.Equal(t, "Dr. Sarah", vet.FirstName)
		assert.Equal(t, "Johnson", vet.LastName)
		assert.Equal(t, "Small Animals", vet.Specialty)
		assert.Equal(t, "+1234567890", vet.Phone)
		assert.Equal(t, "dr.johnson@vetclinic.com", vet.Email)
	})

	t.Run("should handle different specialties", func(t *testing.T) {
		specialties := []string{
			"Small Animals",
			"Large Animals",
			"Exotic Animals",
			"Surgery",
			"Emergency Medicine",
		}

		for i, specialty := range specialties {
			vet := domain.Veterinarian{
				VeterinarianID: uint(i + 1),
				FirstName:      "Dr. Test",
				LastName:       "Vet",
				Specialty:      specialty,
			}
			assert.Equal(t, specialty, vet.Specialty)
		}
	})
}
