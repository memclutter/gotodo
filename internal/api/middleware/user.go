package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
	"strings"
)

func User(next echo.HandlerFunc) echo.HandlerFunc {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  config.Config.Secret,
		TokenLookup: "header:authorization",
		AuthScheme:  "Bearer",
		Claims:      &helpers.JwtClaims{},
		ContextKey:  "auth",
	})
	return jwtMiddleware(func(c echo.Context) error {
		ctx := c.Request().Context()
		auth := c.Get("auth")
		if auth == nil {
			return c.JSON(http.StatusUnauthorized, "Missing authorization access token")
		}
		jwtToken, ok := auth.(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Invalid authorization access token")
		}
		jwtClaims, ok := jwtToken.Claims.(*helpers.JwtClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Invalid authorization access token")
		}
		user := models.User{}
		query := models.DB.NewSelect().Model(&user).
			Where("id = ?", jwtClaims.ID).
			Where("lower(email) = ?", strings.ToLower(jwtClaims.Email)).
			Where("status = ?", models.UserStatusActive)
		if err := query.Scan(ctx); err != nil {
			return fmt.Errorf("middleware user error: %v", err)
		}
		c.Set("auth.user", user)
		return next(c)
	})
}
