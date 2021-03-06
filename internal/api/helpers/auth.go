package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"strings"
	"time"
)

type JwtClaims struct {
	jwt.StandardClaims
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func GetAuthJwtClaims(c echo.Context) *JwtClaims {
	jwtClaims := c.Get("auth.jwtClaims")
	if jwtClaims == nil {
		return nil
	}
	return jwtClaims.(*JwtClaims)
}

func GetAuthUser(c echo.Context) (*models.User, error) {
	authUserValue := c.Get("auth.user")
	if authUserValue == nil {
		jwtClaims := GetAuthJwtClaims(c)
		if jwtClaims == nil {
			return nil, fmt.Errorf("empty auth.jwtClaims")
		}
		user := models.User{}
		query := models.DB.NewSelect().Model(&user).
			Where("id = ?", jwtClaims.ID).
			Where("lower(email) = ?", strings.ToLower(jwtClaims.Email)).
			Where("status = ?", models.UserStatusActive)
		if err := query.Scan(c.Request().Context()); err != nil {
			return nil, fmt.Errorf("middleware user error: %v", err)
		}
		c.Set("auth.user", &user)
		authUserValue = &user
	}
	return authUserValue.(*models.User), nil
}

func CreateTokens(user models.User) (schemas.AuthBaseResponse, error) {
	var err error
	accessTimeout := 2 * time.Minute
	refreshTimeout := 30 * 24 * time.Hour
	if config.Config.Debug {
		accessTimeout = 2 * time.Hour
	}
	response := schemas.AuthBaseResponse{User: user}
	dateCreated := time.Now().UTC()
	response.AccessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(accessTimeout).Unix(),
			Audience:  "access",
		},
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return response, fmt.Errorf("create tokens error: %v", err)
	}
	response.RefreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(refreshTimeout).Unix(),
			Audience:  "refresh",
		},
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return response, fmt.Errorf("create tokens error: %v", err)
	}
	return response, nil
}
