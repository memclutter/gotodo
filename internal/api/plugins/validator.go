package plugins

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Validator struct {
	goPlaygroundValidate *validator.Validate
}

func NewValidator() *Validator {
	goPlaygroundValidate := validator.New()
	goPlaygroundValidate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
	return &Validator{
		goPlaygroundValidate: goPlaygroundValidate,
	}
}

func (v *Validator) Validate(i interface{}) error {
	return v.goPlaygroundValidate.Struct(i)
}
