package endpoints

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/memclutter/gotodo/internal/models"
	"net/http"
	"strconv"
)

func TasksList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	totalCount := 0
	tasks := make([]models.Task, 0)
	if totalCount, err = models.DB.NewSelect().Model(&tasks).ScanAndCount(ctx); err != nil {
		return fmt.Errorf("tasks get error: %v", err)
	}
	return c.JSON(http.StatusOK, struct {
		TotalCount int           `json:"total_count"`
		Items      []models.Task `json:"items"`
	}{
		TotalCount: totalCount,
		Items:      tasks,
	})
}

func TasksCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	task := models.Task{}
	if err = c.Bind(&task); err != nil {
		return fmt.Errorf("tasks create error bind: %v", err)
	}
	if _, err = models.DB.NewInsert().Model(&task).Exec(ctx); err != nil {
		return fmt.Errorf("tasks create error: %v", err)
	}
	return c.JSON(http.StatusCreated, task)
}

func TasksRetrieve(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	if err := models.DB.NewSelect().Model(&task).WherePK("id").Scan(ctx); err != nil {
		return fmt.Errorf("tasks retrieve error: %v", err)
	}
	return c.JSON(http.StatusOK, task)
}

func TasksUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	if err := models.DB.NewSelect().Model(&task).WherePK().Scan(ctx); err != nil {
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
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := models.Task{ID: id}
	if err := models.DB.NewSelect().Model(&task).WherePK("id").Scan(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	} else if _, err := models.DB.NewDelete().Model(&task).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
