package domain

import (
	"time"
)

type VaccinationStatus string

const (
	VaccinationStatusScheduled VaccinationStatus = "scheduled"
	VaccinationStatusCompleted VaccinationStatus = "completed"
	VaccinationStatusOverdue   VaccinationStatus = "overdue"
	VaccinationStatusCancelled VaccinationStatus = "cancelled"
)

type Vaccination struct {
	VaccinationID    uint              `gorm:"primaryKey" json:"vaccination_id"`
	PetID            uint              `json:"pet_id"`
	VeterinarianID   *uint             `json:"veterinarian_id,omitempty"`
	VaccineName      string            `json:"vaccine_name"`
	Manufacturer     string            `json:"manufacturer,omitempty"`
	BatchNumber      string            `json:"batch_number,omitempty"`
	DateAdministered *time.Time        `json:"date_administered,omitempty"`
	DateScheduled    *time.Time        `json:"date_scheduled,omitempty"`
	NextDueDate      *time.Time        `json:"next_due_date,omitempty"`
	Status           VaccinationStatus `json:"status" gorm:"default:'scheduled'"`
	Notes            string            `json:"notes,omitempty"`
	SideEffects      string            `json:"side_effects,omitempty"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
	Pet              Pet               `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	Veterinarian     *Veterinarian     `gorm:"foreignKey:VeterinarianID" json:"veterinarian,omitempty"`
}

type VaccinationRepository interface {
	FindAll() ([]Vaccination, error)
	FindByID(id uint) (Vaccination, error)
	FindByPetID(petID uint) ([]Vaccination, error)
	FindDueVaccinations(beforeDate time.Time) ([]Vaccination, error)
	FindOverdueVaccinations() ([]Vaccination, error)
	Create(vaccination *Vaccination) error
	Update(vaccination *Vaccination) error
	Delete(id uint) error
}
