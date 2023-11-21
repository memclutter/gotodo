package utils

import (
	"gotodo/internal/mocks"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewValidate(t *testing.T) {
	// Arrange
	validate := validator.New()

	// Act
	actual := NewValidate(validate)

	// Asserts
	assert.NotNil(t, actual, "must be correct type")
}

func TestNewValidator(t *testing.T) {
	// Arrange
	validate := &mocks.Validate{}
	validate.On("RegisterTagNameFunc", mock.AnythingOfType("validator.TagNameFunc"))

	// Act
	_, err := NewValidator(validate)

	// Asserts
	require.NoError(t, err, "must be no errors")
	validate.AssertExpectations(t)
}

func TestValidator_Validate(t *testing.T) {
	// Arrange
	data := struct{ a int }{a: 1}
	validate := &mocks.Validate{}
	validate.On("RegisterTagNameFunc", mock.AnythingOfType("validator.TagNameFunc"))
	validate.On("Struct", data).Return(nil)

	// Act
	v, err := NewValidator(validate)
	require.NoError(t, err, "must be no errors")
	err = v.Validate(data)

	// Assert
	assert.Nil(t, err, "must be no error when validate data")
	validate.AssertExpectations(t)
}

func TestValidator_TagNameFunc(t *testing.T) {
	// Arrange
	data := struct {
		Field int `json:"field"`
	}{Field: 1}
	validate := &mocks.Validate{}
	validate.On("RegisterTagNameFunc", mock.AnythingOfType("validator.TagNameFunc"))
	field := reflect.TypeOf(data).Field(0)

	// Act
	v, err := NewValidator(validate)
	require.NoError(t, err, "must be no errors")
	tagName := v.TagNameFunc(field)

	// Assert
	assert.Equal(t, "field", tagName, "must be equal tag name")
	validate.AssertExpectations(t)
}
