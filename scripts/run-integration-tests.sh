#!/bin/bash

# Integration Test Runner Script for Vet-Go API
# This script sets up the test environment and runs integration tests

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Print colored output
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if docker is available
check_docker() {
    if ! command -v docker &> /dev/null; then
        print_error "Docker is not installed or not in PATH"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null; then
        print_error "Docker Compose is not installed or not in PATH"
        exit 1
    fi
}

# Start test database
start_database() {
    print_info "Starting test database..."
    docker-compose -f docker-compose.test.yml up -d
    
    print_info "Waiting for database to be ready..."
    sleep 5
    
    # Check if database is ready
    for i in {1..30}; do
        if docker-compose -f docker-compose.test.yml exec -T postgres-test pg_isready -U postgres &> /dev/null; then
            print_info "Database is ready!"
            return 0
        fi
        echo -n "."
        sleep 1
    done
    
    print_error "Database failed to start"
    docker-compose -f docker-compose.test.yml logs
    exit 1
}

# Stop test database
stop_database() {
    print_info "Stopping test database..."
    docker-compose -f docker-compose.test.yml down
}

# Clean test database (remove volumes)
clean_database() {
    print_info "Cleaning test database (removing volumes)..."
    docker-compose -f docker-compose.test.yml down -v
}

# Run integration tests
run_tests() {
    local TEST_PATTERN="${1:-./tests/integration/...}"
    local VERBOSE="${2:-}"
    
    print_info "Running integration tests..."
    
    # Set test environment variables
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
    
    if [ "$VERBOSE" == "-v" ]; then
        go test "$TEST_PATTERN" -v -count=1
    else
        go test "$TEST_PATTERN" -count=1
    fi
    
    TEST_EXIT_CODE=$?
    
    if [ $TEST_EXIT_CODE -eq 0 ]; then
        print_info "All tests passed! ✓"
    else
        print_error "Some tests failed!"
    fi
    
    return $TEST_EXIT_CODE
}

# Run tests with coverage
run_tests_with_coverage() {
    print_info "Running integration tests with coverage..."
    
    # Set test environment variables
    export DB_HOST=localhost
    export DB_USER=postgres
    export DB_PASSWORD=postgres
    export DB_NAME=vet_go_test
    export DB_PORT=5433
    export DB_SSLMODE=disable
    export DB_TIMEZONE=UTC
    export JWT_SECRET_KEY=test_secret_key_for_integration_testing
    export JWT_ISSUER=vet-go-integration-test
    export JWT_TIMEOUT_HOURS=24
    
    go test ./tests/integration/... -v -coverprofile=coverage.integration.out -covermode=atomic
    
    if [ $? -eq 0 ]; then
        print_info "Coverage report generated: coverage.integration.out"
        go tool cover -func=coverage.integration.out | tail -n 1
    fi
}

# Show usage
usage() {
    echo "Usage: $0 [command] [options]"
    echo ""
    echo "Commands:"
    echo "  start           Start test database"
    echo "  stop            Stop test database"
    echo "  clean           Clean test database (remove volumes)"
    echo "  test            Run integration tests"
    echo "  test-verbose    Run integration tests with verbose output"
    echo "  coverage        Run tests with coverage report"
    echo "  all             Start database, run tests, then stop database"
    echo "  help            Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 all                    # Run complete test cycle"
    echo "  $0 test                   # Run tests (database must be running)"
    echo "  $0 test-verbose           # Run tests with verbose output"
    echo "  $0 coverage               # Run tests with coverage"
}

# Main script
main() {
    local COMMAND="${1:-help}"
    
    case "$COMMAND" in
        start)
            check_docker
            start_database
            ;;
        stop)
            check_docker
            stop_database
            ;;
        clean)
            check_docker
            clean_database
            ;;
        test)
            run_tests "./tests/integration/..."
            ;;
        test-verbose)
            run_tests "./tests/integration/..." "-v"
            ;;
        coverage)
            run_tests_with_coverage
            ;;
        all)
            check_docker
            start_database
            run_tests "./tests/integration/..." "-v"
            TEST_RESULT=$?
            stop_database
            exit $TEST_RESULT
            ;;
        help|--help|-h)
            usage
            ;;
        *)
            print_error "Unknown command: $COMMAND"
            echo ""
            usage
            exit 1
            ;;
    esac
}

# Run main function
main "$@"
