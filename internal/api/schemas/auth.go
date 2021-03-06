package schemas

import (
	"github.com/memclutter/gotodo/internal/models"
)

type AuthRegistrationRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email,user_email_exists"`
	Password  string `json:"password" validate:"required"`
}

type AuthConfirmation struct {
	Token string `json:"token" validate:"required"`
}

type AuthBaseResponse struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	AuthBaseResponse
}

type AuthRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type AuthRefreshResponse struct {
	AuthBaseResponse
}
