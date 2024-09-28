package repositories

import (
	"go-vet/domain"
	"go-vet/infrastructure/database"
)

type InvoiceRepository struct {
	db *database.DB
}

func NewInvoiceRepository(db *database.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) FindAll() ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	if err := r.db.Preload("Client").Preload("Appointment").Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InvoiceRepository) FindByID(id uint) (domain.Invoice, error) {
	var invoice domain.Invoice
	if err := r.db.Preload("Client").Preload("Appointment").First(&invoice, id).Error; err != nil {
		return domain.Invoice{}, err
	}
	return invoice, nil
}

func (r *InvoiceRepository) Create(invoice domain.Invoice) error {
	return r.db.Create(&invoice).Error
}

func (r *InvoiceRepository) Update(invoice domain.Invoice) error {
	return r.db.Save(&invoice).Error
}

func (r *InvoiceRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Invoice{}, id).Error
}
