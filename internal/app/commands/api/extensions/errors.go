package extensions

import (
	"database/sql"
	"errors"
	"gotodo/internal/app/commands/api/schemas"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {

	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	response := &schemas.Error{}

	if errors.Is(err, sql.ErrNoRows) {
		code = http.StatusNotFound
		response.Details = err.Error()
	} else {
		var (
			httpError        *echo.HTTPError
			validationErrors validator.ValidationErrors
		)

		if errors.As(err, &httpError) {
			code = httpError.Code
			switch message := httpError.Message.(type) {
			case string:
				response.Details = message
			default:
				response.Details = err.Error()
			}
		} else if errors.As(err, &validationErrors) {
			response.Details = validationErrors.Error()

			items := make(map[string][]string)
			for _, validationError := range validationErrors {
				key := validationError.Field()
				if _, ok := response.ValidationErrors[key]; !ok {
					items[key] = make([]string, 0)
				}
				items[key] = append(items[key], strings.ToUpper(validationError.Tag()))
			}

			if len(items) > 0 {
				response.ValidationErrors = items
			}
		} else {
			response.Details = err.Error()
		}
	}

	// Hide details when not debug
	if !c.Echo().Debug {
		response.Details = ""
	}

	// Send response
	response.Message = http.StatusText(code)
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(code)
	} else {
		err = c.JSON(code, response)
	}

	if err != nil {
		c.Echo().Logger.Error(err)
	}
}
