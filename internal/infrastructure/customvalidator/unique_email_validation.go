package customvalidator

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"mind-help/internal/domain/entity"
)

type UniqueEmailValidator struct {
	DB *pg.DB
}

func (v *UniqueEmailValidator) UniqueEmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	var user entity.User
	err := v.DB.Model(&user).Where("email = ?", email).Select()

	return err != nil
}
