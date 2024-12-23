package service

import (
	"codingquestions/registration/database"
	"codingquestions/registration/entities"
	"errors"
)

type RegistrationService struct {
	db database.Database
}

type RegisterService interface {
	ValidateAndSaveRequest(req entities.RegistrationRequest) error
	FetchSaveRequest(pan string) entities.RegistrationRequest
}

func NewRegistrationService(db database.Database) *RegistrationService {
	return &RegistrationService{db: db}
}

func (service *RegistrationService) ValidateAndSaveRequest(req entities.RegistrationRequest) error {
	if req.PAN == "" || req.Number == "" || req.Name == "" || req.Email == "" {
		return errors.New("PAN , number, name ,email cannot be empty")
	}
	return service.db.SaveRegistrationDetails(req.PAN, req.Number, req.Name, req.Email)
}

func (service *RegistrationService) FetchSaveRequest(pan string) entities.RegistrationRequest {
	return service.db.FetchRegistrationDetails(pan)
}
