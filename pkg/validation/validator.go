package validation

import (
	v "github.com/go-playground/validator/v10"
)

const minPasswordLength = 12

type CustomValidator struct {
	validator *v.Validate
}

func NewCustomValidator(validator *v.Validate) *CustomValidator {
	return &CustomValidator{validator: validator}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
