package schemas

import (
	"github.com/memclutter/gotodo/internal/models"
)

type AuthRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthConfirmation struct {
	Token string `json:"token"`
}

type AuthBaseResponse struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
}

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	AuthBaseResponse
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type AuthRefreshResponse struct {
	AuthBaseResponse
}
