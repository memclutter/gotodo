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
// @Param			request			body		models.Project		true	"Request body"
// @Success			200				{object}	models.Project
// @Failure			400				{object}	schemas.Error		true	"Validation error"
// @Failure			500				{object} 	schemas.Error		true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authJwtClaims := helpers.GetAuthJwtClaims(c)
	project := models.Project{}
	access := models.Access{UserID: authJwtClaims.ID, Role: models.AccessRoleAdmin}
	if err = c.Bind(&project); err != nil {
		return err
	} else if err = c.Validate(&project); err != nil {
		return err
	}
	project.Status = models.GroupStatusActive

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
		members := make([]*models.Access, 0)
		members = append(members, &access)

		if project.Members != nil {
			for _, member := range project.Members {
				members = append(members, &models.Access{
					ProjectID: project.ID,
					UserID:    member.UserID,
					Role:      member.Role,
				})
			}
		}

		if _, err := tx.NewInsert().Model(&members).Exec(ctx); err != nil {
			return fmt.Errorf("insert members error: %v", err)
		}
		project.Members = members
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
	return nil
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
// @Param			request			body		models.Project					true	"Request body"
// @Success			200				{object}	models.Project
// @Failure			400				{object}	schemas.Error					true	"Validation error"
// @Failure			500				{object}	schemas.Error					true	"Server error"
// @Security		ApiHeaderAuth
func ProjectsUpdate(c echo.Context) error {
	return nil
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
	return nil
}
