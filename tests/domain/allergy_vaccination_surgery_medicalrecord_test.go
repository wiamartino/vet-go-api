package domain_test

import (
	"go-vet/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAllergyDomain(t *testing.T) {
	t.Run("should create allergy with valid fields", func(t *testing.T) {
		now := time.Now()
		allergy := domain.Allergy{
			AllergyID:     1,
			PetID:         1,
			Allergen:      "Peanuts",
			AllergyType:   domain.AllergyTypeFood,
			Severity:      domain.AllergySeverityMild,
			Reaction:      "Itching and hives",
			DiagnosedDate: now,
			IsActive:      true,
			CreatedAt:     now,
			UpdatedAt:     now,
		}

		assert.Equal(t, uint(1), allergy.AllergyID)
		assert.Equal(t, uint(1), allergy.PetID)
		assert.Equal(t, "Peanuts", allergy.Allergen)
		assert.Equal(t, domain.AllergyTypeFood, allergy.AllergyType)
		assert.Equal(t, domain.AllergySeverityMild, allergy.Severity)
		assert.Equal(t, "Itching and hives", allergy.Reaction)
		assert.True(t, allergy.IsActive)
	})

	t.Run("should handle all allergy types", func(t *testing.T) {
		types := []domain.AllergyType{
			domain.AllergyTypeFood,
			domain.AllergyTypeMedication,
			domain.AllergyTypeEnvironment,
			domain.AllergyTypeInsect,
			domain.AllergyTypeOther,
		}

		for _, allergyType := range types {
			allergy := domain.Allergy{
				AllergyID:   1,
				PetID:       1,
				Allergen:    "Test",
				AllergyType: allergyType,
				Severity:    domain.AllergySeverityMild,
			}
			assert.Equal(t, allergyType, allergy.AllergyType)
		}
	})

	t.Run("should handle all severity levels", func(t *testing.T) {
		severities := []domain.AllergySeverity{
			domain.AllergySeverityMild,
			domain.AllergySeverityModerate,
			domain.AllergySeveritySevere,
			domain.AllergySeverityFatal,
		}

		for _, severity := range severities {
			allergy := domain.Allergy{
				AllergyID:   1,
				PetID:       1,
				Allergen:    "Test",
				AllergyType: domain.AllergyTypeFood,
				Severity:    severity,
			}
			assert.Equal(t, severity, allergy.Severity)
		}
	})
}

func TestVaccinationDomain(t *testing.T) {
	t.Run("should create vaccination with valid fields", func(t *testing.T) {
		now := time.Now()
		vaccination := domain.Vaccination{
			VaccinationID:    1,
			PetID:            1,
			VaccineName:      "Rabies",
			Manufacturer:     "VetPharma",
			BatchNumber:      "BATCH123",
			DateAdministered: &now,
			Status:           domain.VaccinationStatusCompleted,
			Notes:            "Vaccination completed successfully",
			CreatedAt:        now,
			UpdatedAt:        now,
		}

		assert.Equal(t, uint(1), vaccination.VaccinationID)
		assert.Equal(t, uint(1), vaccination.PetID)
		assert.Equal(t, "Rabies", vaccination.VaccineName)
		assert.Equal(t, "VetPharma", vaccination.Manufacturer)
		assert.Equal(t, "BATCH123", vaccination.BatchNumber)
		assert.Equal(t, domain.VaccinationStatusCompleted, vaccination.Status)
		assert.NotNil(t, vaccination.DateAdministered)
	})

	t.Run("should handle all vaccination statuses", func(t *testing.T) {
		statuses := []domain.VaccinationStatus{
			domain.VaccinationStatusScheduled,
			domain.VaccinationStatusCompleted,
			domain.VaccinationStatusOverdue,
			domain.VaccinationStatusCancelled,
		}

		for _, status := range statuses {
			vaccination := domain.Vaccination{
				VaccinationID: 1,
				PetID:         1,
				VaccineName:   "Test Vaccine",
				Status:        status,
			}
			assert.Equal(t, status, vaccination.Status)
		}
	})
}

