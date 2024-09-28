package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type PetRepository struct {
	db *database.DB
}

func NewPetRepository(db *database.DB) *PetRepository {
	return &PetRepository{db: db}
}

func (r *PetRepository) FindAll() ([]domain.Pet, error) {
	var pets []domain.Pet
	if err := r.db.Preload("Client").Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func (r *PetRepository) FindByID(id uint) (domain.Pet, error) {
	var pet domain.Pet
	if err := r.db.Preload("Client").First(&pet, id).Error; err != nil {
		return domain.Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) Create(pet domain.Pet) error {
	return r.db.Create(&pet).Error
}

func (r *PetRepository) Update(pet domain.Pet) error {
	return r.db.Save(&pet).Error
}

func (r *PetRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Pet{}, id).Error
}
