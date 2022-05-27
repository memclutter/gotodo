package endpoints

import (
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
)

// StatusesCreate godoc
//
// @Router 				/statuses/ 		[post]
// @Summary				Create
// @Description			Create status
// @Tags				statuses
// @Accept 				json
// @Produce				json
// @Param				request			body			schemas.StatusesRequest			true	"Request body"
// @Success				200				{object}		models.Status
// @Failure				400				{object}		schemas.Error					true	"Validation error"
// @Failure				500				{object}		schemas.Error					true	"Server error"
// @Security			ApiHeaderAuth
func StatusesCreate(c echo.Context) (err error) {
	req := schemas.StatusesRequest{}
	if err = c.Bind(&req); err != nil {
		return err
	} else if err = c.Validate(&req); err != nil {
		return err
	}
	project := models.Status{
		ProjectID: req.ProjectID,
		Name:      req.Name,
		SortOrder: req.SortOrder,
		IsFinal:   req.IsFinal,
	}
	return c.JSON(http.StatusCreated, project)
}

func StatusesUpdate(c echo.Context) error {
	return nil
}

func StatusesDelete(c echo.Context) error {
	return nil
}
