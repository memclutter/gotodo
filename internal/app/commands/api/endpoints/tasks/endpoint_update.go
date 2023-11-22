package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Update godoc
//
//	@Router		/tasks/{id}		[PUT]
//	@Summary	Update
//	@Tags		Tasks
//	@Accept		json
//	@Produce	json
//	@Param		id		path		integer				true	"Task ID"
//	@Param		data	body		schemas.TaskForm	true	"Request body"
//	@Failure	400		{object}	schemas.Error
//	@Failure	500		{object}	schemas.Error
//	@Success	201		{object}	models.Task
func (ep *Endpoint) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()
	request := &schemas.TaskForm{}
	if err := c.Bind(request); err != nil {
		return err
	} else if err := c.Validate(request); err != nil {
		return err
	}

	task := &models.Task{
		Title: request.Title,
		Body:  request.Body,
	}
	if task.ID, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}

	query := ep.db.NewUpdate().
		Model(task).
		ExcludeColumn("date_created", "is_done").
		Returning("date_created, is_done").
		WherePK()
	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}
