package entities

type RegistrationRequest struct {
	Name   string `json:"name" binding:"required"`
	PAN    string `json:"pan" binding:"required,pancard"`
	Number string `json:"number" binding:"required,numeric,mobilenumber"`
	Email  string `json:"email" binding:"required,email"`
}
