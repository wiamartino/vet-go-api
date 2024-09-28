package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type AppointmentRepository struct {
	db *database.DB
}

func NewAppointmentRepository(db *database.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) FindAll() ([]domain.Appointment, error) {
	var appointments []domain.Appointment
	if err := r.db.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) FindByID(id uint) (domain.Appointment, error) {
	var appointment domain.Appointment
	if err := r.db.First(&appointment, id).Error; err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *AppointmentRepository) Create(appointment domain.Appointment) error {
	return r.db.Create(&appointment).Error
}

func (r *AppointmentRepository) Update(appointment domain.Appointment) error {
	return r.db.Save(&appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Appointment{}, id).Error
}
