package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateNPWP(fl validator.FieldLevel) bool {
	npwp := fl.Field().String()

	if npwp == "" {
		return true
	}

	match, _ := regexp.MatchString(`^\d{15}$`, npwp)
	return match
}
