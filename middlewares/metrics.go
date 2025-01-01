// middlewares/metrics.go
package middlewares

import (
	"net/http"
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
		totalRequests++
		if c.Writer.Status() >= http.StatusBadRequest {
			totalErrors++
		}

		logrus.WithFields(logrus.Fields{
			"status":        c.Writer.Status(),
			"method":        c.Request.Method,
			"path":          c.Request.URL.Path,
			"latency":       latency,
			"totalRequests": totalRequests,
			"totalErrors":   totalErrors,
		}).Info("Request metrics")
	}
}
