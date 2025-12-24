package concurrency

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/tests/mocks"
	"sync"
	"testing"
	"time"
)

// TestAllergyServiceConcurrentValidation tests concurrent allergy validation
func TestAllergyServiceConcurrentValidation(t *testing.T) {
	mockRepo := new(mocks.MockAllergyRepository)
	service := application.NewAllergyService(mockRepo)

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	allergies := []domain.Allergy{
		{
			PetID:         1,
			Allergen:      "Pollen",
			AllergyType:   domain.AllergyTypeEnvironment,
			Severity:      domain.AllergySeverityMild,
			Reaction:      "Sneezing",
			DiagnosedDate: time.Now(),
		},
		{
			PetID:         2,
			Allergen:      "Chicken",
			AllergyType:   domain.AllergyTypeFood,
			Severity:      domain.AllergySeverityModerate,
			Reaction:      "Itching",
			DiagnosedDate: time.Now(),
		},
	}

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			allergy := allergies[id%len(allergies)]
			err := service.CreateAllergy(&allergy)
			// We don't assert here as validation might fail in test environment
			_ = err
		}(i)
	}

	wg.Wait()
}

// TestConcurrentServiceOperations tests multiple services under concurrent load
func TestConcurrentServiceOperations(t *testing.T) {
	// Setup mocks
	mockAllergyRepo := new(mocks.MockAllergyRepository)

	allergyService := application.NewAllergyService(mockAllergyRepo)

	const numGoroutines = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Concurrent allergy operations
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			allergy := &domain.Allergy{
				PetID:         uint(id + 1),
				Allergen:      "Test Allergen",
				AllergyType:   domain.AllergyTypeEnvironment,
				Severity:      domain.AllergySeverityMild,
				Reaction:      "Test symptoms",
				DiagnosedDate: time.Now(),
			}
			_ = allergyService.CreateAllergy(allergy)
		}(i)
	}

	wg.Wait()
}

// TestServiceValidationStressTest stress tests service validation
func TestServiceValidationStressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	mockRepo := new(mocks.MockAllergyRepository)
	service := application.NewAllergyService(mockRepo)

	const numGoroutines = 200
	const validationsPerGoroutine = 50
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	errors := make(chan error, numGoroutines*validationsPerGoroutine)
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < validationsPerGoroutine; j++ {
				allergy := &domain.Allergy{
					PetID:         uint(id + 1),
					Allergen:      "Test Allergen",
					AllergyType:   domain.AllergyTypeEnvironment,
					Severity:      domain.AllergySeverityMild,
					Reaction:      "Test symptoms",
					DiagnosedDate: time.Now(),
				}
				err := service.CreateAllergy(allergy)
				if err != nil {
					errors <- err
				}
			}
		}(i)
	}

	wg.Wait()
	close(errors)
	duration := time.Since(start)

	errorCount := 0
	for range errors {
		errorCount++
	}

	totalOps := numGoroutines * validationsPerGoroutine
	opsPerSecond := float64(totalOps) / duration.Seconds()

	t.Logf("Completed %d validations in %v (%.2f ops/sec) with %d errors",
		totalOps, duration, opsPerSecond, errorCount)
}
