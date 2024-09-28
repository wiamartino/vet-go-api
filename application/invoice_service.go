package application

import "go-vet/domain"

type InvoiceService struct {
	repo domain.InvoiceRepository
}

func NewInvoiceService(repo domain.InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) GetAllInvoices() ([]domain.Invoice, error) {
	return s.repo.FindAll()
}

func (s *InvoiceService) GetInvoiceByID(id uint) (domain.Invoice, error) {
	return s.repo.FindByID(id)
}

func (s *InvoiceService) CreateInvoice(invoice *domain.Invoice) error {
	return s.repo.Create(invoice)
}

func (s *InvoiceService) UpdateInvoice(invoice *domain.Invoice) error {
	return s.repo.Update(invoice)
}

func (s *InvoiceService) DeleteInvoice(id uint) error {
	return s.repo.Delete(id)
}
