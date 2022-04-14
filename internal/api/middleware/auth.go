package middleware

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/config"
	"net/http"
)

func NewAuth() echo.MiddlewareFunc {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  config.Config.Secret,
		TokenLookup: "header:authorization",
		AuthScheme:  "Bearer",
		Claims:      &helpers.JwtClaims{},
		ContextKey:  "auth",
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return jwtMiddleware(func(c echo.Context) error {
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
			// Set only current user claims
			c.Set("auth.jwtClaims", jwtClaims)
			return next(c)
		})
	}
}
