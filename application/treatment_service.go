package application

import "go-vet/domain"

type TreatmentService struct {
	repo domain.TreatmentRepository
}

func NewTreatmentService(repo domain.TreatmentRepository) *TreatmentService {
	return &TreatmentService{repo: repo}
}

func (s *TreatmentService) GetAllTreatments() ([]domain.Treatment, error) {
	return s.repo.FindAll()
}

func (s *TreatmentService) GetTreatmentByID(id uint) (domain.Treatment, error) {
	return s.repo.FindByID(id)
}

func (s *TreatmentService) CreateTreatment(treatment domain.Treatment) error {
	return s.repo.Create(treatment)
}

func (s *TreatmentService) UpdateTreatment(treatment domain.Treatment) error {
	return s.repo.Update(treatment)
}

func (s *TreatmentService) DeleteTreatment(id uint) error {
	return s.repo.Delete(id)
}
