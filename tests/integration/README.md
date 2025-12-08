# Integration Tests

This directory contains integration tests for the Vet-Go API. Integration tests validate the complete flow of the application, including database interactions, authentication, and API endpoints.

## Prerequisites

Before running integration tests, you need:

1. **PostgreSQL Database**: A running PostgreSQL instance for testing
2. **Go**: Go 1.21 or higher
3. **Environment Setup**: Proper environment variables configured

## Quick Start with Docker

The easiest way to run integration tests is using Docker Compose to set up a test database:

### 1. Start Test Database

```bash
# From the project root directory
docker-compose -f docker-compose.test.yml up -d
```

This will start a PostgreSQL container with the following configuration:
- Database: `vet_go_test`
- User: `postgres`
- Password: `postgres`
- Port: `5432`

### 2. Run Integration Tests

```bash
# Run all integration tests
go test ./tests/integration/... -v

# Run with coverage
go test ./tests/integration/... -v -cover

# Run specific test
go test ./tests/integration/... -v -run TestUserRegistrationAndAuthenticationFlow
```

### 3. Stop Test Database

```bash
docker-compose -f docker-compose.test.yml down

# To remove volumes (clean slate)
docker-compose -f docker-compose.test.yml down -v
```

## Manual Database Setup

If you prefer to set up PostgreSQL manually:

### 1. Create Test Database

```sql
CREATE DATABASE vet_go_test;
CREATE USER postgres WITH PASSWORD 'postgres';
GRANT ALL PRIVILEGES ON DATABASE vet_go_test TO postgres;
```

### 2. Set Environment Variables

The integration tests use the following environment variables:

```bash
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=vet_go_test
export DB_PORT=5432
export DB_SSLMODE=disable
export DB_TIMEZONE=UTC
export JWT_SECRET_KEY=test_secret_key_for_integration_testing
export JWT_ISSUER=vet-go-integration-test
export JWT_TIMEOUT_HOURS=24
```

### 3. Run Tests

```bash
go test ./tests/integration/... -v
```

## Test Coverage

The integration test suite covers the following workflows:

### 1. Authentication Flow
- User registration
- User login
- JWT token generation and validation

### 2. Client and Pet Management
- Create, read, update clients
- Create, read, update pets
- Client-pet relationship management

### 3. Veterinarian and Appointments
- Create veterinarians
- Book appointments
- Update appointment status
- Veterinarian-appointment relationships

### 4. Medication and Treatment
- Create medications
- Create treatments for pets
- Update medication and treatment information
- Treatment cost tracking

### 5. Invoice Management
- Create invoices for clients
- Update invoice payment status
- Invoice tracking and reporting

### 6. Error Handling
- Non-existent resource handling (404)
- Invalid data validation
- Unauthorized access (401)
- Foreign key constraint validation

## Test Structure

Each integration test follows this pattern:

1. **Setup**: Authenticate and prepare test data
2. **Action**: Execute API requests
3. **Assertion**: Verify responses and status codes
4. **Cleanup**: Database is cleaned between tests

## Continuous Integration

### Skip Integration Tests

Integration tests can be skipped in short mode:

```bash
go test ./... -short
```

### CI/CD Configuration

Example GitHub Actions workflow:

```yaml
name: Integration Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_DB: vet_go_test
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432
    
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run Integration Tests
        env:
          DB_HOST: localhost
          DB_USER: postgres
          DB_PASSWORD: postgres
          DB_NAME: vet_go_test
          DB_PORT: 5432
        run: go test ./tests/integration/... -v
```

## Troubleshooting

### Database Connection Failed

If you see "database not available" errors:

1. Verify PostgreSQL is running: `docker ps` or `pg_isready`
2. Check connection parameters match your setup
3. Ensure the test database exists: `psql -U postgres -l`
4. Check firewall/network settings

### Tests Failing

1. **Clean the database**: `docker-compose -f docker-compose.test.yml down -v`
2. **Restart containers**: `docker-compose -f docker-compose.test.yml up -d`
3. **Check logs**: `docker-compose -f docker-compose.test.yml logs`
4. **Verify Go modules**: `go mod tidy`

### Authentication Errors

- Ensure JWT_SECRET_KEY is set
- Check token generation in setupAuth helper
- Verify Authorization header format: `Bearer <token>`

## Best Practices

1. **Isolation**: Each test should be independent and not rely on other tests
2. **Cleanup**: Database is cleaned before each test
3. **Unique Data**: Use timestamps to generate unique test data
4. **Error Handling**: Always check error responses
5. **Meaningful Assertions**: Use descriptive assertion messages

## Performance

Integration tests are slower than unit tests because they:
- Use real database connections
- Perform actual HTTP requests
- Execute complete workflows

Expected runtime: 5-10 seconds for the full suite

## Contributing

When adding new integration tests:

1. Follow the existing test structure
2. Add cleanup in TearDownTest if needed
3. Use helper functions for common operations
4. Document complex test scenarios
5. Ensure tests pass in isolation and as a suite

## Resources

- [Go Testing Package](https://golang.org/pkg/testing/)
- [Testify Suite](https://github.com/stretchr/testify)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
