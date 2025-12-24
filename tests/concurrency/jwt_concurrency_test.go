package concurrency

import (
	"go-vet/utils/jwt"
	"sync"
	"testing"
	"time"
)

// TestJWTConcurrentTokenGeneration tests concurrent token generation
func TestJWTConcurrentTokenGeneration(t *testing.T) {
	const numGoroutines = 100
	tokens := make(chan string, numGoroutines)
	errors := make(chan error, numGoroutines)
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			token, err := jwt.GenerateToken(uint(id), "test@example.com", "user")
			if err != nil {
				errors <- err
				return
			}
			tokens <- token
		}(i)
	}

	wg.Wait()
	close(tokens)
	close(errors)

	// Check for errors
	for err := range errors {
		if err != nil {
			t.Errorf("Error generating token: %v", err)
		}
	}

	// Verify we got the expected number of tokens
	tokenCount := 0
	for range tokens {
		tokenCount++
	}

	if tokenCount > 0 && tokenCount != numGoroutines {
		t.Errorf("Expected %d tokens, got %d", numGoroutines, tokenCount)
	}
}

// TestJWTConcurrentTokenValidation tests concurrent token validation
func TestJWTConcurrentTokenValidation(t *testing.T) {
	// Generate a test token first
	token, err := jwt.GenerateToken(1, "test@example.com", "user")
	if err != nil {
		t.Skip("Cannot generate token, skipping validation test")
	}

	const numGoroutines = 100
	errors := make(chan error, numGoroutines)
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			claims, err := jwt.ValidateToken(token)
			if err != nil {
				errors <- err
				return
			}
			if claims == nil {
				errors <- nil
			}
		}()
	}

	wg.Wait()
	close(errors)

	// Check for validation errors
	for err := range errors {
		if err != nil {
			t.Errorf("Error validating token: %v", err)
		}
	}
}

// TestJWTStressTest applies heavy load to JWT operations
func TestJWTStressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	const numGoroutines = 200
	const opsPerGoroutine = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				token, err := jwt.GenerateToken(uint(id), "test@example.com", "user")
				if err == nil {
					_, _ = jwt.ValidateToken(token)
				}
			}
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	totalOps := numGoroutines * opsPerGoroutine * 2
	opsPerSecond := float64(totalOps) / duration.Seconds()

	t.Logf("Completed %d JWT operations in %v (%.2f ops/sec)", totalOps, duration, opsPerSecond)
}
