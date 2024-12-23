package database

import (
	"codingquestions/registration/entities"
	"fmt"
)

type Database interface {
	SaveRegistrationDetails(pan, number, name, email string) error
	FetchRegistrationDetails(pan string) entities.RegistrationRequest
}

type InMemoryDatabase struct {
	user map[string]entities.RegistrationRequest
}

func NewInMemoryDatabase() *InMemoryDatabase {
	db := make(map[string]entities.RegistrationRequest)
	return &InMemoryDatabase{
		user: db,
	}
}

func (db *InMemoryDatabase) SaveRegistrationDetails(pan, number, name, email string) error {
	if _, ok := db.user[pan]; !ok {
		db.user[pan] = entities.RegistrationRequest{
			Name:   name,
			PAN:    pan,
			Number: number,
			Email:  email,
		}
	}
	fmt.Printf("Saved to DB - PAN: %s, Number: %s\n Name : %s Email: %s", pan, number, name, email)
	return nil
}

func (db *InMemoryDatabase) FetchRegistrationDetails(pan string) entities.RegistrationRequest {
	return db.user[pan]
}
