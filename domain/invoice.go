package domain

import (
	"time"
)

type Invoice struct {
	InvoiceID     uint        `gorm:"primaryKey" json:"invoice_id"`
	Date          time.Time   `json:"date"`
	Total         float64      `json:"total" binding:"required,gt=0"`
	ClientID      uint         `json:"client_id" binding:"required"`
	AppointmentID uint         `json:"appointment_id" binding:"required"`
	Client        *Client      `json:"client,omitempty"`
	Appointment   *Appointment `json:"appointment,omitempty"`
}

type InvoiceRepository interface {
	FindAll() ([]Invoice, error)
	FindByID(id uint) (Invoice, error)
	Create(invoice *Invoice) error
	Update(invoice *Invoice) error
	Delete(id uint) error
}
