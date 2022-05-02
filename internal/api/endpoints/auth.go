package endpoints

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/config"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/memclutter/gotodo/internal/security"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"gopkg.in/gomail.v2"
	"net/http"
	"strings"
)

// AuthRegistration godoc
//
// @Router 			/auth/registration/			[post]
// @Summary			Registration
// @Description		Register new user
// @Tags			auth
// @Accept 			json
// @Produce 		json
// @Param			request						body		schemas.AuthRegistrationRequest	true	"Request body"
// @Success			201
// @Failure			400							{object}	schemas.Error					true	"Validation error"
// @Failure			500							{object}	schemas.Error					true	"Server error"
func AuthRegistration(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthRegistrationRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	} else if err := c.Validate(&req); err != nil {
		return err
	}
	user := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		PwHash:    security.PasswordMustGenerate(req.Password),
		Status:    models.UserStatusPending,
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

	confirmationLink := strings.Join([]string{strings.TrimRight(config.Config.UrlBase, "/"), "confirmation", confirmation.Token}, "/") + "/"
	mailMsg := gomail.NewMessage()
	mailMsg.SetHeader("From", config.Config.DefaultFromMail)
	mailMsg.SetHeader("To", user.Email)
	mailMsg.SetHeader("Subject", "Registration on gotodo")
	mailMsg.SetBody(
		"text/html",
		fmt.Sprintf(
			"<h1>Hello %s %s</h1><p>Your confirmation link <a href=\"%s\">%s</a>",
			user.FirstName,
			user.LastName,
			confirmationLink,
			confirmationLink,
		),
	)

	mailDialer := gomail.NewDialer(config.Config.MailHost, config.Config.MailPort, config.Config.MailUsername, config.Config.MailPassword)
	mailDialer.SSL = config.Config.MailSSL
	if err := mailDialer.DialAndSend(mailMsg); err != nil {
		return fmt.Errorf("auth registration error: %v", err)
	}

	return c.NoContent(http.StatusCreated)
}

// AuthConfirmation godoc
//
// @Router 			/auth/confirmation/ 		[post]
// @Summary			Confirmation
// @Description		User confirmation after registration
// @Tags			auth
// @Accept			json
// @Produce			json
// @Param			request						body		schemas.AuthConfirmation		true	"Request body"
// @Success			204
// @Failure			400							{object}	schemas.Error					true	"Validation error"
// @Failure			500							{object}	schemas.Error					true	"Server error"
func AuthConfirmation(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthConfirmation{}
	if err := c.Bind(&req); err != nil {
		return err
	} else if err := c.Validate(&req); err != nil {
		return err
	}
	logCtx := logrus.WithField("token", req.Token)
	confirmation := models.Confirmation{}
	query := models.DB.NewSelect().Model(&confirmation).Relation("User").Where("token = ?", req.Token)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		logCtx.Info("not found token in db")
		return c.NoContent(http.StatusBadRequest)
	} else if err != nil {
		return fmt.Errorf("auth confimration error: %v", err)
	}
	if confirmation.IsExpired() {
		logCtx.WithField("dateExpired", confirmation.DateExpired).Info("date expired")
		if _, err := models.DB.NewDelete().Model(&confirmation).WherePK().Exec(ctx); err != nil {
			return fmt.Errorf("auth confirmation error: %v", err)
		}
		return c.NoContent(http.StatusBadRequest)
	}
	if err := models.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		user := confirmation.User
		user.Status = models.UserStatusActive
		if _, err := tx.NewUpdate().Model(&user).WherePK().Exec(ctx); err != nil {
			return fmt.Errorf("run in tx error: %v", err)
		}
		if _, err := tx.NewDelete().Model(&confirmation).WherePK().Exec(ctx); err != nil {
			return fmt.Errorf("run in tx error: %v", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("auth confirmation error: %v", err)
	}
	return c.NoContent(http.StatusOK)
}

// AuthLogin godoc
//
// @Router 			/auth/login/				[post]
// @Summary			Login
// @Description		Login in new sessions
// @Tags			auth
// @Accept			json
// @Produce			json
// @Param			request						body		schemas.AuthLoginRequest	true	"Request data"
// @Success			200							{object}	schemas.AuthLoginResponse
// @Failure			400							{object}	schemas.Error					true	"Validation error"
// @Failure			500							{object}	schemas.Error					true	"Server error"
func AuthLogin(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	} else if err := c.Validate(&req); err != nil {
		return err
	}
	user := models.User{}
	query := models.DB.NewSelect().Model(&user).
		Where("lower(email) = ?", strings.ToLower(req.Email)).
		Where("status = ?", models.UserStatusActive)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusBadRequest)
	} else if err != nil {
		return fmt.Errorf("auth login error: %v", err)
	}
	if !security.PasswordValidate(req.Password, user.PwHash) {
		return c.JSON(http.StatusBadRequest, "Invalid credentials")
	}
	baseRes, err := helpers.CreateTokens(user)
	if err != nil {
		return fmt.Errorf("auth login error: %v", err)
	}
	return c.JSON(http.StatusOK, schemas.AuthLoginResponse{
		AuthBaseResponse: baseRes,
	})
}

// AuthRefresh godoc
//
// @Router			/auth/refresh/				[post]
// @Summary			Refresh
// @Description		Refresh exists session
// @Tags			auth
// @Tags			auth
// @Accept			json
// @Produce			json
// @Param			request						body		schemas.AuthRefreshRequest	true	"Request data"
// @Success			200							{object}	schemas.AuthRefreshResponse
// @Failure			400							{object}	schemas.Error					true	"Validation error"
// @Failure			500							{object}	schemas.Error					true	"Server error"
func AuthRefresh(c echo.Context) error {
	ctx := c.Request().Context()
	req := schemas.AuthRefreshRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	} else if err := c.Validate(&req); err != nil {
		return err
	}
	refreshTokenParsed, err := jwt.ParseWithClaims(req.RefreshToken, &helpers.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Secret), nil
	})
	if err != nil {
		c.Logger().Errorf("invalid refresh token: %v", err)
		return c.NoContent(http.StatusBadRequest)
	} else if !refreshTokenParsed.Valid {
		return c.NoContent(http.StatusBadRequest)
	}
	refreshTokenClaims, ok := refreshTokenParsed.Claims.(*helpers.JwtClaims)
	if !ok {
		logrus.Infof("invalid claims %T", refreshTokenParsed.Claims)
		return c.NoContent(http.StatusBadRequest)
	}
	user := models.User{}
	query := models.DB.NewSelect().Model(&user).
		Where("id = ?", refreshTokenClaims.ID).
		Where("lower(email) = ?", strings.ToLower(refreshTokenClaims.Email)).
		Where("status = ?", models.UserStatusActive)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		c.Logger().Errorf("not found user: %v", refreshTokenClaims)
		return c.NoContent(http.StatusBadRequest)
	} else if err != nil {
		return fmt.Errorf("auth refresh error: %v", err)
	}
	baseRes, err := helpers.CreateTokens(user)
	if err != nil {
		return fmt.Errorf("auth refresh error: %v", err)
	}
	return c.JSON(http.StatusOK, schemas.AuthLoginResponse{
		AuthBaseResponse: baseRes,
	})
}
