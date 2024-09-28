package application

import "go-vet/domain"

type VeterinarianService struct {
	repo domain.VeterinarianRepository
}

func NewVeterinarianService(repo domain.VeterinarianRepository) *VeterinarianService {
	return &VeterinarianService{repo: repo}
}

func (s *VeterinarianService) GetAllVeterinarians() ([]domain.Veterinarian, error) {
	return s.repo.FindAll()
}

func (s *VeterinarianService) GetVeterinarianByID(id uint) (domain.Veterinarian, error) {
	return s.repo.FindByID(id)
}

func (s *VeterinarianService) CreateVeterinarian(veterinarian *domain.Veterinarian) error {
	return s.repo.Create(veterinarian)
}

func (s *VeterinarianService) UpdateVeterinarian(veterinarian *domain.Veterinarian) error {
	return s.repo.Update(veterinarian)
}

func (s *VeterinarianService) DeleteVeterinarian(id uint) error {
	return s.repo.Delete(id)
}
