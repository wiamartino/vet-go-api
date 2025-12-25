.PHONY: help build run test test-unit test-integration test-all coverage clean docker-up docker-down

# Variables
APP_NAME=vet-go
MAIN_PATH=./main.go
COVERAGE_DIR=coverage
INTEGRATION_SCRIPT=./scripts/run-integration-tests.sh

# Default target
help:
	@echo "Available targets:"
	@echo "  make build              - Build the application"
	@echo "  make run                - Run the application"
	@echo "  make run-no-auth        - Run the application with authentication disabled"
	@echo "  make test               - Run all tests (unit + integration)"
	@echo "  make test-unit          - Run unit tests only"
	@echo "  make test-integration   - Run integration tests"
	@echo "  make test-coverage      - Run tests with coverage report"
	@echo "  make docker-up          - Start test database"
	@echo "  make docker-down        - Stop test database"
	@echo "  make docker-clean       - Clean test database (remove volumes)"
	@echo "  make clean              - Clean build artifacts and coverage"
	@echo "  make lint               - Run golangci-lint"
	@echo "  make fmt                - Format code"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

# Run the application
run:
	@echo "Running $(APP_NAME)..."
	go run $(MAIN_PATH)

# Run the application with authentication disabled (DEVELOPMENT ONLY)
run-no-auth:
	@echo "⚠️  Running $(APP_NAME) with AUTHENTICATION DISABLED ⚠️"
	@echo "⚠️  This should ONLY be used in development!"
	DISABLE_AUTH=true go run $(MAIN_PATH)

# Run all tests
test: test-unit test-integration

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	go test ./tests/application/... ./tests/controllers/... ./tests/domain/... ./tests/middlewares/... ./tests/utils/... -v -cover

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	@if [ -f $(INTEGRATION_SCRIPT) ]; then \
		$(INTEGRATION_SCRIPT) all; \
	else \
		echo "Integration test script not found"; \
		exit 1; \
	fi

# Run integration tests (database must be running)
test-integration-only:
	@echo "Running integration tests (assuming database is running)..."
	@export DB_HOST=localhost && \
	export DB_USER=postgres && \
	export DB_PASSWORD=postgres && \
	export DB_NAME=vet_go_test && \
	export DB_PORT=5433 && \
	export DB_SSLMODE=disable && \
	export DB_TIMEZONE=UTC && \
	export JWT_SECRET_KEY=test_secret_key_for_integration_testing && \
	export JWT_ISSUER=vet-go-integration-test && \
	export JWT_TIMEOUT_HOURS=24 && \
	go test ./tests/integration/... -v -count=1

# Run all tests
test-all:
	@echo "Running all tests..."
	go test ./... -v

# Generate coverage report
test-coverage:
	@echo "Generating coverage report..."
	@mkdir -p $(COVERAGE_DIR)
	go test ./... -coverprofile=$(COVERAGE_DIR)/coverage.out -covermode=atomic
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report generated: $(COVERAGE_DIR)/coverage.html"

# Start test database
docker-up:
	@echo "Starting test database..."
	docker-compose -f docker-compose.test.yml up -d
	@echo "Waiting for database to be ready..."
	@sleep 5

# Stop test database
docker-down:
	@echo "Stopping test database..."
	docker-compose -f docker-compose.test.yml down

# Clean test database (remove volumes)
docker-clean:
	@echo "Cleaning test database..."
	docker-compose -f docker-compose.test.yml down -v

# Clean build artifacts and coverage
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf $(COVERAGE_DIR)/
	rm -f coverage.*.out
	go clean

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run linter
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: brew install golangci-lint"; \
	fi

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Run the application in development mode with hot reload
dev:
	@echo "Running in development mode..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "air not installed. Install with: go install github.com/cosmtrek/air@latest"; \
		echo "Running without hot reload..."; \
		go run $(MAIN_PATH); \
	fi
