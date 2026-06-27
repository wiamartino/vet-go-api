// middlewares/metrics.go
package middlewares

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var totalRequests int64
var totalErrors int64

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		reqCount := atomic.AddInt64(&totalRequests, 1)
		errCount := totalErrors
		if c.Writer.Status() >= http.StatusBadRequest {
			errCount = atomic.AddInt64(&totalErrors, 1)
		} else {
			errCount = atomic.LoadInt64(&totalErrors)
		}

		logrus.WithFields(logrus.Fields{
			"status":        c.Writer.Status(),
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"latency":       latency,
			"totalRequests": reqCount,
			"totalErrors":   errCount,
		}).Info("Request metrics")
	}
}

func GetTotalRequests() int64 {
	return atomic.LoadInt64(&totalRequests)
}

func GetTotalErrors() int64 {
	return atomic.LoadInt64(&totalErrors)
}
