package endpoints

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/api/helpers"
	"github.com/memclutter/gotodo/internal/api/schemas"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
	"strconv"
)

// TasksList godoc
//
// @Router 			/tasks/			[get]
// @Summary			List
// @Description		Get tasks list
// @Tags			tasks
// @Accept			json
// @Produce			json
// @Success			200	{object}				schemas.TasksListResponse
// @Security		ApiHeaderAuth
func TasksList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authUser := helpers.GetAuthUser(c)
	totalCount := 0
	tasks := make([]models.Task, 0)
	query := models.DB.NewSelect().Model(&tasks).
		Where("user_id = ?", authUser.ID)
	if totalCount, err = query.ScanAndCount(ctx); err != nil {
		return fmt.Errorf("tasks get error: %v", err)
	}
	return c.JSON(http.StatusOK, schemas.TasksListResponse{
		TotalCount: totalCount,
		Items:      tasks,
	})
}

// TasksCreate godoc
//
// @Router 			/tasks/			[post]
// @Summary			Create
// @Description		Create new tasks
// @Tags			tasks
// @Accept			json
// @Produce			json
// @Param			request			body		models.Task		true	"Request body"
// @Success			201	{object}				models.Task
// @Security		ApiHeaderAuth
func TasksCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	authUser := helpers.GetAuthUser(c)
	task := models.Task{}
	if err = c.Bind(&task); err != nil {
		return fmt.Errorf("tasks create error bind: %v", err)
	}
	task.UserID = authUser.ID
	if _, err = models.DB.NewInsert().Model(&task).Exec(ctx); err != nil {
		return fmt.Errorf("tasks create error: %v", err)
	}
	return c.JSON(http.StatusCreated, task)
}

func TasksRetrieve(c echo.Context) error {
	ctx := c.Request().Context()
	authUser := helpers.GetAuthUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	query := models.DB.NewSelect().Model(&task).
		Where("id = ?", task.ID).
		Where("user_id = ?", authUser.ID)
	if err := query.Scan(ctx); err != nil {
		return fmt.Errorf("tasks retrieve error: %v", err)
	}
	return c.JSON(http.StatusOK, task)
}

func TasksUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	authUser := helpers.GetAuthUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	query := models.DB.NewSelect().Model(&task).
		Where("id = ?", task.ID).
		Where("user_id = ?", authUser.ID)
	if err := query.Scan(ctx); err != nil {
		return fmt.Errorf("tasks update error: %v", err)
	}
	if err = c.Bind(&task); err != nil {
		return fmt.Errorf("tasks update error bind: %v", err)
	}
	if _, err := models.DB.NewUpdate().Model(&task).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("tasks update error: %v", err)
	}
	return c.JSON(http.StatusOK, task)
}

func TasksDelete(c echo.Context) error {
	ctx := c.Request().Context()
	authUser := helpers.GetAuthUser(c)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	query := models.DB.NewSelect().Model(&task).
		Where("id = ?", task.ID).
		Where("user_id = ?", authUser.ID)
	if err := query.Scan(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	} else if _, err := models.DB.NewDelete().Model(&task).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
