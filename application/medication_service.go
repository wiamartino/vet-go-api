package application

import "go-vet/domain"

type MedicationService struct {
	repo domain.MedicationRepository
}

func NewMedicationService(repo domain.MedicationRepository) *MedicationService {
	return &MedicationService{repo: repo}
}

func (s *MedicationService) GetAllMedications() ([]domain.Medication, error) {
	return s.repo.FindAll()
}

func (s *MedicationService) GetMedicationByID(id uint) (domain.Medication, error) {
	return s.repo.FindByID(id)
}

func (s *MedicationService) CreateMedication(medication *domain.Medication) error {
	return s.repo.Create(medication)
}

func (s *MedicationService) UpdateMedication(medication *domain.Medication) error {
	return s.repo.Update(medication)
}

func (s *MedicationService) DeleteMedication(id uint) error {
	return s.repo.Delete(id)
}
