package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
)

// Validator provides validation utilities
type Validator struct {
	Errors map[string]string
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// IsValid returns true if there are no validation errors
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

// AddError adds a validation error
func (v *Validator) AddError(field, message string) {
	v.Errors[field] = message
}

// ValidateEmail validates email format
func (v *Validator) ValidateEmail(field, email string) bool {
	if strings.TrimSpace(email) == "" {
		v.AddError(field, "email is required")
		return false
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		v.AddError(field, "invalid email format")
		return false
	}
	return true
}

// ValidatePhone validates phone number format (basic)
func (v *Validator) ValidatePhone(field, phone string) bool {
	if strings.TrimSpace(phone) == "" {
		v.AddError(field, "phone is required")
		return false
	}
	// Remove common formatting characters
	cleaned := regexp.MustCompile(`[\s\-\+\(\)]+`).ReplaceAllString(phone, "")
	if len(cleaned) < 10 || len(cleaned) > 15 {
		v.AddError(field, "phone must be between 10 and 15 digits")
		return false
	}
	if !regexp.MustCompile(`^[\d\+]+$`).MatchString(cleaned) {
		v.AddError(field, "phone must contain only digits and optional formatting")
		return false
	}
	return true
}

// ValidateRequired validates that a string field is not empty
func (v *Validator) ValidateRequired(field, value string) bool {
	if strings.TrimSpace(value) == "" {
		v.AddError(field, field+" is required")
		return false
	}
	return true
}

// ValidateMinLength validates minimum string length
func (v *Validator) ValidateMinLength(field, value string, minLength int) bool {
	if len(strings.TrimSpace(value)) < minLength {
		v.AddError(field, fmt.Sprintf("%s must be at least %d characters", field, minLength))
		return false
	}
	return true
}

// ValidateMaxLength validates maximum string length
func (v *Validator) ValidateMaxLength(field, value string, maxLength int) bool {
	if len(value) > maxLength {
		v.AddError(field, fmt.Sprintf("%s must not exceed %d characters", field, maxLength))
		return false
	}
	return true
}

// ValidateLengthRange validates string length is within range
func (v *Validator) ValidateLengthRange(field, value string, minLength, maxLength int) bool {
	length := len(strings.TrimSpace(value))
	if length == 0 {
		v.AddError(field, field+" is required")
		return false
	}
	if length < minLength {
		v.AddError(field, fmt.Sprintf("%s must be at least %d characters", field, minLength))
		return false
	}
	if length > maxLength {
		v.AddError(field, fmt.Sprintf("%s must not exceed %d characters", field, maxLength))
		return false
	}
	return true
}

// ValidateDate validates ISO 8601 date format
func (v *Validator) ValidateDate(field, date string) bool {
	if strings.TrimSpace(date) == "" {
		v.AddError(field, field+" is required")
		return false
	}
	// Basic ISO 8601 validation (YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DD)
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}(T\d{2}:\d{2}:\d{2}(Z|[\+\-]\d{2}:\d{2})?)?$`)
	if !datePattern.MatchString(date) {
		v.AddError(field, field+" must be a valid date (ISO 8601 format)")
		return false
	}
	return true
}

// ValidateNumericID validates numeric ID field
func (v *Validator) ValidateNumericID(field string, value uint) bool {
	if value == 0 {
		v.AddError(field, field+" must be a valid numeric ID")
		return false
	}
	return true
}

// ValidateInSlice validates that value is in the provided slice
func (v *Validator) ValidateInSlice(field, value string, validValues []string) bool {
	for _, valid := range validValues {
		if value == valid {
			return true
		}
	}
	v.AddError(field, field+" must be one of: "+strings.Join(validValues, ", "))
	return false
}

// ValidateMinValue validates numeric minimum value
func (v *Validator) ValidateMinValue(field string, value float64, minValue float64) bool {
	if value < minValue {
		v.AddError(field, fmt.Sprintf("%s must be at least %g", field, minValue))
		return false
	}
	return true
}

// ValidateMaxValue validates numeric maximum value
func (v *Validator) ValidateMaxValue(field string, value float64, maxValue float64) bool {
	if value > maxValue {
		v.AddError(field, fmt.Sprintf("%s must not exceed %g", field, maxValue))
		return false
	}
	return true
}
