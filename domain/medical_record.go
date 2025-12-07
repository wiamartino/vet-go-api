package domain

import (
	"time"
)

// MedicalRecord represents a pet's medical visit record
type MedicalRecord struct {
	MedicalRecordID uint      `gorm:"primaryKey" json:"medical_record_id"`
	PetID           uint      `json:"pet_id"`
	VeterinarianID  uint      `json:"veterinarian_id"`
	AppointmentID   *uint     `json:"appointment_id,omitempty"` // Optional link to appointment
	VisitDate       time.Time `json:"visit_date"`
	Diagnosis       string    `json:"diagnosis"`
	Symptoms        string    `json:"symptoms"`
	Notes           string    `json:"notes"`
	Weight          *float64  `json:"weight,omitempty"`      // in kg
	Temperature     *float64  `json:"temperature,omitempty"` // in celsius
	HeartRate       *int      `json:"heart_rate,omitempty"`  // beats per minute
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	// Relations
	Pet          Pet          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"pet,omitempty"`
	Veterinarian Veterinarian `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"veterinarian,omitempty"`
	Appointment  *Appointment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"appointment,omitempty"`
}

type MedicalRecordRepository interface {
	FindAll() ([]MedicalRecord, error)
	FindByID(id uint) (MedicalRecord, error)
	FindByPetID(petID uint) ([]MedicalRecord, error)
	Create(record *MedicalRecord) error
	Update(record *MedicalRecord) error
	Delete(id uint) error
}
