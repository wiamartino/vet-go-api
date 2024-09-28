package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type ClientRepository struct {
	db *database.DB
}

func NewClientRepository(db *database.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) FindAll() ([]domain.Client, error) {
	var clients []domain.Client
	if err := r.db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *ClientRepository) FindByID(id uint) (domain.Client, error) {
	var client domain.Client
	if err := r.db.First(&client, id).Error; err != nil {
		return domain.Client{}, err
	}
	return client, nil
}

func (r *ClientRepository) Create(client *domain.Client) error {
	return r.db.Create(&client).Error
}

func (r *ClientRepository) Update(client *domain.Client) error {
	return r.db.Save(&client).Error
}

func (r *ClientRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Client{}, id).Error
}
