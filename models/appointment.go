package models

import (
	"time"
)

type Appointment struct {
	AppointmentID        uint         `gorm:"primaryKey" json:"appointment_id"`
	Date                 time.Time    `json:"date"`
	Time                 time.Time    `json:"time"`
	PetID                uint         `json:"pet_id"`
	VeterinarianID       uint         `json:"veterinarian_id"`
	ReasonForAppointment string       `json:"reason_for_appointment"`
	Pet                  Pet          `gorm:"foreignKey:PetID"`
	Veterinarian         Veterinarian `gorm:"foreignKey:VeterinarianID"`
}
