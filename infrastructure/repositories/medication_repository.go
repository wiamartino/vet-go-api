package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type MedicationRepository struct {
	db *database.DB
}

func NewMedicationRepository(db *database.DB) *MedicationRepository {
	return &MedicationRepository{db: db}
}

func (r *MedicationRepository) FindAll() ([]domain.Medication, error) {
	var medications []domain.Medication
	if err := r.db.Find(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

func (r *MedicationRepository) FindByID(id uint) (domain.Medication, error) {
	var medication domain.Medication
	if err := r.db.First(&medication, id).Error; err != nil {
		return domain.Medication{}, err
	}
	return medication, nil
}

func (r *MedicationRepository) Create(medication domain.Medication) error {
	return r.db.Create(&medication).Error
}

func (r *MedicationRepository) Update(medication domain.Medication) error {
	return r.db.Save(&medication).Error
}

func (r *MedicationRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Medication{}, id).Error
}
