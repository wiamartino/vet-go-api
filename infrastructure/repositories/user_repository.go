package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
	"time"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Create(user domain.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return r.db.Create(&user).Error
}

func (r *UserRepository) UpdateLastLogin(userID uint) error {
	return r.db.Model(&domain.User{}).Where("id = ?", userID).
		Update("last_login", time.Now()).Error
}
