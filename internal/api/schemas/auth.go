package schemas

import (
	"github.com/memclutter/gotodo/internal/models"
)

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	User         models.User    `json:"user"`
	AccessToken  string         `json:"accessToken"`
	RefreshToken models.Session `json:"refreshToken"`
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthRefreshResponse struct {
	User        models.User `json:"user"`
	AccessToken string      `json:"accessToken"`
}

type AuthRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
