package domain

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;size:255" json:"email"`
	Password  string    `json:"password" binding:"required"` // Accept password in requests
	Name      string    `json:"name,omitempty"`
	Role      string    `gorm:"default:'user'" json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	Create(user User) error
	UpdateLastLogin(userID uint) error
}
