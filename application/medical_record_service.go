package application

import (
	"errors"
	"go-vet/domain"
	"time"
)

type MedicalRecordService struct {
	repo domain.MedicalRecordRepository
}

func NewMedicalRecordService(repo domain.MedicalRecordRepository) *MedicalRecordService {
	return &MedicalRecordService{repo: repo}
}

func (s *MedicalRecordService) GetAllRecords() ([]domain.MedicalRecord, error) {
	return s.repo.FindAll()
}

func (s *MedicalRecordService) GetRecordByID(id uint) (domain.MedicalRecord, error) {
	if id == 0 {
		return domain.MedicalRecord{}, errors.New("invalid record ID")
	}
	return s.repo.FindByID(id)
}

func (s *MedicalRecordService) GetRecordsByPetID(petID uint) ([]domain.MedicalRecord, error) {
	if petID == 0 {
		return nil, errors.New("invalid pet ID")
	}
	return s.repo.FindByPetID(petID)
}

func (s *MedicalRecordService) CreateRecord(record *domain.MedicalRecord) error {
	if record.PetID == 0 {
		return errors.New("pet ID is required")
	}
	if record.VeterinarianID == 0 {
		return errors.New("veterinarian ID is required")
	}
	if record.VisitDate.IsZero() {
		record.VisitDate = time.Now()
	}

	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()

	return s.repo.Create(record)
}

func (s *MedicalRecordService) UpdateRecord(record *domain.MedicalRecord) error {
	if record.MedicalRecordID == 0 {
		return errors.New("record ID is required")
	}

	// Check if record exists
	_, err := s.repo.FindByID(record.MedicalRecordID)
	if err != nil {
		return errors.New("record not found")
	}

	record.UpdatedAt = time.Now()
	return s.repo.Update(record)
}

func (s *MedicalRecordService) DeleteRecord(id uint) error {
	if id == 0 {
		return errors.New("invalid record ID")
	}

	// Check if record exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("record not found")
	}

	return s.repo.Delete(id)
}
