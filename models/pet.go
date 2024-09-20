package models

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
