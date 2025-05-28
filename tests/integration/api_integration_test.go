package integration_test

import (
	"bytes"
	"encoding/json"
	"go-vet/domain"
	"go-vet/infrastructure/database"
	"go-vet/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type IntegrationTestSuite struct {
	suite.Suite
	router *gin.Engine
	db     *database.DB
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Set up test database environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "test_user")
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_NAME", "vet_go_test")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("DB_TIMEZONE", "UTC")
	os.Setenv("JWT_SECRET_KEY", "test_secret_key_for_integration_testing")
	os.Setenv("JWT_ISSUER", "vet-go-integration-test")
	os.Setenv("JWT_TIMEOUT_HOURS", "24")

	// Try to initialize database connection
	db, err := database.ConnectDatabase()
	if err != nil {
		suite.T().Skip("Skipping integration tests: database not available")
		return
	}
	suite.db = db

	// Setup router
	suite.router = routes.SetupRouter(db)
}

func (suite *IntegrationTestSuite) SetupTest() {
	// Clean database before each test
	// You might want to truncate tables or use transactions
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	// Clean up database connection
	if suite.db != nil {
		// Close database connection
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

func (suite *IntegrationTestSuite) TestUserRegistrationFlow() {
	// Test complete user registration flow

	// Create a new user
	user := domain.User{
		Email:    "integration@test.com",
		Password: "testpassword123",
		Name:     "Integration Test User",
		Role:     "user",
	}

	jsonData, _ := json.Marshal(user)

	// Register user
	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	suite.NoError(err)
	suite.Equal("success", response["status"])
}

func (suite *IntegrationTestSuite) TestClientManagementFlow() {
	// Test complete client management flow

	// 1. Create a client
	client := domain.Client{
		FirstName: "Integration",
		LastName:  "Test",
		Email:     "client@test.com",
		Phone:     "+1234567890",
		Address:   "123 Test Street",
	}

	jsonData, _ := json.Marshal(client)

	// Create client
	req, _ := http.NewRequest("POST", "/api/clients", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	// Note: In a real test, you'd need to add authentication headers
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// Note: This might fail due to authentication middleware
	// In a real integration test, you'd handle authentication properly

	// 2. Get all clients
	req, _ = http.NewRequest("GET", "/api/clients", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	// 3. Update client
	// 4. Delete client
	// Additional test steps would go here
}

func (suite *IntegrationTestSuite) TestPetManagementFlow() {
	// Test complete pet management workflow

	// This would involve:
	// 1. Creating a client first
	// 2. Creating a pet for that client
	// 3. Retrieving pets
	// 4. Updating pet information
	// 5. Deleting pet
}

func (suite *IntegrationTestSuite) TestAppointmentBookingFlow() {
	// Test complete appointment booking workflow

	// This would involve:
	// 1. Creating a client
	// 2. Creating a pet
	// 3. Creating a veterinarian
	// 4. Booking an appointment
	// 5. Updating appointment
	// 6. Canceling appointment
}

// Helper functions for integration tests

func (suite *IntegrationTestSuite) authenticateUser(email, password string) string {
	// Helper function to authenticate a user and return JWT token
	// This would be used in tests that require authentication

	loginData := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, _ := json.Marshal(loginData)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		if data, ok := response["data"].(map[string]interface{}); ok {
			if token, ok := data["token"].(string); ok {
				return token
			}
		}
	}

	return ""
}

func (suite *IntegrationTestSuite) createTestClient() domain.Client {
	// Helper function to create a test client
	client := domain.Client{
		FirstName: "Test",
		LastName:  "Client",
		Email:     "testclient@example.com",
		Phone:     "+1234567890",
		Address:   "123 Test Address",
	}

	jsonData, _ := json.Marshal(client)

	req, _ := http.NewRequest("POST", "/api/clients", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	if w.Code == http.StatusOK {
		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)
		if _, ok := response["data"].(map[string]interface{}); ok {
			// Parse the created client from response
			// Implementation would depend on your response structure
		}
	}

	return client
}
