package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Status  string                 `json:"status"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Error   string                 `json:"error,omitempty"` // Keep for backward compatibility
	Details map[string]interface{} `json:"details,omitempty"`
}

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Status:  "error",
		Code:    "ERROR",
		Message: message,
		Error:   message, // Backward compatibility
		Details: nil,
	})
}

// RespondWithAppError sends an error response based on AppError
func RespondWithAppError(c *gin.Context, err *AppError) {
	response := ErrorResponse{
		Status:  "error",
		Code:    err.Code,
		Message: err.Message,
		Error:   err.Message, // Backward compatibility
		Details: err.Details,
	}
	c.JSON(err.StatusCode(), response)
}

// RespondWithSuccess sends a standardized success response
func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, SuccessResponse{
		Status: "success",
		Data:   data,
	})
}

// RespondWithCreated sends a 201 Created response
func RespondWithCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, SuccessResponse{
		Status: "success",
		Data:   data,
	})
}

// RespondWithNoContent sends a 204 No Content response
func RespondWithNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// RespondWithData sends data directly (for backward compatibility)
func RespondWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// RespondWithValidationError sends a validation error response
func RespondWithValidationError(c *gin.Context, errors map[string]string) {
	response := ErrorResponse{
		Status:  "error",
		Code:    "VALIDATION_ERROR",
		Message: "validation failed",
		Details: make(map[string]interface{}),
	}
	for field, msg := range errors {
		response.Details[field] = msg
	}
	c.JSON(http.StatusBadRequest, response)
}
