package application_test

import (
	"errors"
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInvoiceService(t *testing.T) {
	t.Run("GetAllInvoices - should return all invoices", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		expectedInvoices := []domain.Invoice{
			{
				InvoiceID:     1,
				Date:          time.Now().AddDate(0, 0, -7),
				Total:         275.50,
				ClientID:      1,
				AppointmentID: 1,
			},
			{
				InvoiceID:     2,
				Date:          time.Now().AddDate(0, 0, -3),
				Total:         150.00,
				ClientID:      2,
				AppointmentID: 2,
			},
		}

		mockRepo.On("FindAll").Return(expectedInvoices, nil)

		// Act
		invoices, err := service.GetAllInvoices()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedInvoices, invoices)
		assert.Len(t, invoices, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllInvoices - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		mockRepo.On("FindAll").Return([]domain.Invoice{}, errors.New("database connection lost"))

		// Act
		invoices, err := service.GetAllInvoices()

		// Assert
		assert.Error(t, err)
		assert.Empty(t, invoices)
		assert.Contains(t, err.Error(), "database connection lost")
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetInvoiceByID - should return invoice when found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		expectedInvoice := domain.Invoice{
			InvoiceID:     1,
			Date:          time.Now().AddDate(0, 0, -1),
			Total:         450.75,
			ClientID:      1,
			AppointmentID: 1,
		}

		mockRepo.On("FindByID", uint(1)).Return(expectedInvoice, nil)

		// Act
		invoice, err := service.GetInvoiceByID(1)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedInvoice, invoice)
		assert.Equal(t, 450.75, invoice.Total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetInvoiceByID - should return error when invoice not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		mockRepo.On("FindByID", uint(999)).Return(domain.Invoice{}, errors.New("invoice not found"))

		// Act
		invoice, err := service.GetInvoiceByID(999)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, domain.Invoice{}, invoice)
		assert.Contains(t, err.Error(), "invoice not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateInvoice - should create invoice successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		invoice := &domain.Invoice{
			Date:          time.Now(),
			Total:         325.00,
			ClientID:      3,
			AppointmentID: 3,
		}

		mockRepo.On("Create", invoice).Return(nil)

		// Act
		err := service.CreateInvoice(invoice)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateInvoice - should return error when repository fails", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		invoice := &domain.Invoice{
			Date:          time.Now(),
			Total:         -100.00, // Invalid negative total
			ClientID:      0,       // Invalid client ID
			AppointmentID: 0,       // Invalid appointment ID
		}

		mockRepo.On("Create", invoice).Return(errors.New("invalid invoice data: negative total"))

		// Act
		err := service.CreateInvoice(invoice)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid invoice data")
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateInvoice - should update invoice successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		invoice := &domain.Invoice{
			InvoiceID:     1,
			Date:          time.Now(),
			Total:         500.25,
			ClientID:      1,
			AppointmentID: 1,
		}

		mockRepo.On("Update", invoice).Return(nil)

		// Act
		err := service.UpdateInvoice(invoice)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateInvoice - should return error when invoice not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		invoice := &domain.Invoice{
			InvoiceID:     999,
			Date:          time.Now(),
			Total:         100.00,
			ClientID:      1,
			AppointmentID: 1,
		}

		mockRepo.On("Update", invoice).Return(errors.New("invoice not found"))

		// Act
		err := service.UpdateInvoice(invoice)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invoice not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteInvoice - should delete invoice successfully", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		mockRepo.On("Delete", uint(1)).Return(nil)

		// Act
		err := service.DeleteInvoice(1)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteInvoice - should return error when invoice not found", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		mockRepo.On("Delete", uint(999)).Return(errors.New("invoice not found"))

		// Act
		err := service.DeleteInvoice(999)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invoice not found")
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetInvoiceByID - should handle invoice with zero total", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		expectedInvoice := domain.Invoice{
			InvoiceID:     2,
			Date:          time.Now(),
			Total:         0.00, // Zero total for complimentary service
			ClientID:      2,
			AppointmentID: 2,
		}

		mockRepo.On("FindByID", uint(2)).Return(expectedInvoice, nil)

		// Act
		invoice, err := service.GetInvoiceByID(2)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedInvoice, invoice)
		assert.Equal(t, 0.00, invoice.Total)
		mockRepo.AssertExpectations(t)
	})

	t.Run("CreateInvoice - should handle invoice with current date", func(t *testing.T) {
		// Arrange
		mockRepo := new(mocks.MockInvoiceRepository)
		service := application.NewInvoiceService(mockRepo)

		now := time.Now()
		invoice := &domain.Invoice{
			Date:          now,
			Total:         199.99,
			ClientID:      4,
			AppointmentID: 4,
		}

		mockRepo.On("Create", invoice).Return(nil)

		// Act
		err := service.CreateInvoice(invoice)

		// Assert
		assert.NoError(t, err)
		assert.True(t, invoice.Date.Equal(now) || invoice.Date.After(now.Add(-time.Second)))
		mockRepo.AssertExpectations(t)
	})
}
