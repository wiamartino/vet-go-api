package mocks

import (
	"go-vet/domain"
	"time"

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

// MockAllergyRepository is a mock implementation of domain.AllergyRepository
type MockAllergyRepository struct {
	mock.Mock
}

func (m *MockAllergyRepository) FindAll() ([]domain.Allergy, error) {
	args := m.Called()
	return args.Get(0).([]domain.Allergy), args.Error(1)
}

func (m *MockAllergyRepository) FindByID(id uint) (domain.Allergy, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Allergy), args.Error(1)
}

func (m *MockAllergyRepository) FindByPetID(petID uint) ([]domain.Allergy, error) {
	args := m.Called(petID)
	return args.Get(0).([]domain.Allergy), args.Error(1)
}

func (m *MockAllergyRepository) FindActiveByPetID(petID uint) ([]domain.Allergy, error) {
	args := m.Called(petID)
	return args.Get(0).([]domain.Allergy), args.Error(1)
}

func (m *MockAllergyRepository) FindBySeverity(severity domain.AllergySeverity) ([]domain.Allergy, error) {
	args := m.Called(severity)
	return args.Get(0).([]domain.Allergy), args.Error(1)
}

func (m *MockAllergyRepository) Create(allergy *domain.Allergy) error {
	args := m.Called(allergy)
	return args.Error(0)
}

func (m *MockAllergyRepository) Update(allergy *domain.Allergy) error {
	args := m.Called(allergy)
	return args.Error(0)
}

func (m *MockAllergyRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockVaccinationRepository is a mock implementation of domain.VaccinationRepository
type MockVaccinationRepository struct {
	mock.Mock
}

func (m *MockVaccinationRepository) FindAll() ([]domain.Vaccination, error) {
	args := m.Called()
	return args.Get(0).([]domain.Vaccination), args.Error(1)
}

func (m *MockVaccinationRepository) FindByID(id uint) (domain.Vaccination, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Vaccination), args.Error(1)
}

func (m *MockVaccinationRepository) FindByPetID(petID uint) ([]domain.Vaccination, error) {
	args := m.Called(petID)
	return args.Get(0).([]domain.Vaccination), args.Error(1)
}

func (m *MockVaccinationRepository) FindDueVaccinations(beforeDate time.Time) ([]domain.Vaccination, error) {
	args := m.Called(beforeDate)
	return args.Get(0).([]domain.Vaccination), args.Error(1)
}

func (m *MockVaccinationRepository) FindOverdueVaccinations() ([]domain.Vaccination, error) {
	args := m.Called()
	return args.Get(0).([]domain.Vaccination), args.Error(1)
}

func (m *MockVaccinationRepository) Create(vaccination *domain.Vaccination) error {
	args := m.Called(vaccination)
	return args.Error(0)
}

func (m *MockVaccinationRepository) Update(vaccination *domain.Vaccination) error {
	args := m.Called(vaccination)
	return args.Error(0)
}

func (m *MockVaccinationRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockSurgeryRepository is a mock implementation of domain.SurgeryRepository
type MockSurgeryRepository struct {
	mock.Mock
}

func (m *MockSurgeryRepository) FindAll() ([]domain.Surgery, error) {
	args := m.Called()
	return args.Get(0).([]domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) FindByID(id uint) (domain.Surgery, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) FindByPetID(petID uint) ([]domain.Surgery, error) {
	args := m.Called(petID)
	return args.Get(0).([]domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) FindByStatus(status domain.SurgeryStatus) ([]domain.Surgery, error) {
	args := m.Called(status)
	return args.Get(0).([]domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) FindScheduledSurgeries(startDate, endDate time.Time) ([]domain.Surgery, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) FindUpcoming() ([]domain.Surgery, error) {
	args := m.Called()
	return args.Get(0).([]domain.Surgery), args.Error(1)
}

func (m *MockSurgeryRepository) Create(surgery *domain.Surgery) error {
	args := m.Called(surgery)
	return args.Error(0)
}

func (m *MockSurgeryRepository) Update(surgery *domain.Surgery) error {
	args := m.Called(surgery)
	return args.Error(0)
}

func (m *MockSurgeryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockMedicalRecordRepository is a mock implementation of domain.MedicalRecordRepository
type MockMedicalRecordRepository struct {
	mock.Mock
}

func (m *MockMedicalRecordRepository) FindAll() ([]domain.MedicalRecord, error) {
	args := m.Called()
	return args.Get(0).([]domain.MedicalRecord), args.Error(1)
}

func (m *MockMedicalRecordRepository) FindByID(id uint) (domain.MedicalRecord, error) {
	args := m.Called(id)
	return args.Get(0).(domain.MedicalRecord), args.Error(1)
}

func (m *MockMedicalRecordRepository) FindByPetID(petID uint) ([]domain.MedicalRecord, error) {
	args := m.Called(petID)
	return args.Get(0).([]domain.MedicalRecord), args.Error(1)
}

func (m *MockMedicalRecordRepository) Create(record *domain.MedicalRecord) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockMedicalRecordRepository) Update(record *domain.MedicalRecord) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockMedicalRecordRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
