package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type TreatmentRepository struct {
	db *database.DB
}

func NewTreatmentRepository(db *database.DB) *TreatmentRepository {
	return &TreatmentRepository{db: db}
}

func (r *TreatmentRepository) FindAll() ([]domain.Treatment, error) {
	var treatments []domain.Treatment
	if err := r.db.Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *TreatmentRepository) FindByID(id uint) (domain.Treatment, error) {
	var treatment domain.Treatment
	if err := r.db.First(&treatment, id).Error; err != nil {
		return domain.Treatment{}, err
	}
	return treatment, nil
}

func (r *TreatmentRepository) Create(treatment domain.Treatment) error {
	return r.db.Create(&treatment).Error
}

func (r *TreatmentRepository) Update(treatment domain.Treatment) error {
	return r.db.Save(&treatment).Error
}

func (r *TreatmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Treatment{}, id).Error
}
