#!/bin/bash

# Race Condition and Concurrency Testing Script
# This script runs Go tests with the race detector enabled

set -e

echo "======================================"
echo "Race Condition & Concurrency Testing"
echo "======================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to run tests with race detector
run_race_tests() {
    local package=$1
    local description=$2
    
    echo -e "${YELLOW}Testing: $description${NC}"
    echo "Package: $package"
    echo ""
    
    if go test -race -v "$package" -timeout 60s 2>&1 | tee /tmp/race_test_output.log; then
        echo -e "${GREEN}✓ PASSED${NC}"
        echo ""
        return 0
    else
        echo -e "${RED}✗ FAILED${NC}"
        echo ""
        return 1
    fi
}

# Track failures
FAILED_TESTS=()

echo "1. Testing Database Concurrency"
echo "--------------------------------"
if ! run_race_tests "./tests/concurrency" "Database concurrent access and race conditions"; then
    FAILED_TESTS+=("Database Concurrency")
fi

echo ""
echo "2. Testing JWT Concurrency"
echo "--------------------------------"
if ! run_race_tests "./tests/concurrency" "JWT token generation and validation"; then
    FAILED_TESTS+=("JWT Concurrency")
fi

echo ""
echo "3. Testing Metrics Middleware Concurrency"
echo "--------------------------------"
if ! run_race_tests "./tests/concurrency" "Metrics middleware concurrent requests"; then
    FAILED_TESTS+=("Metrics Concurrency")
fi

echo ""
echo "4. Testing Service Layer Concurrency"
echo "--------------------------------"
if ! run_race_tests "./tests/concurrency" "Service validation and operations"; then
    FAILED_TESTS+=("Service Concurrency")
fi

echo ""
echo "5. Running All Application Tests with Race Detector"
echo "--------------------------------"
if ! go test -race -v ./application/... -timeout 60s; then
    FAILED_TESTS+=("Application Layer")
fi

echo ""
echo "6. Running All Repository Tests with Race Detector"
echo "--------------------------------"
if ! go test -race -v ./infrastructure/repositories/... -timeout 60s; then
    FAILED_TESTS+=("Repository Layer")
fi

echo ""
echo "7. Running Middleware Tests with Race Detector"
echo "--------------------------------"
if ! go test -race -v ./tests/middlewares/... -timeout 60s; then
    FAILED_TESTS+=("Middleware Layer")
fi

echo ""
echo "======================================"
echo "Race Detection Summary"
echo "======================================"

if [ ${#FAILED_TESTS[@]} -eq 0 ]; then
    echo -e "${GREEN}✓ All tests passed! No race conditions detected.${NC}"
    exit 0
else
    echo -e "${RED}✗ Some tests failed or race conditions detected:${NC}"
    for test in "${FAILED_TESTS[@]}"; do
        echo -e "${RED}  - $test${NC}"
    done
    exit 1
fi
