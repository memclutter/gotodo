package utils

import (
	"gotodo/internal/interfaces"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func NewValidate(validate *validator.Validate) interfaces.Validate { return validate }

type Validator struct {
	validate interfaces.Validate
}

func NewValidator(validate interfaces.Validate) (v *Validator, err error) {
	v = &Validator{validate: validate}
	v.validate.RegisterTagNameFunc(v.TagNameFunc)
	return
}

func (v *Validator) Validate(i interface{}) error {
	return v.validate.Struct(i)
}

func (v *Validator) TagNameFunc(field reflect.StructField) string {
	return field.Tag.Get("json")
}
