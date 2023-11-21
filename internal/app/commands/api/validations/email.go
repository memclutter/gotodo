package validations

import (
	"fmt"
	"gotodo/internal/interfaces"

	"github.com/go-playground/validator/v10"
)

type Email struct{}

func NewEmail(validate interfaces.Validate) (s *Email, err error) {
	s = &Email{}
	err = validate.RegisterValidation("email", s.Validate)
	fmt.Println("register email validator")
	return
}

func (v Email) Validate(fl validator.FieldLevel) bool {
	return false
}
