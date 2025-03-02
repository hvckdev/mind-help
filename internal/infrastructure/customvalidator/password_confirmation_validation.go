package customvalidator

import "github.com/go-playground/validator/v10"

func PasswordConfirmationValidation(fl validator.FieldLevel) bool {
	passwordConfirm := fl.Field().String()

	password, ok := fl.Parent().FieldByName("Password").Interface().(string)
	if !ok {
		return false
	}

	return password == passwordConfirm
}
