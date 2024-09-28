package domain

import (
	"time"
)

type Pet struct {
	PetID       uint      `gorm:"primaryKey" json:"pet_id"`
	Name        string    `json:"name"`
	Species     string    `json:"species"`
	Breed       string    `json:"breed"`
	DateOfBirth time.Time `json:"date_of_birth"`
	ClientID    uint      `json:"client_id"`
}

type PetRepository interface {
	FindAll() ([]Pet, error)
	FindByID(id uint) (Pet, error)
	Create(pet *Pet) error
	Update(pet *Pet) error
	Delete(id uint) error
}
