package helpers

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
