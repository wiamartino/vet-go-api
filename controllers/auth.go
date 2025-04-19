package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	jwtUtils "go-vet/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	service *application.UserService
}

func NewAuthController(service *application.UserService) *AuthController {
	return &AuthController{service: service}
}

// Register creates a new user account
func (ctrl *AuthController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid input data",
		})
		return
	}

	// Validate if the email is already registered
	if _, err := ctrl.service.FindByEmail(user.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": "error",
			"error":  "Email is already registered",
		})
		return
	}

	if err := ctrl.service.Register(user); err != nil {
		logrus.WithError(err).Error("Failed to register user")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
	})
}

// Login authenticates a user and provides a JWT token
func (ctrl *AuthController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  "Invalid credentials format",
		})
		return
	}

	user, err := ctrl.service.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"error":  "Invalid email or password",
		})
		return
	}

	// Generate JWT token
	token, err := jwtUtils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Authentication failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"token":      token,
			"user_id":    user.ID,
			"email":      user.Email,
			"name":       user.Name,
			"role":       user.Role,
			"created_at": user.CreatedAt,
		},
	})
}

// RefreshToken generates a new token if the current one is valid
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"error":  "Valid authorization token required",
		})
		return
	}

	tokenString := authHeader[7:]
	newToken, err := jwtUtils.RefreshToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "error",
			"error":  "Invalid token: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"token": newToken,
		},
	})
}