func TestSurgeryDomain(t *testing.T) {
	t.Run("should create surgery with valid fields", func(t *testing.T) {
		now := time.Now()
		scheduledDate := now.AddDate(0, 0, 1)
		duration := 120
		cost := 500.00

		surgery := domain.Surgery{
			SurgeryID:        1,
			PetID:            1,
			VeterinarianID:   1,
			SurgeryName:      "Spay",
			SurgeryType:      domain.SurgeryTypeRoutine,
			Status:           domain.SurgeryStatusScheduled,
			ScheduledDate:    scheduledDate,
			Duration:         &duration,
			PreOpNotes:       "Patient fasted for 12 hours",
			PostOpNotes:      "Surgery successful",
			AnesthesiaUsed:   "Isoflurane",
			FollowUpRequired: true,
			Cost:             &cost,
			CreatedAt:        now,
			UpdatedAt:        now,
		}

		assert.Equal(t, uint(1), surgery.SurgeryID)
		assert.Equal(t, uint(1), surgery.PetID)
		assert.Equal(t, uint(1), surgery.VeterinarianID)
		assert.Equal(t, "Spay", surgery.SurgeryName)
		assert.Equal(t, domain.SurgeryTypeRoutine, surgery.SurgeryType)
		assert.Equal(t, domain.SurgeryStatusScheduled, surgery.Status)
		assert.NotNil(t, surgery.Duration)
		assert.Equal(t, 120, *surgery.Duration)
		assert.NotNil(t, surgery.Cost)
		assert.Equal(t, 500.00, *surgery.Cost)
	})

	t.Run("should handle all surgery types", func(t *testing.T) {
		types := []domain.SurgeryType{
			domain.SurgeryTypeRoutine,
			domain.SurgeryTypeEmergency,
			domain.SurgeryTypeElective,
		}

		for _, surgeryType := range types {
			surgery := domain.Surgery{
				SurgeryID:      1,
				PetID:          1,
				VeterinarianID: 1,
				SurgeryName:    "Test Surgery",
				SurgeryType:    surgeryType,
				ScheduledDate:  time.Now(),
			}
			assert.Equal(t, surgeryType, surgery.SurgeryType)
		}
	})

	t.Run("should handle all surgery statuses", func(t *testing.T) {
		statuses := []domain.SurgeryStatus{
			domain.SurgeryStatusScheduled,
			domain.SurgeryStatusInProgress,
			domain.SurgeryStatusCompleted,
			domain.SurgeryStatusCancelled,
		}

		for _, status := range statuses {
			surgery := domain.Surgery{
				SurgeryID:      1,
				PetID:          1,
				VeterinarianID: 1,
				SurgeryName:    "Test Surgery",
				Status:         status,
				ScheduledDate:  time.Now(),
			}
			assert.Equal(t, status, surgery.Status)
		}
	})
}

func TestMedicalRecordDomain(t *testing.T) {
	t.Run("should create medical record with valid fields", func(t *testing.T) {
		now := time.Now()
		weight := 5.5
		temperature := 38.5
		heartRate := 120

		record := domain.MedicalRecord{
			MedicalRecordID: 1,
			PetID:           1,
			VeterinarianID:  1,
			VisitDate:       now,
			Diagnosis:       "Healthy",
			Symptoms:        "None observed",
			Notes:           "Annual checkup",
			Weight:          &weight,
			Temperature:     &temperature,
			HeartRate:       &heartRate,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		assert.Equal(t, uint(1), record.MedicalRecordID)
		assert.Equal(t, uint(1), record.PetID)
		assert.Equal(t, uint(1), record.VeterinarianID)
		assert.Equal(t, "Healthy", record.Diagnosis)
		assert.Equal(t, "None observed", record.Symptoms)
		assert.NotNil(t, record.Weight)
		assert.Equal(t, 5.5, *record.Weight)
		assert.NotNil(t, record.Temperature)
		assert.Equal(t, 38.5, *record.Temperature)
		assert.NotNil(t, record.HeartRate)
		assert.Equal(t, 120, *record.HeartRate)
	})

	t.Run("should allow optional fields to be nil", func(t *testing.T) {
		record := domain.MedicalRecord{
			MedicalRecordID: 1,
			PetID:           1,
			VeterinarianID:  1,
			VisitDate:       time.Now(),
			Diagnosis:       "Healthy",
		}

		assert.Nil(t, record.Weight)
		assert.Nil(t, record.Temperature)
		assert.Nil(t, record.HeartRate)
		assert.Nil(t, record.AppointmentID)
	})
}
