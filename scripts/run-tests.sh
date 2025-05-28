#!/bin/bash

# Test runner script for the veterinary Go application

set -e

echo "🐾 Running Veterinary Management System Tests 🐾"
echo "================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go first."
    exit 1
fi

print_status "Go version: $(go version)"

# Navigate to project directory
cd "$(dirname "$0")/.."

# Download dependencies
print_status "Downloading dependencies..."
go mod download
go mod tidy

# Run different types of tests based on arguments
case "${1:-all}" in
    "unit")
        print_status "Running unit tests..."
        go test -v -short ./tests/domain/... ./tests/application/... ./tests/mocks/...
        ;;
    "integration")
        print_status "Running integration tests..."
        go test -v ./tests/integration/...
        ;;
    "controllers")
        print_status "Running controller tests..."
        go test -v ./tests/controllers/...
        ;;
    "middlewares")
        print_status "Running middleware tests..."
        go test -v ./tests/middlewares/...
        ;;
    "coverage")
        print_status "Running tests with coverage..."
        go test -v -coverprofile=coverage.out ./tests/...
        go tool cover -html=coverage.out -o coverage.html
        print_success "Coverage report generated: coverage.html"
        ;;
    "all")
        print_status "Running all tests..."
        
        print_status "1. Running domain tests..."
        go test -v ./tests/domain/...
        
        print_status "2. Running application service tests..."
        go test -v ./tests/application/...
        
        print_status "3. Running controller tests..."
        go test -v ./tests/controllers/...
        
        print_status "4. Running middleware tests..."
        go test -v ./tests/middlewares/...
        
        print_status "5. Running integration tests (skipped in short mode)..."
        go test -v -short ./tests/integration/...
        
        print_status "6. Generating coverage report..."
        go test -v -coverprofile=coverage.out ./tests/...
        go tool cover -html=coverage.out -o coverage.html
        
        print_success "All tests completed! Coverage report: coverage.html"
        ;;
    "clean")
        print_status "Cleaning test artifacts..."
        rm -f coverage.out coverage.html
        print_success "Test artifacts cleaned"
        ;;
    "help")
        echo "Usage: $0 [OPTION]"
        echo ""
        echo "Options:"
        echo "  unit         Run unit tests only"
        echo "  integration  Run integration tests only"
        echo "  controllers  Run controller tests only"
        echo "  middlewares  Run middleware tests only"
        echo "  coverage     Run all tests with coverage report"
        echo "  all          Run all tests (default)"
        echo "  clean        Clean test artifacts"
        echo "  help         Show this help message"
        ;;
    *)
        print_error "Unknown option: $1"
        print_warning "Use '$0 help' to see available options"
        exit 1
        ;;
esac

print_success "Test execution completed!"
