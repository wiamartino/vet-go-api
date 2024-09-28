package application

import "go-vet/domain"

type AppointmentService struct {
	repo domain.AppointmentRepository
}

func NewAppointmentService(repo domain.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (s *AppointmentService) GetAllAppointments() ([]domain.Appointment, error) {
	return s.repo.FindAll()
}

func (s *AppointmentService) GetAppointmentByID(id uint) (domain.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *AppointmentService) CreateAppointment(appointment domain.Appointment) error {
	return s.repo.Create(appointment)
}

func (s *AppointmentService) UpdateAppointment(appointment domain.Appointment) error {
	return s.repo.Update(appointment)
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	return s.repo.Delete(id)
}
