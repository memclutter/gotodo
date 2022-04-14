package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/config"
	"net/http"
)

func NewAuth() echo.MiddlewareFunc {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.Config.Secret),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		Claims:      &helpers.JwtClaims{},
		ContextKey:  "auth",
		ErrorHandler: func(err error) error {
			if he, ok := err.(*echo.HTTPError); ok {
				he.Code = http.StatusUnauthorized
				return he
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "Unauthorized",
				Internal: err,
			}
		},
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return jwtMiddleware(func(c echo.Context) error {
			auth := c.Get("auth")
			if auth == nil {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Unauthorized",
					Internal: errors.New("context key 'auth' is nil"),
				}
			}
			jwtToken, ok := auth.(*jwt.Token)
			if !ok {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Unauthorized",
					Internal: errors.New("invalid jwt token"),
				}
			}
			jwtClaims, ok := jwtToken.Claims.(*helpers.JwtClaims)
			if !ok {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Unauthorized",
					Internal: errors.New("invalid jwt claims"),
				}
			}
			// only access tokens
			if jwtClaims.Audience != "access" {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Unauthorized",
					Internal: errors.New("not access audience"),
				}
			}
			// Set only current user claims
			c.Set("auth.jwtClaims", jwtClaims)
			return next(c)
		})
	}
}
