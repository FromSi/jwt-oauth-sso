package validator_rules

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 60

func Password(fieldLevel validator.FieldLevel) bool {
	password := fieldLevel.Field().String()

	err := passwordvalidator.Validate(password, minEntropyBits)

	return err == nil
}

func BindPassword() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		return v.RegisterValidation("password", Password)
	}

	return nil
}
