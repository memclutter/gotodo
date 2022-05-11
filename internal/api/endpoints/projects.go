package endpoints

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
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
		Join("LEFT JOIN access AS pa ON pa.project_id = p.id").
		Join("LEFT JOIN access AS ga ON ga.group_id = p.group_id").
		Where("status = ?", models.ProjectStatusActive).
		Where("pa.user_id = ? OR ga.user_id = ?", authJwtClaims.ID, authJwtClaims.ID).
		OrderExpr("id")
	if totalCount, err = query.ScanAndCount(ctx); err != nil {
		return fmt.Errorf("groups get error: %v", err)
	}

	return c.JSON(http.StatusOK, schemas.ProjectsListResponse{
		TotalCount: totalCount,
		Items:      projects,
	})
}

func ProjectsCreate(c echo.Context) (err error) {
	return nil
}

func ProjectsRetrieve(c echo.Context) error {
	return nil
}

func ProjectsUpdate(c echo.Context) error {
	return nil
}

func ProjectsDelete(c echo.Context) error {
	return nil
}
