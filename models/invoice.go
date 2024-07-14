package models

import (
	"time"
)

type Invoice struct {
	InvoiceID     uint        `gorm:"primaryKey" json:"invoice_id"`
	Date          time.Time   `json:"date"`
	Total         float64     `json:"total"`
	ClientID      uint        `json:"client_id"`
	AppointmentID uint        `json:"appointment_id"`
	Client        Client      `gorm:"foreignKey:ClientID"`
	Appointment   Appointment `gorm:"foreignKey:AppointmentID"`
}
