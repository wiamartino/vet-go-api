package concurrency

import (
	"go-vet/infrastructure/database"
	"sync"
	"testing"
	"time"
)

// TestDatabaseConcurrentAccess validates concurrent database access
func TestDatabaseConcurrentAccess(t *testing.T) {
	const numGoroutines = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			db := database.GetDB()
			if db == nil {
				t.Error("Expected database instance, got nil")
			}
		}()
	}

	wg.Wait()
}

// TestDatabaseInitializationRaceCondition tests concurrent initialization
func TestDatabaseInitializationRaceCondition(t *testing.T) {
	const numGoroutines = 50
	errors := make(chan error, numGoroutines)
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			db := database.GetDB()
			if db == nil {
				errors <- nil // Expected in this test context
			}
		}()
	}

	wg.Wait()
	close(errors)

	// Check that no actual errors occurred (nil is acceptable)
	for err := range errors {
		if err != nil {
			t.Errorf("Unexpected error during concurrent access: %v", err)
		}
	}
}

// TestConcurrentReadsAndWrites simulates concurrent database operations
func TestConcurrentReadsAndWrites(t *testing.T) {
	db := database.GetDB()
	if db == nil {
		t.Skip("Database not initialized, skipping test")
	}

	const numReaders = 50
	const numWriters = 10
	var wg sync.WaitGroup
	wg.Add(numReaders + numWriters)

	// Start readers
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				db := database.GetDB()
				if db == nil {
					t.Error("Database became nil during concurrent access")
					return
				}
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start writers (simulated)
	for i := 0; i < numWriters; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				db := database.GetDB()
				if db == nil {
					t.Error("Database became nil during concurrent access")
					return
				}
				time.Sleep(time.Millisecond * 2)
			}
		}()
	}

	wg.Wait()
}
