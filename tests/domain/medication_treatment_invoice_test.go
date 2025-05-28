package domain_test

import (
	"go-vet/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMedicationModel(t *testing.T) {
	t.Run("should create medication with valid fields", func(t *testing.T) {
		medication := domain.Medication{
			ID:          1,
			Name:        "Amoxicillin",
			Description: "Antibiotic for bacterial infections",
			Price:       25.99,
		}

		assert.Equal(t, uint(1), medication.ID)
		assert.Equal(t, "Amoxicillin", medication.Name)
		assert.Equal(t, "Antibiotic for bacterial infections", medication.Description)
		assert.Equal(t, 25.99, medication.Price)
	})

	t.Run("should handle different price ranges", func(t *testing.T) {
		prices := []float64{10.50, 25.99, 100.00, 250.75}

		for i, price := range prices {
			medication := domain.Medication{
				ID:    uint(i + 1),
				Name:  "TestMed",
				Price: price,
			}
			assert.Equal(t, price, medication.Price)
		}
	})
}

func TestTreatmentModel(t *testing.T) {
	t.Run("should create treatment with valid fields", func(t *testing.T) {
		treatment := domain.Treatment{
			TreatmentID: 1,
			Name:        "Annual Vaccination",
			Description: "Complete vaccination package for pets",
			Cost:        150.00,
		}

		assert.Equal(t, uint(1), treatment.TreatmentID)
		assert.Equal(t, "Annual Vaccination", treatment.Name)
		assert.Equal(t, "Complete vaccination package for pets", treatment.Description)
		assert.Equal(t, 150.00, treatment.Cost)
	})
}

func TestInvoiceModel(t *testing.T) {
	t.Run("should create invoice with valid fields", func(t *testing.T) {
		invoiceDate := time.Now()
		invoice := domain.Invoice{
			InvoiceID:     1,
			Date:          invoiceDate,
			Total:         275.50,
			ClientID:      1,
			AppointmentID: 1,
		}

		assert.Equal(t, uint(1), invoice.InvoiceID)
		assert.Equal(t, invoiceDate, invoice.Date)
		assert.Equal(t, 275.50, invoice.Total)
		assert.Equal(t, uint(1), invoice.ClientID)
		assert.Equal(t, uint(1), invoice.AppointmentID)
	})

	t.Run("should handle different invoice totals", func(t *testing.T) {
		totals := []float64{50.00, 125.75, 300.00, 1200.50}

		for i, total := range totals {
			invoice := domain.Invoice{
				InvoiceID: uint(i + 1),
				Total:     total,
			}
			assert.Equal(t, total, invoice.Total)
		}
	})
}
