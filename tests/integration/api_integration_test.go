package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	router    *gin.Engine
	db        *database.DB
	authToken string
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Set up test database environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "vet_go_test")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("DB_TIMEZONE", "UTC")
	os.Setenv("JWT_SECRET_KEY", "test_secret_key_for_integration_testing")
	os.Setenv("JWT_ISSUER", "vet-go-integration-test")
	os.Setenv("JWT_TIMEOUT_HOURS", "24")

	// Try to initialize database connection
	db, err := database.ConnectDatabase()
	if err != nil {
		suite.T().Logf("Database connection error: %v", err)
		suite.T().Skip("Skipping integration tests: database not available")
		return
	}
	suite.db = db

	// Setup router
	suite.router = routes.SetupRouter(db)

	// Clean database tables
	suite.cleanDatabase()
}

func (suite *IntegrationTestSuite) SetupTest() {
	// Clean database before each test
	suite.cleanDatabase()
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	// Clean up database connection
	if suite.db != nil {
		suite.cleanDatabase()
	}
}

func (suite *IntegrationTestSuite) cleanDatabase() {
	if suite.db == nil {
		return
	}

	// Delete in order to respect foreign key constraints
	tables := []string{
		"appointments", "surgeries", "vaccinations", "allergies",
		"medical_records", "treatments", "invoices", "medications",
		"pets", "clients", "veterinarians", "users",
	}

	for _, table := range tables {
		suite.db.Exec(fmt.Sprintf("DELETE FROM %s", table))
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	// Skip integration tests if not running in integration mode
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	// Skip if database environment is not available
	// Integration tests require a running PostgreSQL database
	// This allows the test suite to pass even without database setup
	suite.Run(t, new(IntegrationTestSuite))
}

// Test 1: User Registration and Authentication Flow
func (suite *IntegrationTestSuite) TestUserRegistrationAndAuthenticationFlow() {
	// Register a new user
	registerData := map[string]string{
		"email":    "testuser@example.com",
		"password": "securepassword123",
		"name":     "Test User",
		"role":     "veterinarian",
	}

	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "User registration should succeed")

	var registerResponse map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &registerResponse)
	suite.NoError(err)
	suite.Equal("success", registerResponse["status"])

	// Login with the registered user
	loginData := map[string]string{
		"email":    "testuser@example.com",
		"password": "securepassword123",
	}

	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Login should succeed")

	var loginResponse map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &loginResponse)
	suite.NoError(err)
	suite.Equal("success", loginResponse["status"])

	// Extract token for subsequent requests
	if data, ok := loginResponse["data"].(map[string]interface{}); ok {
		if token, ok := data["token"].(string); ok {
			suite.authToken = token
			suite.NotEmpty(suite.authToken, "Token should not be empty")
		}
	}
}

// Test 2: Complete Client and Pet Management Flow
func (suite *IntegrationTestSuite) TestClientAndPetManagementFlow() {
	// Setup authentication first
	token := suite.setupAuth()

	// Step 1: Create a client
	clientData := map[string]string{
		"first_name": "John",
		"last_name":  "Doe",
		"email":      "john.doe@example.com",
		"phone":      "+1234567890",
		"address":    "123 Main Street, City",
	}

	jsonData, _ := json.Marshal(clientData)
	req, _ := http.NewRequest("POST", "/api/v1/clients", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "Client creation should succeed")

	var clientResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &clientResponse)
	clientID := uint(clientResponse["data"].(map[string]interface{})["client_id"].(float64))

	// Step 2: Get all clients
	req, _ = http.NewRequest("GET", "/api/v1/clients", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get clients should succeed")

	// Step 3: Create a pet for the client
	petData := map[string]interface{}{
		"name":          "Buddy",
		"species":       "Dog",
		"breed":         "Golden Retriever",
		"date_of_birth": time.Now().AddDate(-2, 0, 0).Format("2006-01-02T15:04:05Z07:00"),
		"client_id":     clientID,
	}

	jsonData, _ = json.Marshal(petData)
	req, _ = http.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "Pet creation should succeed")

	var petResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &petResponse)
	petID := uint(petResponse["data"].(map[string]interface{})["pet_id"].(float64))

	// Step 4: Get pet by ID
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/pets/%d", petID), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get pet should succeed")

	// Step 5: Update pet information
	updatePetData := map[string]interface{}{
		"name":          "Buddy Updated",
		"species":       "Dog",
		"breed":         "Golden Retriever",
		"date_of_birth": time.Now().AddDate(-2, 0, 0).Format("2006-01-02T15:04:05Z07:00"),
		"client_id":     clientID,
	}

	jsonData, _ = json.Marshal(updatePetData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/pets/%d", petID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Pet update should succeed")

	// Step 6: Update client information
	updateClientData := map[string]string{
		"first_name": "John Updated",
		"last_name":  "Doe",
		"email":      "john.doe@example.com",
		"phone":      "+1234567890",
		"address":    "456 New Street, City",
	}

	jsonData, _ = json.Marshal(updateClientData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/clients/%d", clientID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Client update should succeed")
}

