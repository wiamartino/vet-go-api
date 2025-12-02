package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
	"time"
)

type vaccinationRepository struct {
	db *database.DB
}

func NewVaccinationRepository(db *database.DB) domain.VaccinationRepository {
	return &vaccinationRepository{db: db}
}

func (r *vaccinationRepository) FindAll() ([]domain.Vaccination, error) {
	var vaccinations []domain.Vaccination
	err := r.db.Preload("Pet").Preload("Veterinarian").Find(&vaccinations).Error
	return vaccinations, err
}

func (r *vaccinationRepository) FindByID(id uint) (domain.Vaccination, error) {
	var vaccination domain.Vaccination
	err := r.db.Preload("Pet").Preload("Veterinarian").First(&vaccination, id).Error
	return vaccination, err
}

func (r *vaccinationRepository) FindByPetID(petID uint) ([]domain.Vaccination, error) {
	var vaccinations []domain.Vaccination
	err := r.db.Where("pet_id = ?", petID).Preload("Veterinarian").Order("date_administered DESC, date_scheduled DESC").Find(&vaccinations).Error
	return vaccinations, err
}

func (r *vaccinationRepository) FindDueVaccinations(beforeDate time.Time) ([]domain.Vaccination, error) {
	var vaccinations []domain.Vaccination
	err := r.db.Where("status = ? AND next_due_date <= ?", domain.VaccinationStatusScheduled, beforeDate).Preload("Pet").Preload("Veterinarian").Order("next_due_date ASC").Find(&vaccinations).Error
	return vaccinations, err
}

func (r *vaccinationRepository) FindOverdueVaccinations() ([]domain.Vaccination, error) {
	var vaccinations []domain.Vaccination
	now := time.Now()
	err := r.db.Where("status = ? AND next_due_date < ?", domain.VaccinationStatusScheduled, now).Preload("Pet").Preload("Veterinarian").Order("next_due_date ASC").Find(&vaccinations).Error
	return vaccinations, err
}

func (r *vaccinationRepository) Create(vaccination *domain.Vaccination) error {
	return r.db.Create(vaccination).Error
}

func (r *vaccinationRepository) Update(vaccination *domain.Vaccination) error {
	return r.db.Save(vaccination).Error
}

func (r *vaccinationRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Vaccination{}, id).Error
}
