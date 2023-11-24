package tasks

import (
	"gotodo/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Delete godoc
//
//	@Router		/tasks/{id}		[DELETE]
//	@Summary	Delete
//	@Tags		Tasks
//	@Produce	json
//	@Param		id	path		integer	true	"Task ID"
//	@Failure	400	{object}	schemas.Error
//	@Failure	500	{object}	schemas.Error
//	@Success	204
func (ep *Endpoint) Delete(c echo.Context) (err error) {
	ctx := c.Request().Context()
	task := &models.Task{}
	if task.ID, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound}
	}
	query := ep.db.NewDelete().
		Model(task).
		WherePK()

	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
