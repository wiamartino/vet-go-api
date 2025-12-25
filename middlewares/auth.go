package middlewares

import (
	"net/http"
	"os"
	"strings"

	jwtUtils "go-vet/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthMiddleware validates JWT tokens and sets user information in the context
// Can be disabled by setting DISABLE_AUTH=true environment variable
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if auth is disabled via environment variable
		if os.Getenv("DISABLE_AUTH") == "true" {
			logrus.Warn("⚠️  Authentication is DISABLED - DEVELOPMENT MODE ONLY")
			// Set default user context for development
			c.Set("userID", uint(1))
			c.Set("email", "dev@example.com")
			c.Set("role", "admin")
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Authorization header required",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Invalid token format",
			})
			c.Abort()
			return
		}

		tokenString := authHeader[7:] // Remove "Bearer " prefix
		claims, err := jwtUtils.ValidateToken(tokenString)
		if err != nil {
			logrus.WithError(err).Info("Invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Store user information in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleAuthMiddleware checks if the user has the required role
func RoleAuthMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// This middleware should be used after AuthMiddleware
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "error",
				"error":  "Authentication required",
			})
			c.Abort()
			return
		}

		userRole := role.(string)

		// Check if the user's role is in the list of required roles
		hasPermission := false
		for _, requiredRole := range requiredRoles {
			if userRole == requiredRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, gin.H{
				"status": "error",
				"error":  "Insufficient permissions",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
