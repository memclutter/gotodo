package helpers

import (
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/models"
)

func GetAuthUser(c echo.Context) *models.User {
	authUserValue := c.Get("auth.user")
	if authUserValue == nil {
		return nil
	}
	return authUserValue.(*models.User)
}
