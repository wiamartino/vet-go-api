package utils

import "net/http"

// ErrorType defines the type of error that occurred
type ErrorType int

const (
	ValidationError ErrorType = iota
	NotFoundError
	ConflictError
	UnauthorizedError
	ForbiddenError
	InternalError
	BadRequestError
)

// AppError represents an application error with proper HTTP status mapping
type AppError struct {
	Type    ErrorType
	Message string
	Code    string
	Details map[string]interface{}
}

// Error implements the error interface
func (e *AppError) Error() string {
	return e.Message
}

// StatusCode returns the appropriate HTTP status code for this error
func (e *AppError) StatusCode() int {
	switch e.Type {
	case ValidationError:
		return http.StatusBadRequest
	case NotFoundError:
		return http.StatusNotFound
	case ConflictError:
		return http.StatusConflict
	case UnauthorizedError:
		return http.StatusUnauthorized
	case ForbiddenError:
		return http.StatusForbidden
	case BadRequestError:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// NewValidationError creates a validation error
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ValidationError,
		Message: message,
		Code:    "VALIDATION_ERROR",
		Details: make(map[string]interface{}),
	}
}

// NewNotFoundError creates a not found error
func NewNotFoundError(resource string) *AppError {
	return &AppError{
		Type:    NotFoundError,
		Message: resource + " not found",
		Code:    "NOT_FOUND",
		Details: make(map[string]interface{}),
	}
}

// NewConflictError creates a conflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Type:    ConflictError,
		Message: message,
		Code:    "CONFLICT",
		Details: make(map[string]interface{}),
	}
}

// NewUnauthorizedError creates an unauthorized error
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:    UnauthorizedError,
		Message: message,
		Code:    "UNAUTHORIZED",
		Details: make(map[string]interface{}),
	}
}

// NewForbiddenError creates a forbidden error
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Type:    ForbiddenError,
		Message: message,
		Code:    "FORBIDDEN",
		Details: make(map[string]interface{}),
	}
}

// NewBadRequestError creates a bad request error
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Type:    BadRequestError,
		Message: message,
		Code:    "BAD_REQUEST",
		Details: make(map[string]interface{}),
	}
}

// NewInternalError creates an internal server error
func NewInternalError(message string) *AppError {
	return &AppError{
		Type:    InternalError,
		Message: "An internal error occurred",
		Code:    "INTERNAL_ERROR",
		Details: make(map[string]interface{}),
	}
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details map[string]interface{}) *AppError {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	for k, v := range details {
		e.Details[k] = v
	}
	return e
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// AsAppError converts an error to AppError if possible
func AsAppError(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return NewInternalError(err.Error())
}
