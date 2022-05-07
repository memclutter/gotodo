package endpoints

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
)

// ProfileRetrieve godoc
//
// @Router 			/profile/	[get]
// @Summary			Retrieve
// @Description		Get profile
// @Tags			profile
// @Accept			json
// @Produce 		json
// @Success			200				{object}	models.User
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProfileRetrieve(c echo.Context) error {
	authUser, err := helpers.GetAuthUser(c)
	if err != nil {
		return fmt.Errorf("get auth user error: %v", err)
	}
	return c.JSON(http.StatusOK, authUser)
}

// ProfileUpdate godoc
//
// @Router 			/profile/ 	[put]
// @Summary			Update
// @Tags			profile
// @Accept			json
// @Produce			json
// @Param			request			body		schemas.ProfileUpdate			true	"Request body"
// @Success			200				{object}	models.User
// @Failure			500				{object} 	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProfileUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	authUser, err := helpers.GetAuthUser(c)
	if err != nil {
		return fmt.Errorf("get auth user error: %v", err)
	}
	req := schemas.ProfileUpdate{}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err := c.Validate(&req); err != nil {
		return err
	}
	authUser.FirstName = req.FirstName
	authUser.LastName = req.LastName
	if _, err = models.DB.NewInsert().Model(&authUser).Exec(ctx); err != nil {
		return fmt.Errorf("auth user update error: %v", err)
	}
	return c.JSON(http.StatusCreated, authUser)
}
