package interfaces

import "github.com/go-playground/validator/v10"

type Validate interface {
	RegisterTagNameFunc(fn validator.TagNameFunc)
	RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
	Struct(i interface{}) error
}
