package concurrency

import (
	"go-vet/middlewares"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMetricsMiddlewareConcurrency tests concurrent requests through metrics middleware
func TestMetricsMiddlewareConcurrency(t *testing.T) {
	gin.SetMode(gin.TestMode)

	const numGoroutines = 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	router := gin.New()
	router.Use(middlewares.MetricsMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "/test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("Expected status 200, got %d", w.Code)
			}
		}()
	}

	wg.Wait()
}

// TestMetricsMiddlewareErrorCounting tests concurrent error counting
func TestMetricsMiddlewareErrorCounting(t *testing.T) {
	gin.SetMode(gin.TestMode)

	const numSuccess = 50
	const numErrors = 50
	var wg sync.WaitGroup
	wg.Add(numSuccess + numErrors)

	router := gin.New()
	router.Use(middlewares.MetricsMiddleware())
	router.GET("/success", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	router.GET("/error", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	})

	// Concurrent successful requests
	for i := 0; i < numSuccess; i++ {
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "/success", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}()
	}

	// Concurrent error requests
	for i := 0; i < numErrors; i++ {
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "/error", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
		}()
	}

	wg.Wait()
}

// TestMetricsMiddlewareStressTest applies heavy concurrent load
func TestMetricsMiddlewareStressTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	gin.SetMode(gin.TestMode)

	const numGoroutines = 500
	const requestsPerGoroutine = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	router := gin.New()
	router.Use(middlewares.MetricsMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerGoroutine; j++ {
				req, _ := http.NewRequest("GET", "/test", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, req)
			}
		}()
	}

	wg.Wait()
}
