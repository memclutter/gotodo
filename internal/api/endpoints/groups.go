package endpoints

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
	"github.com/uptrace/bun"
	"net/http"
	"strconv"
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
		ColumnExpr("g.*").
		Relation("Projects").
		Relation("Members").
		Relation("Members.User").
		Join("INNER JOIN access AS a ON a.group_id = g.id").
		Where("status = ?", models.GroupStatusActive).
		Where("a.user_id = ?", authJwtClaims.ID).
		OrderExpr("id")
	if totalCount, err = query.ScanAndCount(ctx); err != nil {
		return fmt.Errorf("groups get error: %v", err)
	}

	// @FIXME convert nil -> [] in projects
	for i := range groups {
		if groups[i].Projects == nil {
			groups[i].Projects = make([]*models.Project, 0)
		}
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
// @Param			request			body		schemas.GroupsRequest		true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object} 	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	req := schemas.GroupsRequest{}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err = c.Validate(&req); err != nil {
		return err
	}

	group := models.Group{
		Name:     req.Name,
		Status:   models.GroupStatusActive,
		Projects: make([]*models.Project, 0),
		Members:  make([]*models.Access, 0),
	}

	// Create group + access in tx
	if err := models.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(&group).Exec(ctx); err != nil {
			return fmt.Errorf("insert group error: %v", err)
		}
		group.Members = append(group.Members, &models.Access{
			UserID:  authJwtClaims.ID,
			GroupID: group.ID,
			Role:    models.AccessRoleAdmin,
		})

		if req.Members != nil {
			for _, member := range req.Members {
				group.Members = append(group.Members, &models.Access{
					GroupID: group.ID,
					UserID:  member.UserID,
					Role:    member.Role,
				})
			}
		}

		if _, err := tx.NewInsert().Model(&group.Members).Exec(ctx); err != nil {
			return fmt.Errorf("insert group members error: %v", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("group create error: %v", err)
	}

	return c.JSON(http.StatusCreated, group)
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
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := models.Group{ID: id}
	query := models.DB.NewSelect().Model(&group).
		ColumnExpr("g.*").
		Relation("Projects").
		Relation("Members").
		Relation("Members.User").
		Join("INNER JOIN access AS a ON a.group_id = g.id").
		Where("status = ?", models.GroupStatusActive).
		Where("a.user_id = ?", authJwtClaims.ID)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	// @FIXME replace nil -> [] in projects
	if group.Projects == nil {
		group.Projects = make([]*models.Project, 0)
	}

	return c.JSON(http.StatusOK, group)
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
// @Param			request			body		schemas.GroupsRequest			true	"Request body"
// @Success			200				{object}	models.Group
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func GroupsUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := models.Group{ID: id}
	query := models.DB.NewSelect().Model(&group).
		ColumnExpr("g.*").
		Relation("Projects").
		Relation("Members").
		Relation("Members.User").
		Join("INNER JOIN access AS a ON a.group_id = g.id").
		Where("status = ?", models.GroupStatusActive).
		Where("a.user_id = ?", authJwtClaims.ID).
		Where("a.role = ?", models.AccessRoleAdmin)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	// Parse request and update model
	req := schemas.GroupsRequest{}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err = c.Validate(&req); err != nil {
		return err
	}

	group.Name = req.Name

	// Create group + access in tx
	if err := models.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewUpdate().Model(&group).Exec(ctx); err != nil {
			return fmt.Errorf("update group error: %v", err)
		}
		// @TODO update members
		return nil
	}); err != nil {
		return fmt.Errorf("group create error: %v", err)
	}

	return c.JSON(http.StatusOK, group)
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
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := models.Group{ID: id}
	query := models.DB.NewSelect().Model(&group).
		Join("INNER JOIN access AS a ON a.group_id = g.id").
		Where("status = ?", models.GroupStatusActive).
		Where("a.user_id = ?", authJwtClaims.ID).
		Where("a.role = ?", models.AccessRoleAdmin)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	// Soft delete
	group.Status = models.GroupStatusDeleted
	if _, err := models.DB.NewUpdate().Model(&group).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("group delete error: %v", err)
	}

	return c.NoContent(http.StatusNoContent)
}
