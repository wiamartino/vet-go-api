package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
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
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	// Validate user data
	validator := utils.NewValidator()
	validator.ValidateLengthRange("name", user.Name, 2, 100)
	validator.ValidateEmail("email", user.Email)
	validator.ValidateMinLength("password", user.Password, 8)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	// Validate if the email is already registered
	if _, err := ctrl.service.FindByEmail(user.Email); err == nil {
		utils.RespondWithAppError(c, utils.NewConflictError("Email is already registered"))
		return
	}

	if err := ctrl.service.Register(user); err != nil {
		logrus.WithError(err).Error("Failed to register user")
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewBadRequestError(err.Error()))
		}
		return
	}

	utils.RespondWithCreated(c, gin.H{
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
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	user, err := ctrl.service.Login(credentials.Email, credentials.Password)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewUnauthorizedError("Invalid email or password"))
		return
	}

	// Generate JWT token
	token, err := jwtUtils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate token")
		utils.RespondWithAppError(c, utils.NewInternalError("Authentication failed"))
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"token":      token,
		"user_id":    user.ID,
		"email":      user.Email,
		"name":       user.Name,
		"role":       user.Role,
		"created_at": user.CreatedAt,
	})
}

// RefreshToken generates a new token if the current one is valid
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		utils.RespondWithError(c, http.StatusUnauthorized, "Valid authorization token required")
		return
	}

	tokenString := authHeader[7:]
	newToken, err := jwtUtils.RefreshToken(tokenString)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token: "+err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"token": newToken,
	})
}
