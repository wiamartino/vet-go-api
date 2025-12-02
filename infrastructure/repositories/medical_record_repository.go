package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type medicalRecordRepository struct {
	db *database.DB
}

func NewMedicalRecordRepository(db *database.DB) domain.MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}

func (r *medicalRecordRepository) FindAll() ([]domain.MedicalRecord, error) {
	var records []domain.MedicalRecord
	err := r.db.Preload("Pet").Preload("Veterinarian").Preload("Appointment").Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) FindByID(id uint) (domain.MedicalRecord, error) {
	var record domain.MedicalRecord
	err := r.db.Preload("Pet").Preload("Veterinarian").Preload("Appointment").First(&record, id).Error
	return record, err
}

func (r *medicalRecordRepository) FindByPetID(petID uint) ([]domain.MedicalRecord, error) {
	var records []domain.MedicalRecord
	err := r.db.Where("pet_id = ?", petID).
		Preload("Veterinarian").
		Preload("Appointment").
		Order("visit_date DESC").
		Find(&records).Error
	return records, err
}

func (r *medicalRecordRepository) Create(record *domain.MedicalRecord) error {
	return r.db.Create(record).Error
}

func (r *medicalRecordRepository) Update(record *domain.MedicalRecord) error {
	return r.db.Save(record).Error
}

func (r *medicalRecordRepository) Delete(id uint) error {
	return r.db.Delete(&domain.MedicalRecord{}, id).Error
}
