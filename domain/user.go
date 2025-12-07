package domain

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;size:255" json:"email" binding:"required,email"`
	Password  string    `json:"password,omitempty" binding:"required,min=8"`
	Name      string    `json:"name" binding:"required,min=2,max=100"`
	Role      string    `gorm:"default:'user'" json:"role" binding:"omitempty,oneof=admin user veterinarian"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Create(user User) error
	UpdateLastLogin(userID uint) error
}
