package extensions

import (
	"encoding/json"
	"errors"
	"gotodo/internal/app/commands/api/schemas"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}

	var he *echo.HTTPError
	ok := errors.As(err, &he)
	if ok {
		if he.Internal != nil {
			var herr *echo.HTTPError
			if ok := errors.As(he.Internal, &herr); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// Issue #1426
	code := he.Code
	response := &schemas.Error{
		Message: http.StatusText(code),
	}

	switch m := he.Message.(type) {
	case string:
		if c.Echo().Debug {
			response.Message = m
			response.Details = err.Error()
		} else {
			response.Message = m
		}
	case json.Marshaler:
		// do nothing - this type knows how to format itself to JSON
	case error:
		response.Message = http.StatusText(code)
		response.Details = m.Error()
	}

	// Send response
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, response)
	}

	if err != nil {
		c.Echo().Logger.Error(err)
	}
}
