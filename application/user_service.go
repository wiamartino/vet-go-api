package application

import (
	"errors"
	"go-vet/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindByEmail(email string) (domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Register(user domain.User) error {
	// Input validation
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password are required")
	}

	// Check if password meets minimum requirements
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Hash password with appropriate cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (domain.User, error) {
	// Input validation
	if email == "" || password == "" {
		return domain.User{}, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		// Don't expose if email exists or not to prevent enumeration
		return domain.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, errors.New("invalid email or password")
	}

	// Update last login time
	if err := s.repo.UpdateLastLogin(user.ID); err != nil {
		// Log this error but don't return it to the user
		// The login was successful, this is just an auxiliary operation
	}

	return user, nil
}
