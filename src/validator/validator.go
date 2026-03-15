package validator

import (
	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single field validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse represents the overall validation error response
type ErrorResponse struct {
	Status string            `json:"status"`
	Errors []ValidationError `json:"errors"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct validates a struct and returns structured errors
func ValidateStruct(data any) *ErrorResponse {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errors := make([]ValidationError, len(validationErrors))

	for i, fe := range validationErrors {
		errors[i] = ValidationError{
			Field:   fe.Field(),
			Message: getErrorMessage(fe),
		}
	}

	return &ErrorResponse{
		Status: "validation_failed",
		Errors: errors,
	}
}

// getErrorMessage returns human-readable error message based on validation tag
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "max":
		return fe.Field() + " must be at most " + fe.Param() + " characters"
	case "oneof":
		return fe.Field() + " must be one of: " + fe.Param()
	case "number":
		return fe.Field() + " must be a number"
	default:
		return fe.Field() + " failed validation"
	}
}