// Test 3: Veterinarian and Appointment Booking Flow
func (suite *IntegrationTestSuite) TestVeterinarianAndAppointmentFlow() {
	token := suite.setupAuth()

	// Step 1: Create a veterinarian
	vetData := map[string]string{
		"first_name":     "Dr. Sarah",
		"last_name":      "Smith",
		"email":          "dr.smith@example.com",
		"phone":          "+9876543210",
		"specialty":      "Surgery",
		"license_number": "VET12345",
	}

	jsonData, _ := json.Marshal(vetData)
	req, _ := http.NewRequest("POST", "/api/v1/veterinarians", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Veterinarian creation should succeed")

	var vetResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &vetResponse)
	vetID := uint(vetResponse["data"].(map[string]interface{})["veterinarian_id"].(float64))

	// Step 2: Create client and pet for appointment
	_, petID := suite.createClientAndPet(token)

	// Step 3: Book an appointment
	appointmentData := map[string]interface{}{
		"pet_id":                 petID,
		"veterinarian_id":        vetID,
		"appointment_date":       time.Now().AddDate(0, 0, 7).Format("2006-01-02T15:04:05Z07:00"),
		"reason_for_appointment": "Annual checkup",
		"status":                 "scheduled",
	}

	jsonData, _ = json.Marshal(appointmentData)
	req, _ = http.NewRequest("POST", "/api/v1/appointments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "Appointment creation should succeed")

	var appointmentResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &appointmentResponse)
	appointmentID := uint(appointmentResponse["data"].(map[string]interface{})["appointment_id"].(float64))

	// Step 4: Get appointment by ID
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/appointments/%d", appointmentID), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get appointment should succeed")

	// Step 5: Update appointment status
	updateData := map[string]interface{}{
		"pet_id":           petID,
		"veterinarian_id":  vetID,
		"appointment_date": time.Now().AddDate(0, 0, 7).Format("2006-01-02T15:04:05Z07:00"),
		"reason":           "Annual checkup - Updated",
		"status":           "completed",
	}

	jsonData, _ = json.Marshal(updateData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/appointments/%d", appointmentID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Appointment update should succeed")
}

