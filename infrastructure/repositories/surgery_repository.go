package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
	"time"
)

type surgeryRepository struct {
	db *database.DB
}

func NewSurgeryRepository(db *database.DB) domain.SurgeryRepository {
	return &surgeryRepository{db: db}
}

func (r *surgeryRepository) FindAll() ([]domain.Surgery, error) {
	var surgeries []domain.Surgery
	err := r.db.Preload("Pet").Preload("Veterinarian").Find(&surgeries).Error
	return surgeries, err
}

func (r *surgeryRepository) FindByID(id uint) (domain.Surgery, error) {
	var surgery domain.Surgery
	err := r.db.Preload("Pet").Preload("Veterinarian").First(&surgery, id).Error
	return surgery, err
}

func (r *surgeryRepository) FindByPetID(petID uint) ([]domain.Surgery, error) {
	var surgeries []domain.Surgery
	err := r.db.Where("pet_id = ?", petID).Preload("Veterinarian").Order("scheduled_date DESC").Find(&surgeries).Error
	return surgeries, err
}

func (r *surgeryRepository) FindByStatus(status domain.SurgeryStatus) ([]domain.Surgery, error) {
	var surgeries []domain.Surgery
	err := r.db.Where("status = ?", status).Preload("Pet").Preload("Veterinarian").Order("scheduled_date ASC").Find(&surgeries).Error
	return surgeries, err
}

func (r *surgeryRepository) FindScheduledSurgeries(startDate, endDate time.Time) ([]domain.Surgery, error) {
	var surgeries []domain.Surgery
	err := r.db.Where("scheduled_date BETWEEN ? AND ?", startDate, endDate).Preload("Pet").Preload("Veterinarian").Order("scheduled_date ASC").Find(&surgeries).Error
	return surgeries, err
}

func (r *surgeryRepository) Create(surgery *domain.Surgery) error {
	return r.db.Create(surgery).Error
}

func (r *surgeryRepository) Update(surgery *domain.Surgery) error {
	return r.db.Save(surgery).Error
}

func (r *surgeryRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Surgery{}, id).Error
}
