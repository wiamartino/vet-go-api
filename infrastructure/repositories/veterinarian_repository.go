package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type VeterinarianRepository struct {
	db *database.DB
}

func NewVeterinarianRepository(db *database.DB) *VeterinarianRepository {
	return &VeterinarianRepository{db: db}
}

func (r *VeterinarianRepository) FindAll() ([]domain.Veterinarian, error) {
	var veterinarians []domain.Veterinarian
	if err := r.db.Find(&veterinarians).Error; err != nil {
		return nil, err
	}
	return veterinarians, nil
}

func (r *VeterinarianRepository) FindByID(id uint) (domain.Veterinarian, error) {
	var veterinarian domain.Veterinarian
	if err := r.db.First(&veterinarian, id).Error; err != nil {
		return domain.Veterinarian{}, err
	}
	return veterinarian, nil
}

func (r *VeterinarianRepository) Create(veterinarian domain.Veterinarian) error {
	return r.db.Create(&veterinarian).Error
}

func (r *VeterinarianRepository) Update(veterinarian domain.Veterinarian) error {
	return r.db.Save(&veterinarian).Error
}

func (r *VeterinarianRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Veterinarian{}, id).Error
}
