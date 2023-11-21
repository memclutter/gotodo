package mocks

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
)

type Validate struct {
	mock.Mock
}

func (m *Validate) RegisterTagNameFunc(fn validator.TagNameFunc) {
	m.Called(fn)
}

func (m *Validate) RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	args := m.Called(tag, fn, callValidationEvenIfNull)
	return args.Error(0)
}

func (m *Validate) Struct(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}
