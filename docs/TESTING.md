# Test Suite Documentation

## Overview

This document describes the comprehensive test suite for the Veterinary Management System. The test suite follows Go testing best practices and includes unit tests, integration tests, and mock implementations.

## Test Structure

```
tests/
├── domain/                     # Domain model tests
│   ├── user_test.go
│   ├── client_pet_test.go
│   ├── appointment_vet_test.go
│   └── medication_treatment_invoice_test.go
├── application/                # Service layer tests
│   ├── user_service_test.go
│   ├── client_pet_service_test.go
│   └── appointment_vet_service_test.go
├── controllers/                # HTTP controller tests
│   ├── auth_controller_test.go
│   └── client_pet_controller_test.go
├── middlewares/                # Middleware tests
│   └── auth_middleware_test.go
├── integration/                # End-to-end tests
│   └── api_integration_test.go
├── mocks/                      # Mock implementations
│   ├── repositories.go
│   └── additional_repositories.go
└── testhelpers/                # Test utility functions
    └── helpers.go
```

## Test Categories

### 1. Domain Tests (`tests/domain/`)

These tests validate the domain models and their business rules:

- **User Model Tests**: Validates user creation, roles, and field validation
- **Client/Pet Model Tests**: Tests client-pet relationships and data integrity
- **Appointment/Veterinarian Tests**: Validates appointment scheduling logic
- **Medication/Treatment/Invoice Tests**: Tests pricing and billing models

**Running Domain Tests:**
```bash
go test -v ./tests/domain/...
```

### 2. Application Service Tests (`tests/application/`)

These tests validate the business logic layer using mocked repositories:

- **User Service Tests**: Registration, authentication, and user management
- **Client/Pet Service Tests**: CRUD operations with business validation
- **Appointment Service Tests**: Scheduling logic and conflict resolution
- **Veterinarian Service Tests**: Staff management and specializations

**Running Service Tests:**
```bash
go test -v ./tests/application/...
```

### 3. Controller Tests (`tests/controllers/`)

These tests validate HTTP endpoints and request/response handling:

- **Auth Controller Tests**: Registration, login, and token management
- **Client/Pet Controller Tests**: REST API endpoints for client and pet management
- **Error Handling Tests**: Validates proper error responses and status codes

**Running Controller Tests:**
```bash
go test -v ./tests/controllers/...
```

### 4. Middleware Tests (`tests/middlewares/`)

These tests validate authentication and other middleware:

- **Auth Middleware Tests**: JWT token validation and context setting
- **Error Handling**: Invalid tokens, missing headers, expired tokens

**Running Middleware Tests:**
```bash
go test -v ./tests/middlewares/...
```

### 5. Integration Tests (`tests/integration/`)

These tests validate complete workflows using the full application stack:

- **User Registration Flow**: Complete user registration and authentication
- **Client Management Flow**: End-to-end client lifecycle
- **Appointment Booking Flow**: Complete appointment scheduling workflow

**Running Integration Tests:**
```bash
go test -v ./tests/integration/...
```

## Mock Implementations

The test suite includes comprehensive mock implementations for all repository interfaces:

- `MockUserRepository`
- `MockClientRepository`
- `MockPetRepository`
- `MockAppointmentRepository`
- `MockVeterinarianRepository`
- `MockMedicationRepository`
- `MockTreatmentRepository`
- `MockInvoiceRepository`

These mocks use the `testify/mock` package and provide:
- Method call verification
- Return value configuration
- Call count assertions

## Test Helpers

The `testhelpers` package provides utility functions for creating test data:

- `CreateTestUser()`: Creates a test user with default values
- `CreateTestClient()`: Creates a test client
- `CreateTestPet()`: Creates a test pet
- `CreateTestVeterinarian()`: Creates a test veterinarian
- `CreateTestAppointment()`: Creates a test appointment
- And more...

## Running Tests

### Using the Test Runner Script

The project includes a convenient test runner script:

```bash
# Run all tests
./scripts/run-tests.sh

# Run specific test categories
./scripts/run-tests.sh unit
./scripts/run-tests.sh integration
./scripts/run-tests.sh controllers
./scripts/run-tests.sh middlewares

# Generate coverage report
./scripts/run-tests.sh coverage

# Clean test artifacts
./scripts/run-tests.sh clean
```

### Manual Test Execution

```bash
# Run all tests
go test -v ./tests/...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html

# Run only unit tests (fast)
go test -v -short ./tests/...

# Run specific package tests
go test -v ./tests/domain/
go test -v ./tests/application/
go test -v ./tests/controllers/
```

## Test Database Setup

For integration tests, you may need to set up a test database:

1. Create a separate test database
2. Set environment variables:
   ```bash
   export DB_NAME=vet_go_test
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=your_user
   export DB_PASSWORD=your_password
   ```

## Best Practices

### Test Naming
- Use descriptive test names: `TestUserService_Register_ShouldReturnErrorForInvalidEmail`
- Group related tests using subtests: `t.Run("should create user successfully", func(t *testing.T) {...})`

### Test Structure
- Follow the Arrange-Act-Assert pattern
- Use table-driven tests for multiple test cases
- Keep tests independent and isolated

### Mocking
- Mock external dependencies (databases, APIs, etc.)
- Verify mock expectations in every test
- Use dependency injection to make mocking easier

### Coverage
- Aim for high test coverage (>80%)
- Focus on critical business logic
- Don't sacrifice test quality for coverage numbers

## Continuous Integration

The test suite is designed to work with CI/CD pipelines:

```yaml
# Example GitHub Actions workflow
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.21
      - run: go mod download
      - run: ./scripts/run-tests.sh coverage
```

## Contributing

When adding new features:

1. Write tests first (TDD approach)
2. Ensure all existing tests pass
3. Add appropriate mocks for new dependencies
4. Update integration tests for new workflows
5. Maintain test documentation

## Troubleshooting

### Common Issues

1. **Import Cycle**: Avoid importing application packages in domain tests
2. **Database Connection**: Ensure test database is properly configured for integration tests
3. **Mock Expectations**: Always call `mockRepo.AssertExpectations(t)` in tests using mocks
4. **Test Isolation**: Each test should be independent and not rely on other tests

### Debug Tips

- Use `go test -v` for verbose output
- Add `t.Log()` statements for debugging
- Use `go test -run TestSpecificTest` to run single tests
- Check test coverage with `go test -cover`
