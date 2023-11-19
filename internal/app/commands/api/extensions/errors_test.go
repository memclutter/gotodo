package extensions

import (
	"errors"
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/mocks"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestErrorHandler(t *testing.T) {
	// Arrange
	err := echo.NewHTTPError(http.StatusInternalServerError, "Test error")
	e := echo.New()

	c := &mocks.EchoContext{}
	c.On("Response").Return(&echo.Response{Committed: false})
	c.On("Echo").Return(e)
	c.On("Request").Return(&http.Request{Method: http.MethodPost})
	c.On("JSON", http.StatusInternalServerError, &schemas.Error{Message: "Test error"}).Return(nil)

	// Act
	ErrorHandler(err, c)

	// Assert
	c.AssertExpectations(t)
}

func TestErrorHandlerForGoError(t *testing.T) {
	// Arrange
	err := errors.New("test error")
	e := echo.New()
	e.Debug = true
	c := &mocks.EchoContext{}
	c.On("Response").Return(&echo.Response{Committed: false})
	c.On("Echo").Return(e)
	c.On("Request").Return(&http.Request{Method: http.MethodPost})
	c.On("JSON", http.StatusInternalServerError, &schemas.Error{Message: "Internal Server Error", Details: "test error"}).Return(nil)

	// Act
	ErrorHandler(err, c)

	// Assert
	c.AssertExpectations(t)
}
