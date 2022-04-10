package validators

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/sirupsen/logrus"
	"strings"
)

func UserEmailExists(ctx context.Context, fl validator.FieldLevel) bool {
	value := fl.Field().String()
	user := models.User{}
	if exists, err := models.DB.NewSelect().Model(&user).Where("lower(email) = ?", strings.ToLower(value)).Exists(ctx); err == sql.ErrNoRows {
		return true
	} else if err != nil {
		logrus.Errorf("user email exists validation error: %v", err)
		return false
	} else {
		return !exists
	}
}
