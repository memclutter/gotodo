package endpoints

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/memclutter/gotodo/internal/security"
	"net/http"
	"strings"
	"time"
)

// AuthLogin godoc
//
// @Router 			/auth/login/	[post]
// @Summary			Login
// @Description		Login in new sessions
// @Tags			auth
// @Accept			json
// @Produce			json
// @Param			request			body		schemas.AuthLoginRequest	true	"Request data"
// @Success			200	{object}				schemas.AuthLoginResponse
func AuthLogin(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("auth login error: %v", err)
	}
	user := models.User{}
	query := models.DB.NewSelect().Model(&user).
		Where("lower(email) = ?", strings.ToLower(req.Email)).
		Where("status = ?", models.UserStatusActive)
	if err := query.Scan(ctx); err != nil {
		return fmt.Errorf("auth login error: %v", err)
	}
	if !security.PasswordValidate(req.Password, user.PwHash) {
		return c.JSON(http.StatusBadRequest, "Invalid credentials")
	}
	dateCreated := time.Now().UTC()
	session := models.Session{
		UserID:      user.ID,
		Token:       security.TokenMustGenerate(32),
		DateCreated: dateCreated,
		DateExpired: dateCreated.Add(30 * 24 * time.Hour),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, helpers.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(2 * time.Minute).Unix(),
		},
		ID:    user.ID,
		Email: user.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return fmt.Errorf("auth login error: %v", err)
	} else if _, err := models.DB.NewInsert().Model(&session).Exec(ctx); err != nil {
		return fmt.Errorf("auth login error: %v", err)
	}

	return c.JSON(http.StatusOK, schemas.AuthLoginResponse{
		User:         user,
		AccessToken:  token,
		RefreshToken: session,
	})
}

// AuthRefresh godoc
//
// @Router 			/auth/refresh/		[post]
// @Summary 		Refresh
// @Description		Refresh current session
// @Tags			auth
// @Accept			json
// @Produce			json
// @Param			request				body		schemas.AuthRefreshRequest	true	"Request body"
// @Success			200	{object}					schemas.AuthRefreshResponse
func AuthRefresh(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthRefreshRequest{}
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("auth refresh error: %v", err)
	}
	session := models.Session{}
	query := models.DB.NewSelect().Model(&session).
		Relation("User").
		Where("token = ?", req.RefreshToken).
		Where("date_expired > NOW()")
	if err := query.Scan(ctx); err != nil {
		return fmt.Errorf("auth refresh error: %v", err)
	}
	dateCreated := time.Now().UTC()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, helpers.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: dateCreated.Add(2 * time.Minute).Unix(),
		},
		ID:    session.User.ID,
		Email: session.User.Email,
	}).SignedString([]byte(config.Config.Secret))
	if err != nil {
		return fmt.Errorf("auth refresh error: %v", err)
	}

	return c.JSON(http.StatusOK, schemas.AuthRefreshResponse{
		User:        session.User,
		AccessToken: token,
	})
}

func AuthLogout(c echo.Context) error {
	return nil
}

// AuthRegistration godoc
//
// @Router 			/auth/registration/		[post]
// @Summary			Registration
// @Description		Register new user
// @Tags			auth
// @Accept 			json
// @Produce 		json
// @Param			request					body		schemas.AuthRegistrationRequest	true	"Request body"
// @Success			201
func AuthRegistration(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthRegistrationRequest{}
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("auth registration error: %v", err)
	}
	user := models.User{
		Email:  req.Email,
		PwHash: security.PasswordMustGenerate(req.Password),
	}
	if _, err := models.DB.NewInsert().Model(&user).Exec(ctx); err != nil {
		return fmt.Errorf("auth registration error: %v", err)
	}
	confirmation := models.Confirmation{
		UserID: user.ID,
		Token:  security.TokenMustGenerate(32),
	}
	if _, err := models.DB.NewInsert().Model(&confirmation).Exec(ctx); err != nil {
		return fmt.Errorf("auth registration error: %v", err)
	}

	//confirmationLink := strings.Join([]string{strings.TrimRight(config.Config.UrlBase, "/"), "auth", "confirmation", confirmation.Token}, "/")
	//mailMsg := gomail.NewMessage()
	//mailMsg.SetHeader("From", "robot@gotodo")
	//mailMsg.SetHeader("To", user.Email)
	//mailMsg.SetHeader("Subject", "Registration on gotodo")
	//mailMsg.SetBody("text/html", fmt.Sprintf("<h1>Hello</h1><p>Your confirmation link <a href=\"%s\">%s</a>", confirmationLink, confirmationLink))
	//
	//mailDialer := gomail.NewDialer(config.Config.MailHost, config.Config.MailPort, config.Config.MailUsername, config.Config.MailPassword)
	//if err := mailDialer.DialAndSend(mailMsg); err != nil {
	//	return fmt.Errorf("auth registration error: %v", err)
	//}

	return c.NoContent(http.StatusCreated)
}

func AuthInfo(c echo.Context) error {
	return nil
}
