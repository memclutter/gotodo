package endpoints

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
)

// GroupsList godoc
//
// @Router			/groups/		[get]
// @Summary			List
// @Tags			groups
// @Accept			json
// @Produce			json
// @Success			200				{array}		schemas.GroupsListResponse
// @Failure			500				{object} 	schemas.Error		true	"Server error"
// @Security		ApiHeaderAuth
func GroupsList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	totalCount := 0
	groups := make([]models.Group, 0)
	query := models.DB.NewSelect().Model(&groups).
		ColumnExpr("groups.*").
		Relation("Projects").
		Join("INNER JOIN access AS a ON a.group_id = group.id").
		Where("status = ?", models.GroupStatusActive).
		Where("a.user_id = ?", authJwtClaims.ID).
		OrderExpr("id")
	if totalCount, err = query.ScanAndCount(ctx); err != nil {
		return fmt.Errorf("tasks get error: %v", err)
	}
	return c.JSON(http.StatusOK, schemas.GroupsListResponse{
		TotalCount: totalCount,
		Items:      groups,
	})
}

// GroupsCreate godoc
//
// @Router			/groups/		[post]
// @Summary			Create
// @Description		Create a new group
// @Tags 			groups
// @Accept			json
// @Produce			json
// @Param			request			body		models.Group		true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error		true	"Validation error"
// @Failure			500				{object} 	schemas.Error		true	"Server error"
// @Security		ApiHeaderAuth
func GroupsCreate(c echo.Context) error {
	return nil
}

// GroupsRetrieve godoc
//
// @Router 			/groups/{id}/	[get]
// @Summary			Retrieve
// @Description		Get group detail
// @Tags			groups
// @Accept			json
// @Produce 		json
// @Param			id				path		integer							true	"Group identifier"
// @Success			200				{object}	models.Group
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsRetrieve(c echo.Context) error {
	return nil
}

// GroupsUpdate godoc
//
// @Router			/groups/{id}/	[put]
// @Summary			Update
// @Description		Update group
// @Tags			groups
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Param			request			body		models.Group					true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsUpdate(c echo.Context) error {
	return nil
}

// GroupsDelete godoc
//
// @Router			/groups/{id}/	[delete]
// @Summary			Delete
// @Description		Delete group
// @Tags			groups
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Success			204
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsDelete(c echo.Context) error {
	return nil
}
