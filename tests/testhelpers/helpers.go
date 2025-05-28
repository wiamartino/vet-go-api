package testhelpers

import (
	"go-vet/domain"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupTestEnvironment sets up the test environment with required environment variables
func SetupTestEnvironment() {
	// Set required JWT environment variables for testing
	os.Setenv("JWT_SECRET_KEY", "test-secret-key-for-testing-only")
	os.Setenv("JWT_ISSUER", "vet-go-test")
	os.Setenv("JWT_TIMEOUT_HOURS", "24")
}

// SetupTestRouter creates a new Gin router for testing with proper environment setup
func SetupTestRouter() *gin.Engine {
	SetupTestEnvironment()
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// CreateTestUser creates a test user with default values
func CreateTestUser() domain.User {
	return domain.User{
		ID:        1,
		Email:     "test@example.com",
		Password:  "hashedpassword123",
		Name:      "Test User",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// CreateTestClient creates a test client with default values
func CreateTestClient() domain.Client {
	return domain.Client{
		ClientID:  1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "+1234567890",
		Address:   "123 Test Street",
	}
}

// CreateTestPet creates a test pet with default values
func CreateTestPet() domain.Pet {
	return domain.Pet{
		PetID:       1,
		Name:        "Fluffy",
		Species:     "Cat",
		Breed:       "Persian",
		DateOfBirth: time.Now().AddDate(-2, 0, 0),
		ClientID:    1,
	}
}

// CreateTestVeterinarian creates a test veterinarian with default values
func CreateTestVeterinarian() domain.Veterinarian {
	return domain.Veterinarian{
		VeterinarianID: 1,
		FirstName:      "Dr. Sarah",
		LastName:       "Johnson",
		Specialty:      "Small Animals",
		Phone:          "+1234567890",
		Email:          "dr.johnson@vetclinic.com",
	}
}

// CreateTestAppointment creates a test appointment with default values
func CreateTestAppointment() domain.Appointment {
	return domain.Appointment{
		AppointmentID:        1,
		Date:                 time.Now().AddDate(0, 0, 7),
		Time:                 time.Now().Add(2 * time.Hour),
		PetID:                1,
		VeterinarianID:       1,
		ReasonForAppointment: "Annual checkup",
	}
}

// CreateTestMedication creates a test medication with default values
func CreateTestMedication() domain.Medication {
	return domain.Medication{
		ID:          1,
		Name:        "Amoxicillin",
		Description: "Antibiotic for bacterial infections",
		Price:       25.99,
	}
}

// CreateTestTreatment creates a test treatment with default values
func CreateTestTreatment() domain.Treatment {
	return domain.Treatment{
		TreatmentID: 1,
		Name:        "Annual Vaccination",
		Description: "Complete vaccination package for pets",
		Cost:        150.00,
	}
}

// CreateTestInvoice creates a test invoice with default values
func CreateTestInvoice() domain.Invoice {
	return domain.Invoice{
		InvoiceID:     1,
		Date:          time.Now(),
		Total:         275.50,
		ClientID:      1,
		AppointmentID: 1,
	}
}

// AssertValidationError helper function for checking validation errors
func AssertValidationError(err error, field string) bool {
	if err == nil {
		return false
	}
	// You can implement specific validation error checking logic here
	// This depends on how your application handles validation errors
	return true
}
