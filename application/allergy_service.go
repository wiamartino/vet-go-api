package application

import (
	"errors"
	"go-vet/domain"
	"time"
)

type AllergyService struct {
	repo domain.AllergyRepository
}

func NewAllergyService(repo domain.AllergyRepository) *AllergyService {
	return &AllergyService{repo: repo}
}

func (s *AllergyService) GetAllAllergies() ([]domain.Allergy, error) {
	return s.repo.FindAll()
}

func (s *AllergyService) GetAllergyByID(id uint) (domain.Allergy, error) {
	if id == 0 {
		return domain.Allergy{}, errors.New("invalid allergy ID")
	}
	return s.repo.FindByID(id)
}

func (s *AllergyService) GetAllergiesByPetID(petID uint) ([]domain.Allergy, error) {
	if petID == 0 {
		return nil, errors.New("invalid pet ID")
	}
	return s.repo.FindByPetID(petID)
}

func (s *AllergyService) GetActiveAllergiesByPetID(petID uint) ([]domain.Allergy, error) {
	if petID == 0 {
		return nil, errors.New("invalid pet ID")
	}
	return s.repo.FindActiveByPetID(petID)
}

func (s *AllergyService) GetAllergiesBySeverity(severity domain.AllergySeverity) ([]domain.Allergy, error) {
	return s.repo.FindBySeverity(severity)
}

func (s *AllergyService) CreateAllergy(allergy *domain.Allergy) error {
	if allergy.PetID == 0 {
		return errors.New("pet ID is required")
	}
	if allergy.Allergen == "" {
		return errors.New("allergen is required")
	}
	if allergy.AllergyType == "" {
		return errors.New("allergy type is required")
	}
	if allergy.Severity == "" {
		return errors.New("severity is required")
	}
	validTypes := map[domain.AllergyType]bool{
		domain.AllergyTypeFood:        true,
		domain.AllergyTypeMedication:  true,
		domain.AllergyTypeEnvironment: true,
		domain.AllergyTypeInsect:      true,
		domain.AllergyTypeOther:       true,
	}
	if !validTypes[allergy.AllergyType] {
		return errors.New("invalid allergy type")
	}
	validSeverities := map[domain.AllergySeverity]bool{
		domain.AllergySeverityMild:     true,
		domain.AllergySeverityModerate: true,
		domain.AllergySeveritySevere:   true,
		domain.AllergySeverityFatal:    true,
	}
	if !validSeverities[allergy.Severity] {
		return errors.New("invalid severity level")
	}
	if allergy.DiagnosedDate.IsZero() {
		allergy.DiagnosedDate = time.Now()
	}
	allergy.CreatedAt = time.Now()
	allergy.UpdatedAt = time.Now()
	return s.repo.Create(allergy)
}

func (s *AllergyService) UpdateAllergy(allergy *domain.Allergy) error {
	if allergy.AllergyID == 0 {
		return errors.New("allergy ID is required")
	}
	_, err := s.repo.FindByID(allergy.AllergyID)
	if err != nil {
		return errors.New("allergy not found")
	}
	allergy.UpdatedAt = time.Now()
	return s.repo.Update(allergy)
}

func (s *AllergyService) DeactivateAllergy(id uint) error {
	allergy, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("allergy not found")
	}
	allergy.IsActive = false
	allergy.UpdatedAt = time.Now()
	return s.repo.Update(&allergy)
}

func (s *AllergyService) ReactivateAllergy(id uint) error {
	allergy, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("allergy not found")
	}
	allergy.IsActive = true
	allergy.UpdatedAt = time.Now()
	return s.repo.Update(&allergy)
}

func (s *AllergyService) DeleteAllergy(id uint) error {
	if id == 0 {
		return errors.New("invalid allergy ID")
	}
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("allergy not found")
	}
	return s.repo.Delete(id)
}
