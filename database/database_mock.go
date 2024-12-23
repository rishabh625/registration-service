package database

import (
	"codingquestions/registration/entities"
	"errors"
)

type MockDatabase struct {
	// Simulate a database in memory
	data map[string]entities.RegistrationRequest
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		data: make(map[string]entities.RegistrationRequest),
	}
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

func (db *MockDatabase) FetchRegistrationDetails(pan string) entities.RegistrationRequest {
	// Simulate fetching data
	return db.data[pan]
}
