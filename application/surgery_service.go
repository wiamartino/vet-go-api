package application

import (
	"errors"
	"go-vet/domain"
	"time"
)

type SurgeryService struct {
	repo domain.SurgeryRepository
}

func NewSurgeryService(repo domain.SurgeryRepository) *SurgeryService {
	return &SurgeryService{repo: repo}
}

func (s *SurgeryService) GetAllSurgeries() ([]domain.Surgery, error) {
	return s.repo.FindAll()
}

func (s *SurgeryService) GetSurgeryByID(id uint) (domain.Surgery, error) {
	if id == 0 {
		return domain.Surgery{}, errors.New("invalid surgery ID")
	}
	return s.repo.FindByID(id)
}

func (s *SurgeryService) GetSurgeriesByPetID(petID uint) ([]domain.Surgery, error) {
	if petID == 0 {
		return nil, errors.New("invalid pet ID")
	}
	return s.repo.FindByPetID(petID)
}

func (s *SurgeryService) GetSurgeriesByStatus(status domain.SurgeryStatus) ([]domain.Surgery, error) {
	return s.repo.FindByStatus(status)
}

func (s *SurgeryService) GetScheduledSurgeries(startDate, endDate time.Time) ([]domain.Surgery, error) {
	if startDate.After(endDate) {
		return nil, errors.New("start date must be before end date")
	}
	return s.repo.FindScheduledSurgeries(startDate, endDate)
}

func (s *SurgeryService) CreateSurgery(surgery *domain.Surgery) error {
	if surgery.PetID == 0 {
		return errors.New("pet ID is required")
	}
	if surgery.VeterinarianID == 0 {
		return errors.New("veterinarian ID is required")
	}
	if surgery.SurgeryName == "" {
		return errors.New("surgery name is required")
	}
	if surgery.ScheduledDate.IsZero() {
		return errors.New("scheduled date is required")
	}
	if surgery.Status == "" {
		surgery.Status = domain.SurgeryStatusScheduled
	}
	if surgery.SurgeryType == "" {
		surgery.SurgeryType = domain.SurgeryTypeRoutine
	}
	surgery.CreatedAt = time.Now()
	surgery.UpdatedAt = time.Now()
	return s.repo.Create(surgery)
}

func (s *SurgeryService) UpdateSurgery(surgery *domain.Surgery) error {
	if surgery.SurgeryID == 0 {
		return errors.New("surgery ID is required")
	}
	_, err := s.repo.FindByID(surgery.SurgeryID)
	if err != nil {
		return errors.New("surgery not found")
	}
	surgery.UpdatedAt = time.Now()
	return s.repo.Update(surgery)
}

func (s *SurgeryService) StartSurgery(id uint) error {
	surgery, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("surgery not found")
	}
	if surgery.Status != domain.SurgeryStatusScheduled {
		return errors.New("surgery must be in scheduled status to start")
	}
	surgery.Status = domain.SurgeryStatusInProgress
	surgery.UpdatedAt = time.Now()
	return s.repo.Update(&surgery)
}

func (s *SurgeryService) CompleteSurgery(id uint, postOpNotes string, complications string) error {
	surgery, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("surgery not found")
	}
	now := time.Now()
	surgery.Status = domain.SurgeryStatusCompleted
	surgery.ActualDate = &now
	surgery.PostOpNotes = postOpNotes
	surgery.Complications = complications
	surgery.UpdatedAt = now
	return s.repo.Update(&surgery)
}

func (s *SurgeryService) CancelSurgery(id uint) error {
	surgery, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("surgery not found")
	}
	if surgery.Status == domain.SurgeryStatusCompleted {
		return errors.New("cannot cancel completed surgery")
	}
	surgery.Status = domain.SurgeryStatusCancelled
	surgery.UpdatedAt = time.Now()
	return s.repo.Update(&surgery)
}

func (s *SurgeryService) DeleteSurgery(id uint) error {
	if id == 0 {
		return errors.New("invalid surgery ID")
	}
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("surgery not found")
	}
	return s.repo.Delete(id)
}
