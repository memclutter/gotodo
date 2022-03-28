package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/models"
)

type JwtClaims struct {
	jwt.StandardClaims
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func GetAuthUser(c echo.Context) *models.User {
	authUserValue := c.Get("auth.user")
	if authUserValue == nil {
		return nil
	}
	return authUserValue.(*models.User)
}
