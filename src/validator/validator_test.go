package validator

import (
	"go-api-app/src/models"
	"strings"
	"testing"
)

// TestValidateStructValid tests validation with valid employee data
func TestValidateStructValid(t *testing.T) {
	validEmployee := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		JobTitle:  "Software Engineer",
		Email:     "john@example.com",
		Gender:    "Male",
	}

	err := ValidateStruct(validEmployee)
	if err != nil {
		t.Errorf("Expected no validation errors, but got: %v", err)
	}
}

// TestValidateStructMissingRequiredFields tests validation with missing required fields
func TestValidateStructMissingRequiredFields(t *testing.T) {
	tests := []struct {
		name      string
		employee  models.Employee
		wantError bool
		fieldName string
	}{
		{
			name: "missing_first_name",
			employee: models.Employee{
				ID:        1,
				FirstName: "",
				LastName:  "Doe",
				JobTitle:  "Engineer",
				Email:     "test@example.com",
				Gender:    "Male",
			},
			wantError: true,
			fieldName: "FirstName",
		},
		{
			name: "missing_email",
			employee: models.Employee{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				JobTitle:  "Engineer",
				Email:     "",
				Gender:    "Male",
			},
			wantError: true,
			fieldName: "Email",
		},
		{
			name: "missing_gender",
			employee: models.Employee{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				JobTitle:  "Engineer",
				Email:     "test@example.com",
				Gender:    "",
			},
			wantError: true,
			fieldName: "Gender",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStruct(tt.employee)
			if tt.wantError && err == nil {
				t.Errorf("Expected validation error for %s, but got none", tt.fieldName)
			}
			if !tt.wantError && err != nil {
				t.Errorf("Did not expect validation error, but got: %v", err)
			}
			if tt.wantError && err != nil {
				// Check if the error contains the expected field
				fieldFound := false
				for _, validErr := range err.Errors {
					if validErr.Field == tt.fieldName {
						fieldFound = true
						break
					}
				}
				if !fieldFound {
					t.Errorf("Expected error for field %s, but got errors for: %v", tt.fieldName, err.Errors)
				}
			}
		})
	}
}

// TestValidateStructInvalidEmail tests email validation
func TestValidateStructInvalidEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		isValid bool
	}{
		{
			name:    "valid_email",
			email:   "john@example.com",
			isValid: true,
		},
		{
			name:    "invalid_email_no_at",
			email:   "johnexample.com",
			isValid: false,
		},
		{
			name:    "invalid_email_no_domain",
			email:   "john@",
			isValid: false,
		},
		{
			name:    "invalid_email_spaces",
			email:   "john doe@example.com",
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employee := models.Employee{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				JobTitle:  "Engineer",
				Email:     tt.email,
				Gender:    "Male",
			}

			err := ValidateStruct(employee)
			if tt.isValid && err != nil {
				t.Errorf("Expected email %s to be valid, but got error: %v", tt.email, err)
			}
			if !tt.isValid && err == nil {
				t.Errorf("Expected email %s to be invalid, but validation passed", tt.email)
			}
		})
	}
}

// TestValidateStructInvalidGender tests gender validation (oneof constraint)
func TestValidateStructInvalidGender(t *testing.T) {
	invalidEmployee := models.Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		JobTitle:  "Engineer",
		Email:     "john@example.com",
		Gender:    "InvalidGender",
	}

	err := ValidateStruct(invalidEmployee)
	if err == nil {
		t.Error("Expected validation error for invalid gender, but got none")
	}
	if err != nil && err.Status != "validation_failed" {
		t.Errorf("Expected status 'validation_failed', but got: %s", err.Status)
	}
}

// TestValidateStructStringLengthConstraints tests min/max length validation
func TestValidateStructStringLengthConstraints(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		isValid   bool
	}{
		{
			name:      "valid_first_name",
			firstName: "John",
			isValid:   true,
		},
		{
			name:      "short_first_name",
			firstName: "",
			isValid:   false,
		},
		{
			name:      "very_long_first_name",
			firstName: "a" + strings.Repeat("a", 100), // 101 characters
			isValid:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employee := models.Employee{
				ID:        1,
				FirstName: tt.firstName,
				LastName:  "Doe",
				JobTitle:  "Engineer",
				Email:     "test@example.com",
				Gender:    "Male",
			}

			err := ValidateStruct(employee)
			if tt.isValid && err != nil {
				t.Errorf("Expected %s to be valid, but got error: %v", tt.firstName, err)
			}
			if !tt.isValid && err == nil {
				t.Errorf("Expected %s to be invalid, but validation passed", tt.firstName)
			}
		})
	}
}

// TestErrorResponseStructure tests the structure of error responses
func TestErrorResponseStructure(t *testing.T) {
	invalidEmployee := models.Employee{
		ID:        1,
		FirstName: "",
		LastName:  "",
		JobTitle:  "Engineer",
		Email:     "invalid-email",
		Gender:    "",
	}

	err := ValidateStruct(invalidEmployee)
	if err == nil {
		t.Fatal("Expected validation errors, but got none")
	}

	// Check response structure
	if err.Status != "validation_failed" {
		t.Errorf("Expected status 'validation_failed', got: %s", err.Status)
	}

	// Check that errors list is not empty
	if len(err.Errors) == 0 {
		t.Error("Expected validation errors, but got empty list")
	}

	// Check structure of individual errors
	for _, validErr := range err.Errors {
		if validErr.Field == "" {
			t.Error("Field name should not be empty")
		}
		if validErr.Message == "" {
			t.Error("Error message should not be empty")
		}
	}
}
