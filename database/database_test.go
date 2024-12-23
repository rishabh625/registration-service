package database_test

import (
	"codingquestions/registration/database"
	"codingquestions/registration/entities"
	"reflect"
	"testing"
)

func TestInMemoryDatabaseSaveAndRetrieve(t *testing.T) {
	db := database.NewInMemoryDatabase()
	pan := "ABCDE1234F"
	number := "9876543210"
	name := "John Doe"
	email := "john.doe@example.com"

	err := db.SaveRegistrationDetails(pan, number, name, email)
	if err != nil {
		t.Fatalf("Failed to save registration details: %v", err)
	}

	// Verify the data was saved
	savedDetails := db.FetchRegistrationDetails(pan)
	if savedDetails.PAN != pan {
		t.Errorf("Expected PAN %s, got %s", pan, savedDetails.PAN)
	}
	if savedDetails.Number != number {
		t.Errorf("Expected Number %s, got %s", number, savedDetails.Number)
	}
	if savedDetails.Name != name {
		t.Errorf("Expected Name %s, got %s", name, savedDetails.Name)
	}
	if savedDetails.Email != email {
		t.Errorf("Expected Email %s, got %s", email, savedDetails.Email)
	}

	// Test FetchRegistrationDetails for non-existent PAN
	nonExistentPAN := "ZZZZZ9999Z"
	nonExistentDetails := db.FetchRegistrationDetails(nonExistentPAN)
	if !reflect.DeepEqual(nonExistentDetails, entities.RegistrationRequest{}) {
		t.Errorf("Expected empty details for non-existent PAN, got %+v", nonExistentDetails)
	}
}