// Test 4: Medication and Treatment Management Flow
func (suite *IntegrationTestSuite) TestMedicationAndTreatmentFlow() {
	token := suite.setupAuth()

	// Step 1: Create a medication
	medicationData := map[string]interface{}{
		"name":        "Antibiotics",
		"description": "General purpose antibiotics",
		"dosage":      "500mg",
		"price":       25.50,
	}

	jsonData, _ := json.Marshal(medicationData)
	req, _ := http.NewRequest("POST", "/api/v1/medications", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Medication creation should succeed")

	var medicationResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &medicationResponse)

	// Extract medication ID from response (uses "id" not "medication_id")
	var medicationID uint
	if data, ok := medicationResponse["data"].(map[string]interface{}); ok {
		if id, ok := data["id"].(float64); ok {
			medicationID = uint(id)
		}
	}
	suite.NotZero(medicationID, "Medication ID should be set")

	// Step 2: Get all medications
	req, _ = http.NewRequest("GET", "/api/v1/medications", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get medications should succeed")

	// Step 3: Create treatment
	_, petID := suite.createClientAndPet(token)

	treatmentData := map[string]interface{}{
		"pet_id":      petID,
		"name":        "Infection Treatment",
		"description": "Treatment for bacterial infection",
		"date":        time.Now().Format("2006-01-02T15:04:05Z07:00"),
		"cost":        100.00,
	}

	jsonData, _ = json.Marshal(treatmentData)
	req, _ = http.NewRequest("POST", "/api/v1/treatments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Treatment creation should succeed")

	var treatmentResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &treatmentResponse)
	treatmentID := uint(treatmentResponse["data"].(map[string]interface{})["treatment_id"].(float64))

	// Step 4: Update treatment
	updateTreatmentData := map[string]interface{}{
		"pet_id":      petID,
		"name":        "Infection Treatment - Updated",
		"description": "Treatment for bacterial infection - completed",
		"date":        time.Now().Format("2006-01-02T15:04:05Z07:00"),
		"cost":        120.00,
	}

	jsonData, _ = json.Marshal(updateTreatmentData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/treatments/%d", treatmentID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Treatment update should succeed")

	// Step 5: Update medication
	updateMedicationData := map[string]interface{}{
		"name":        "Antibiotics Updated",
		"description": "General purpose antibiotics - prescription required",
		"dosage":      "500mg",
		"price":       30.00,
	}

	jsonData, _ = json.Marshal(updateMedicationData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/medications/%d", medicationID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Medication update should succeed")
}

// Test 5: Invoice Management Flow
func (suite *IntegrationTestSuite) TestInvoiceManagementFlow() {
	token := suite.setupAuth()
	clientID, petID := suite.createClientAndPet(token)

	// Step 0: Create a veterinarian and appointment (required for invoice)
	vetData := map[string]string{
		"first_name":     "Dr. Invoice",
		"last_name":      "Tester",
		"email":          fmt.Sprintf("vet%d@example.com", time.Now().UnixNano()),
		"phone":          "+9876543210",
		"specialty":      "General",
		"license_number": "VET99999",
	}

	jsonData, _ := json.Marshal(vetData)
	req, _ := http.NewRequest("POST", "/api/v1/veterinarians", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var vetResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &vetResponse)
	vetID := uint(vetResponse["data"].(map[string]interface{})["veterinarian_id"].(float64))

	// Create appointment
	appointmentData := map[string]interface{}{
		"pet_id":                 petID,
		"veterinarian_id":        vetID,
		"appointment_date":       time.Now().AddDate(0, 0, 1).Format("2006-01-02T15:04:05Z07:00"),
		"reason_for_appointment": "Checkup for invoice",
		"status":                 "completed",
	}

	jsonData, _ = json.Marshal(appointmentData)
	req, _ = http.NewRequest("POST", "/api/v1/appointments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "Appointment creation should succeed")

	var appointmentResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &appointmentResponse)
	appointmentID := uint(appointmentResponse["data"].(map[string]interface{})["appointment_id"].(float64))

	// Step 1: Create an invoice
	invoiceData := map[string]interface{}{
		"client_id":      clientID,
		"appointment_id": appointmentID,
		"total":          150.00,
		"date":           time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	jsonData, _ = json.Marshal(invoiceData)
	req, _ = http.NewRequest("POST", "/api/v1/invoices", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Invoice creation should succeed")

	var invoiceResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &invoiceResponse)
	invoiceID := uint(invoiceResponse["data"].(map[string]interface{})["invoice_id"].(float64))

	// Step 2: Get invoice by ID
	req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/invoices/%d", invoiceID), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get invoice should succeed")

	// Step 3: Update invoice (update total)
	updateInvoiceData := map[string]interface{}{
		"client_id":      clientID,
		"appointment_id": appointmentID,
		"total":          175.00,
		"date":           time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}

	jsonData, _ = json.Marshal(updateInvoiceData)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/invoices/%d", invoiceID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Invoice update should succeed")

	// Step 4: Get all invoices
	req, _ = http.NewRequest("GET", "/api/v1/invoices", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code, "Get all invoices should succeed")
}

