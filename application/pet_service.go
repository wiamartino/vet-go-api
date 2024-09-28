package application

import "go-vet/domain"

type PetService struct {
	repo domain.PetRepository
}

func NewPetService(repo domain.PetRepository) *PetService {
	return &PetService{repo: repo}
}

func (s *PetService) GetAllPets() ([]domain.Pet, error) {
	return s.repo.FindAll()
}

func (s *PetService) GetPetByID(id uint) (domain.Pet, error) {
	return s.repo.FindByID(id)
}

func (s *PetService) CreatePet(pet *domain.Pet) error {
	return s.repo.Create(pet)
}

func (s *PetService) UpdatePet(pet *domain.Pet) error {
	return s.repo.Update(pet)
}

func (s *PetService) DeletePet(id uint) error {
	return s.repo.Delete(id)
}
