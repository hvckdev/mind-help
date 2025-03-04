package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"mind-help/internal/infrastructure/repository"
)

type UniqueEmailValidator struct {
	R repository.UserRepository
}

func (v *UniqueEmailValidator) UniqueEmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	_, err := v.R.GetByEmail(email)

	return err != nil
}
