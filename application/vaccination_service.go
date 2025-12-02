package application

import (
	"errors"
	"go-vet/domain"
	"time"
)

type VaccinationService struct {
	repo domain.VaccinationRepository
}

func NewVaccinationService(repo domain.VaccinationRepository) *VaccinationService {
	return &VaccinationService{repo: repo}
}

func (s *VaccinationService) GetAllVaccinations() ([]domain.Vaccination, error) {
	return s.repo.FindAll()
}

func (s *VaccinationService) GetVaccinationByID(id uint) (domain.Vaccination, error) {
	if id == 0 {
		return domain.Vaccination{}, errors.New("invalid vaccination ID")
	}
	return s.repo.FindByID(id)
}

func (s *VaccinationService) GetVaccinationsByPetID(petID uint) ([]domain.Vaccination, error) {
	if petID == 0 {
		return nil, errors.New("invalid pet ID")
	}
	return s.repo.FindByPetID(petID)
}

func (s *VaccinationService) GetDueVaccinations(daysAhead int) ([]domain.Vaccination, error) {
	beforeDate := time.Now().AddDate(0, 0, daysAhead)
	return s.repo.FindDueVaccinations(beforeDate)
}

func (s *VaccinationService) GetOverdueVaccinations() ([]domain.Vaccination, error) {
	return s.repo.FindOverdueVaccinations()
}

func (s *VaccinationService) CreateVaccination(vaccination *domain.Vaccination) error {
	if vaccination.PetID == 0 {
		return errors.New("pet ID is required")
	}
	if vaccination.VaccineName == "" {
		return errors.New("vaccine name is required")
	}
	if vaccination.Status == "" {
		vaccination.Status = domain.VaccinationStatusScheduled
	}
	vaccination.CreatedAt = time.Now()
	vaccination.UpdatedAt = time.Now()
	return s.repo.Create(vaccination)
}

func (s *VaccinationService) UpdateVaccination(vaccination *domain.Vaccination) error {
	if vaccination.VaccinationID == 0 {
		return errors.New("vaccination ID is required")
	}
	_, err := s.repo.FindByID(vaccination.VaccinationID)
	if err != nil {
		return errors.New("vaccination not found")
	}
	vaccination.UpdatedAt = time.Now()
	return s.repo.Update(vaccination)
}

func (s *VaccinationService) CompleteVaccination(id uint, veterinarianID uint, sideEffects string) error {
	vaccination, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("vaccination not found")
	}
	now := time.Now()
	vaccination.DateAdministered = &now
	vaccination.VeterinarianID = &veterinarianID
	vaccination.Status = domain.VaccinationStatusCompleted
	vaccination.SideEffects = sideEffects
	vaccination.UpdatedAt = now
	return s.repo.Update(&vaccination)
}

func (s *VaccinationService) DeleteVaccination(id uint) error {
	if id == 0 {
		return errors.New("invalid vaccination ID")
	}
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("vaccination not found")
	}
	return s.repo.Delete(id)
}
