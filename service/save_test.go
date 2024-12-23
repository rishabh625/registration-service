package service

import (
	"codingquestions/registration/database"
	"codingquestions/registration/entities"
	"errors"
	"testing"
)

type MockDatabase struct {
	// Simulate a database in memory
	data map[string]entities.RegistrationRequest
}

func (db *MockDatabase) SaveRegistrationDetails(pan, number, name, email string) error {
	// Simulate saving data
	if _, exists := db.data[pan]; exists {
		return errors.New("duplicate PAN")
	}
	db.data[pan] = entities.RegistrationRequest{
		PAN:    pan,
		Number: number,
		Name:   name,
		Email:  email,
	}
	return nil
}

func TestValidateAndSaveRequest(t *testing.T) {
	mockDB := database.NewMockDatabase()
	service := NewRegistrationService(mockDB)

	tests := []struct {
		req       entities.RegistrationRequest
		expectErr bool
	}{
		{
			req:       entities.RegistrationRequest{PAN: "ABCDE1234F", Number: "9876543210", Name: "John Doe", Email: "john@example.com"},
			expectErr: false, // Valid input
		},
		{
			req:       entities.RegistrationRequest{PAN: "", Number: "9876543210", Name: "John Doe", Email: "john@example.com"},
			expectErr: true, // Invalid PAN
		},
		{
			req:       entities.RegistrationRequest{PAN: "ABCDE1234F", Number: "", Name: "John Doe", Email: "john@example.com"},
			expectErr: true, // Invalid Number
		},
		{
			req:       entities.RegistrationRequest{PAN: "ABCDE1234F", Number: "9876543210", Name: "", Email: "john@example.com"},
			expectErr: true, // Invalid Name
		},
		{
			req:       entities.RegistrationRequest{PAN: "ABCDE1234F", Number: "9876543210", Name: "John Doe", Email: ""},
			expectErr: true, // Invalid Email
		},
	}

	for _, test := range tests {
		t.Run(test.req.PAN, func(t *testing.T) {
			err := service.ValidateAndSaveRequest(test.req)
			if (err != nil) != test.expectErr {
				t.Errorf("expected error: %v, got: %v", test.expectErr, err)
			}
		})
	}
}

func TestFetchSaveRequest(t *testing.T) {
	mockDB := database.NewMockDatabase()
	service := NewRegistrationService(mockDB)

	// Test case where the PAN is found
	mockDB.SaveRegistrationDetails("ABCDE1234F", "9876543210", "John Doe", "john@example.com")

	tests := []struct {
		pan       string
		expectErr bool
	}{
		{
			pan:       "ABCDE1234F", // Valid PAN
			expectErr: false,
		},
		{
			pan:       "INVALID1234", // Invalid PAN (not in DB)
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.pan, func(t *testing.T) {
			result := service.FetchSaveRequest(test.pan)
			if (result.PAN == "" && !test.expectErr) || (result.PAN != "" && test.expectErr) {
				t.Errorf("expected error: %v, got: %v", test.expectErr, result.PAN)
			}
		})
	}
}