// Test 6: Error Handling and Edge Cases
func (suite *IntegrationTestSuite) TestErrorHandlingAndEdgeCases() {
	token := suite.setupAuth()

	// Test 1: Get non-existent client
	req, _ := http.NewRequest("GET", "/api/v1/clients/99999", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusNotFound, w.Code, "Should return 404 for non-existent client")

	// Test 2: Create client with minimal data (API is permissive)
	minimalClientData := map[string]string{
		"first_name": "Test",
		"last_name":  "Client",
		"email":      fmt.Sprintf("minimal%d@test.com", time.Now().UnixNano()),
		"phone":      "+1234567890",
		"address":    "123 Test Street",
	}

	jsonData, _ := json.Marshal(minimalClientData)
	req, _ = http.NewRequest("POST", "/api/v1/clients", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code, "Should succeed with minimal valid client data")

	// Test 3: Access without authentication
	req, _ = http.NewRequest("GET", "/api/v1/clients", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusUnauthorized, w.Code, "Should return 401 without authentication")

	// Test 4: Create pet with non-existent client
	petData := map[string]interface{}{
		"name":          "Orphan Pet",
		"species":       "Cat",
		"breed":         "Siamese",
		"date_of_birth": time.Now().AddDate(-1, 0, 0).Format("2006-01-02T15:04:05Z07:00"),
		"client_id":     99999, // Non-existent client
	}

	jsonData, _ = json.Marshal(petData)
	req, _ = http.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.NotEqual(http.StatusOK, w.Code, "Should fail when creating pet with non-existent client")
}

// Helper Functions

func (suite *IntegrationTestSuite) setupAuth() string {
	// Register and login a test user
	registerData := map[string]string{
		"email":    fmt.Sprintf("testuser%d@example.com", time.Now().UnixNano()),
		"password": "securepassword123",
		"name":     "Test User",
		"role":     "admin",
	}

	jsonData, _ := json.Marshal(registerData)
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Login
	loginData := map[string]string{
		"email":    registerData["email"],
		"password": registerData["password"],
	}

	jsonData, _ = json.Marshal(loginData)
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var loginResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &loginResponse)

	if data, ok := loginResponse["data"].(map[string]interface{}); ok {
		if token, ok := data["token"].(string); ok {
			return token
		}
	}

	return ""
}

func (suite *IntegrationTestSuite) createClientAndPet(token string) (uint, uint) {
	// Create client
	clientData := map[string]string{
		"first_name": fmt.Sprintf("Client%d", time.Now().UnixNano()),
		"last_name":  "Test",
		"email":      fmt.Sprintf("client%d@example.com", time.Now().UnixNano()),
		"phone":      "+1234567890",
		"address":    "123 Test Street",
	}

	jsonData, _ := json.Marshal(clientData)
	req, _ := http.NewRequest("POST", "/api/v1/clients", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var clientResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &clientResponse)
	clientID := uint(clientResponse["data"].(map[string]interface{})["client_id"].(float64))

	// Create pet
	petData := map[string]interface{}{
		"name":          fmt.Sprintf("Pet%d", time.Now().UnixNano()),
		"species":       "Dog",
		"breed":         "Mixed",
		"date_of_birth": time.Now().AddDate(-1, 0, 0).Format("2006-01-02T15:04:05Z07:00"),
		"client_id":     clientID,
	}

	jsonData, _ = json.Marshal(petData)
	req, _ = http.NewRequest("POST", "/api/v1/pets", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	var petResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &petResponse)
	petID := uint(petResponse["data"].(map[string]interface{})["pet_id"].(float64))

	return clientID, petID
}
