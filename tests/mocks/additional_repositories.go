package mocks

import (
	"go-vet/domain"

	"github.com/stretchr/testify/mock"
)

// MockAppointmentRepository is a mock implementation of domain.AppointmentRepository
type MockAppointmentRepository struct {
	mock.Mock
}

func (m *MockAppointmentRepository) FindAll() ([]domain.Appointment, error) {
	args := m.Called()
	return args.Get(0).([]domain.Appointment), args.Error(1)
}

func (m *MockAppointmentRepository) FindByID(id uint) (domain.Appointment, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Appointment), args.Error(1)
}

func (m *MockAppointmentRepository) Create(appointment *domain.Appointment) error {
	args := m.Called(appointment)
	return args.Error(0)
}

func (m *MockAppointmentRepository) Update(appointment *domain.Appointment) error {
	args := m.Called(appointment)
	return args.Error(0)
}

func (m *MockAppointmentRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockVeterinarianRepository is a mock implementation of domain.VeterinarianRepository
type MockVeterinarianRepository struct {
	mock.Mock
}

func (m *MockVeterinarianRepository) FindAll() ([]domain.Veterinarian, error) {
	args := m.Called()
	return args.Get(0).([]domain.Veterinarian), args.Error(1)
}

func (m *MockVeterinarianRepository) FindByID(id uint) (domain.Veterinarian, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Veterinarian), args.Error(1)
}

func (m *MockVeterinarianRepository) Create(veterinarian *domain.Veterinarian) error {
	args := m.Called(veterinarian)
	return args.Error(0)
}

func (m *MockVeterinarianRepository) Update(veterinarian *domain.Veterinarian) error {
	args := m.Called(veterinarian)
	return args.Error(0)
}

func (m *MockVeterinarianRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockMedicationRepository is a mock implementation of domain.MedicationRepository
type MockMedicationRepository struct {
	mock.Mock
}

func (m *MockMedicationRepository) FindAll() ([]domain.Medication, error) {
	args := m.Called()
	return args.Get(0).([]domain.Medication), args.Error(1)
}

func (m *MockMedicationRepository) FindByID(id uint) (domain.Medication, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Medication), args.Error(1)
}

func (m *MockMedicationRepository) Create(medication *domain.Medication) error {
	args := m.Called(medication)
	return args.Error(0)
}

func (m *MockMedicationRepository) Update(medication *domain.Medication) error {
	args := m.Called(medication)
	return args.Error(0)
}

func (m *MockMedicationRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockTreatmentRepository is a mock implementation of domain.TreatmentRepository
type MockTreatmentRepository struct {
	mock.Mock
}

func (m *MockTreatmentRepository) FindAll() ([]domain.Treatment, error) {
	args := m.Called()
	return args.Get(0).([]domain.Treatment), args.Error(1)
}

func (m *MockTreatmentRepository) FindByID(id uint) (domain.Treatment, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Treatment), args.Error(1)
}

func (m *MockTreatmentRepository) Create(treatment *domain.Treatment) error {
	args := m.Called(treatment)
	return args.Error(0)
}

func (m *MockTreatmentRepository) Update(treatment *domain.Treatment) error {
	args := m.Called(treatment)
	return args.Error(0)
}

func (m *MockTreatmentRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockInvoiceRepository is a mock implementation of domain.InvoiceRepository
type MockInvoiceRepository struct {
	mock.Mock
}

func (m *MockInvoiceRepository) FindAll() ([]domain.Invoice, error) {
	args := m.Called()
	return args.Get(0).([]domain.Invoice), args.Error(1)
}

func (m *MockInvoiceRepository) FindByID(id uint) (domain.Invoice, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Invoice), args.Error(1)
}

func (m *MockInvoiceRepository) Create(invoice *domain.Invoice) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *MockInvoiceRepository) Update(invoice *domain.Invoice) error {
	args := m.Called(invoice)
	return args.Error(0)
}

func (m *MockInvoiceRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
