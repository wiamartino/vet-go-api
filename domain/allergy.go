package domain

import "time"

type AllergySeverity string

const (
	AllergySeverityMild     AllergySeverity = "mild"
	AllergySeverityModerate AllergySeverity = "moderate"
	AllergySeveritySevere   AllergySeverity = "severe"
	AllergySeverityFatal    AllergySeverity = "fatal"
)

type AllergyType string

const (
	AllergyTypeFood        AllergyType = "food"
	AllergyTypeMedication  AllergyType = "medication"
	AllergyTypeEnvironment AllergyType = "environment"
	AllergyTypeInsect      AllergyType = "insect"
	AllergyTypeOther       AllergyType = "other"
)

type Allergy struct {
	AllergyID     uint            `gorm:"primaryKey" json:"allergy_id"`
	PetID         uint            `json:"pet_id"`
	Allergen      string          `json:"allergen"`
	AllergyType   AllergyType     `json:"allergy_type"`
	Severity      AllergySeverity `json:"severity"`
	Reaction      string          `json:"reaction"`
	DiagnosedDate time.Time       `json:"diagnosed_date"`
	DiagnosedBy   *uint           `json:"diagnosed_by,omitempty"`
	Notes         string          `json:"notes,omitempty"`
	IsActive      bool            `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	// Relations removed to prevent FK constraint conflicts
	// Use repository methods with Preload("Pet") and Preload("Veterinarian") when needed
}

type AllergyRepository interface {
	FindAll() ([]Allergy, error)
	FindByID(id uint) (Allergy, error)
	FindByPetID(petID uint) ([]Allergy, error)
	FindActiveByPetID(petID uint) ([]Allergy, error)
	FindBySeverity(severity AllergySeverity) ([]Allergy, error)
	Create(allergy *Allergy) error
	Update(allergy *Allergy) error
	Delete(id uint) error
}
