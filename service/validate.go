package service

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pancard", validatePAN)
		v.RegisterValidation("mobilenumber", validateMobileNumber)
	}
}

func validatePAN(fl validator.FieldLevel) bool {
	panRegex := `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	return regexp.MustCompile(panRegex).MatchString(fl.Field().String())
}

func validateMobileNumber(fl validator.FieldLevel) bool {
	mobileRegex := `^[6-9][0-9]{9}$`
	return regexp.MustCompile(mobileRegex).MatchString(fl.Field().String())
}
