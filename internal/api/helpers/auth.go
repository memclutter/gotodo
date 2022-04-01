package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"time"
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

func CreateTokens(user models.User) (schemas.AuthBaseResponse, error) {
	var err error
	response := schemas.AuthBaseResponse{User: user}
	dateCreated := time.Now().UTC()
	response.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(30 * 24 * time.Hour).Unix(),
		},
		ID:    user.ID,
		Email: user.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return response, fmt.Errorf("create tokens error: %v", err)
	}
	response.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(30 * 24 * time.Hour).Unix(),
		},
		ID:    user.ID,
		Email: user.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return response, fmt.Errorf("create tokens error: %v", err)
	}
	return response, nil
}
