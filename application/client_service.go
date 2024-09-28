package application

import "go-vet/domain"

type ClientService struct {
	repo domain.ClientRepository
}

func NewClientService(repo domain.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetAllClients() ([]domain.Client, error) {
	return s.repo.FindAll()
}

func (s *ClientService) GetClientByID(id uint) (domain.Client, error) {
	return s.repo.FindByID(id)
}

func (s *ClientService) CreateClient(client *domain.Client) error {
	return s.repo.Create(client)
}

func (s *ClientService) UpdateClient(client *domain.Client) error {
	return s.repo.Update(client)
}

func (s *ClientService) DeleteClient(id uint) error {
	return s.repo.Delete(id)
}
