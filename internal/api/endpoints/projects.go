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

// ProjectsList godoc
//
// @Router			/projects/		[get]
// @Summary			List
// @Tags			projects
// @Accept			json
// @Produce			json
// @Param			groupId			query		schemas.ProjectsListRequest		false	"List filter"
// @Success			200				{array}		schemas.ProjectsListResponse
// @Failure			500				{object} 	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	req := schemas.ProjectsListRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	totalCount := 0
	projects := make([]models.Project, 0)
	query := models.DB.NewSelect().Model(&projects).
		ColumnExpr("p.*").
		Relation("Group").
		Relation("Members").
		Relation("Members.User").
		Relation("Statuses").
		Join("LEFT JOIN access AS pa ON pa.project_id = p.id").
		Join("LEFT JOIN access AS ga ON ga.group_id = p.group_id").
		Where("status = ?", models.ProjectStatusActive).
		Where("pa.user_id = ? OR ga.user_id = ?", authJwtClaims.ID, authJwtClaims.ID).
		OrderExpr("id")
	if totalCount, err = query.ScanAndCount(ctx); err != nil {
		return fmt.Errorf("projects get error: %v", err)
	}

	return c.JSON(http.StatusOK, schemas.ProjectsListResponse{
		TotalCount: totalCount,
		Items:      projects,
	})
}

// ProjectsCreate godoc
//
// @Router			/projects/		[post]
// @Summary			Create
// @Description		Create a new project
// @Tags 			projects
// @Accept			json
// @Produce			json
// @Param			request			body		schemas.ProjectsRequest		true	"Request body"
// @Success			200				{object}	models.Project
// @Failure			400				{object}	schemas.Error				true	"Validation error"
// @Failure			500				{object} 	schemas.Error				true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	req := schemas.ProjectsRequest{}
	access := models.Access{UserID: authJwtClaims.ID, Role: models.AccessRoleAdmin}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err = c.Validate(&req); err != nil {
		return err
	}
	project := models.Project{
		Name:    req.Name,
		GroupID: req.GroupID,
		Status:  models.GroupStatusActive,
		Members: make([]*models.Access, 0),
	}

	// Check group access
	group := models.Group{ID: project.GroupID}
	if err := models.DB.NewSelect().Model(&group).Relation("Members").WherePK().Scan(ctx); err != nil {
		return fmt.Errorf("get project group error: %v", err)
	} else if can, err := group.Can(authJwtClaims.ID, []string{}); err != nil {
		return fmt.Errorf("check project group access error: %v", err)
	} else if !can {
		return c.NoContent(http.StatusForbidden)
	}
	project.Group = group

	// Create project + access in tx
	if err := models.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(&project).Exec(ctx); err != nil {
			return fmt.Errorf("insert project error: %v", err)
		}
		access.ProjectID = project.ID
		project.Members = append(project.Members, &access)

		if req.Members != nil {
			for _, member := range req.Members {
				project.Members = append(project.Members, &models.Access{
					ProjectID: project.ID,
					UserID:    member.UserID,
					Role:      member.Role,
				})
			}
		}

		if _, err := tx.NewInsert().Model(&project.Members).Exec(ctx); err != nil {
			return fmt.Errorf("insert members error: %v", err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("project create error: %v", err)
	}

	return c.JSON(http.StatusCreated, project)
}

// ProjectsRetrieve godoc
//
// @Router 			/projects/{id}/	[get]
// @Summary			Retrieve
// @Description		Get project detail
// @Tags			projects
// @Accept			json
// @Produce 		json
// @Param			id				path		integer							true	"Group identifier"
// @Success			200				{object}	models.Project
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsRetrieve(c echo.Context) error {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	project := models.Project{ID: id}
	query := models.DB.NewSelect().Model(&project).
		ColumnExpr("p.*").
		Relation("Group").
		Relation("Members").
		Relation("Members.User").
		Join("LEFT JOIN access AS pa ON pa.project_id = p.id").
		Join("LEFT JOIN access AS ga ON ga.group_id = p.group_id").
		WherePK().
		Where("status = ?", models.ProjectStatusActive).
		Where("pa.user_id = ? OR ga.user_id = ?", authJwtClaims.ID, authJwtClaims.ID)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	return c.JSON(http.StatusOK, project)
}

// ProjectsUpdate godoc
//
// @Router			/projects/{id}/	[put]
// @Summary			Update
// @Description		Update project
// @Tags			projects
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Param			request			body		schemas.ProjectsRequest			true	"Request body"
// @Success			200				{object}	models.Project
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsUpdate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	project := models.Project{ID: id}
	query := models.DB.NewSelect().Model(&project).
		ColumnExpr("p.*").
		Relation("Group").
		Relation("Group.Members").
		Relation("Members").
		Relation("Members.User").
		Join("LEFT JOIN access AS pa ON pa.project_id = p.id").
		Join("LEFT JOIN access AS ga ON ga.group_id = p.group_id").
		WherePK().
		Where("status = ?", models.ProjectStatusActive).
		Where("pa.user_id = ? OR ga.user_id = ?", authJwtClaims.ID, authJwtClaims.ID)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	// Parse request and update model
	req := schemas.ProjectsRequest{}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err = c.Validate(&req); err != nil {
		return err
	}
	project.Name = req.Name
	project.GroupID = req.GroupID

	// Check group access
	if can, err := project.Group.Can(authJwtClaims.ID, []string{}); err != nil {
		return fmt.Errorf("check project group access error: %v", err)
	} else if !can {
		return c.NoContent(http.StatusForbidden)
	}

	// Create project + access in tx
	if err := models.DB.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(&project).Exec(ctx); err != nil {
			return fmt.Errorf("insert project error: %v", err)
		}
		// @TODO update members
		return nil
	}); err != nil {
		return fmt.Errorf("project create error: %v", err)
	}

	return c.JSON(http.StatusCreated, project)
}

// ProjectsDelete godoc
//
// @Router			/projects/{id}/	[delete]
// @Summary			Delete
// @Description		Delete project
// @Tags			projects
// @Accept			json
// @Produce			json
// @Param			id				path		integer							true	"Group identifier"
// @Success			204
// @Failure			404
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsDelete(c echo.Context) error {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	project := models.Project{ID: id}
	query := models.DB.NewSelect().Model(&project).
		ColumnExpr("p.*").
		Relation("Group").
		Relation("Members").
		Relation("Members.User").
		Join("LEFT JOIN access AS pa ON pa.project_id = p.id").
		Join("LEFT JOIN access AS ga ON ga.group_id = p.group_id").
		WherePK().
		Where("status = ?", models.ProjectStatusActive).
		Where(
			"(pa.user_id = ? AND pa.role = ?) OR (ga.user_id = ? AND ga.role = ?)",
			authJwtClaims.ID,
			models.AccessRoleAdmin,
			authJwtClaims.ID,
			models.AccessRoleAdmin,
		)
	if err := query.Scan(ctx); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return fmt.Errorf("group retrieve error: %v", err)
	}

	// Soft delete
	project.Status = models.ProjectStatusDelete
	if _, err := models.DB.NewUpdate().Model(&project).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("project delete error: %v", err)
	}

	return c.NoContent(http.StatusNoContent)
}
