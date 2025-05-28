package mocks

import (
	"go-vet/domain"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Create(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateLastLogin(userID uint) error {
	args := m.Called(userID)
	return args.Error(0)
}

// MockClientRepository is a mock implementation of domain.ClientRepository
type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) FindAll() ([]domain.Client, error) {
	args := m.Called()
	return args.Get(0).([]domain.Client), args.Error(1)
}

func (m *MockClientRepository) FindByID(id uint) (domain.Client, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Client), args.Error(1)
}

func (m *MockClientRepository) Create(client *domain.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) Update(client *domain.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockPetRepository is a mock implementation of domain.PetRepository
type MockPetRepository struct {
	mock.Mock
}

func (m *MockPetRepository) FindAll() ([]domain.Pet, error) {
	args := m.Called()
	return args.Get(0).([]domain.Pet), args.Error(1)
}

func (m *MockPetRepository) FindByID(id uint) (domain.Pet, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Pet), args.Error(1)
}

func (m *MockPetRepository) Create(pet *domain.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func (m *MockPetRepository) Update(pet *domain.Pet) error {
	args := m.Called(pet)
	return args.Error(0)
}

func (m *MockPetRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
