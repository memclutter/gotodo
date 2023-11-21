package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// List godoc
//
//	@Router		/tasks/ 		[GET]
//	@Summary	List
//	@Tags		Tasks
//	@Produce	json
//	@Param		data	query		schemas.TaskListRequest	true	"Request params"
//	@Failure	500		{object}	schemas.Error
//	@Success	200		{array}		models.Task
func (ep *Endpoint) List(c echo.Context) error {
	ctx := c.Request().Context()
	request := &schemas.TaskListRequest{}
	if err := c.Bind(request); err != nil {
		return err
	} else if err := c.Validate(request); err != nil {
		return err
	}

	tasks := make([]models.Task, 0)
	query := ep.db.
		NewSelect().
		Model(&tasks).
		Offset(request.Offset).
		Limit(request.Limit)

	if err := query.Scan(ctx); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}
