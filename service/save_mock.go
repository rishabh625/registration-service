package service

import (
	"codingquestions/registration/database"
	"codingquestions/registration/entities"
)

type RegistrationServiceMock struct {
	db database.Database
}

type RegisterServiceMock interface {
	ValidateAndSaveRequest(req entities.RegistrationRequest) error
	FetchSaveRequest(pan string) entities.RegistrationRequest
}

func NewRegistrationServiceMock() *RegistrationService {
	db := database.NewInMemoryDatabase()
	return &RegistrationService{
		db: db,
	}
}

func (service *RegistrationServiceMock) ValidateAndSaveRequest(req entities.RegistrationRequest) error {
	return nil
}

func (service *RegistrationServiceMock) FetchSaveRequest(pan string) entities.RegistrationRequest {
	return entities.RegistrationRequest{
		PAN: pan,
	}
}
