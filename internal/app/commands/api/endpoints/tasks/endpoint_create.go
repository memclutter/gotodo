package tasks

import (
	"gotodo/internal/app/commands/api/schemas"
	"gotodo/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create godoc
//
// @Router			/tasks/		[POST]
// @Summary			Create
func (ep *Endpoint) Create(c echo.Context) error {
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
	if _, err := ep.db.NewInsert().Model(task).Exec(ctx); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, task)
}
