package middlewares

import (
	"go-vet/infrastructure/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuditMiddleware(db *database.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodDelete {
			auditLog := &database.AuditLog{
				Method:    c.Request.Method,
				Path:      c.Request.URL.Path,
				Timestamp: time.Now(),
				Status:    c.Writer.Status(),
				User:      c.GetString("email"),
			}
			if err := db.Create(auditLog).Error; err != nil {
				logrus.Error("Failed to log audit data:", err)
			}
		}
	}
}
