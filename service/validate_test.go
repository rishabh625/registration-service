package service

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

// Test for PAN validation
func TestValidatePAN(t *testing.T) {
	// Initialize validator instance
	validate := validator.New()

	// Define test cases for PAN validation
	tests := []struct {
		pan      string
		expected bool
	}{
		{"ABCDE1234F", true},   // Valid PAN
		{"ABCDE1234", false},   // Invalid PAN (missing last character)
		{"ABCDE1234FA", false}, // Invalid PAN (last character not uppercase letter)
		{"1234567890A", false}, // Invalid PAN (doesn't start with uppercase letters)
		{"abcdE1234F", false},  // Invalid PAN (first characters are lowercase)
	}

	// Register the custom PAN validator
	validate.RegisterValidation("pan", validatePAN)

	// Run the test cases
	for _, test := range tests {
		t.Run(test.pan, func(t *testing.T) {
			err := validate.Var(test.pan, "pan")
			isValid := err == nil
			if isValid != test.expected {
				t.Errorf("expected %v for PAN %s, got %v", test.expected, test.pan, isValid)
			}
		})
	}
}

// Test for Mobile Number validation
func TestValidateMobileNumber(t *testing.T) {
	// Initialize validator instance
	validate := validator.New()

	// Define test cases for Mobile Number validation
	tests := []struct {
		mobile   string
		expected bool
	}{
		{"9876543210", true},   // Valid mobile number
		{"8888888888", true},   // Valid mobile number (starts with 8)
		{"0123456789", false},  // Invalid mobile number (starts with 0)
		{"98765432", false},    // Invalid mobile number (less than 10 digits)
		{"abcdefg1234", false}, // Invalid mobile number (non-numeric characters)
	}

	// Register the custom mobile number validator
	validate.RegisterValidation("mobile", validateMobileNumber)

	// Run the test cases
	for _, test := range tests {
		t.Run(test.mobile, func(t *testing.T) {
			err := validate.Var(test.mobile, "mobile")
			isValid := err == nil
			if isValid != test.expected {
				t.Errorf("expected %v for mobile number %s, got %v", test.expected, test.mobile, isValid)
			}
		})
	}
}
