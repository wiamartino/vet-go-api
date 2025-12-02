package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type allergyRepository struct {
	db *database.DB
}

func NewAllergyRepository(db *database.DB) domain.AllergyRepository {
	return &allergyRepository{db: db}
}

func (r *allergyRepository) FindAll() ([]domain.Allergy, error) {
	var allergies []domain.Allergy
	err := r.db.Preload("Pet").Preload("Veterinarian").Find(&allergies).Error
	return allergies, err
}

func (r *allergyRepository) FindByID(id uint) (domain.Allergy, error) {
	var allergy domain.Allergy
	err := r.db.Preload("Pet").Preload("Veterinarian").First(&allergy, id).Error
	return allergy, err
}

func (r *allergyRepository) FindByPetID(petID uint) ([]domain.Allergy, error) {
	var allergies []domain.Allergy
	err := r.db.Where("pet_id = ?", petID).Preload("Veterinarian").Order("diagnosed_date DESC").Find(&allergies).Error
	return allergies, err
}

func (r *allergyRepository) FindActiveByPetID(petID uint) ([]domain.Allergy, error) {
	var allergies []domain.Allergy
	err := r.db.Where("pet_id = ? AND is_active = ?", petID, true).Preload("Veterinarian").Order("severity DESC, diagnosed_date DESC").Find(&allergies).Error
	return allergies, err
}

func (r *allergyRepository) FindBySeverity(severity domain.AllergySeverity) ([]domain.Allergy, error) {
	var allergies []domain.Allergy
	err := r.db.Where("severity = ? AND is_active = ?", severity, true).Preload("Pet").Preload("Veterinarian").Order("diagnosed_date DESC").Find(&allergies).Error
	return allergies, err
}

func (r *allergyRepository) Create(allergy *domain.Allergy) error {
	return r.db.Create(allergy).Error
}

func (r *allergyRepository) Update(allergy *domain.Allergy) error {
	return r.db.Save(allergy).Error
}

func (r *allergyRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Allergy{}, id).Error
}
