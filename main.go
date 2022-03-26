package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/urfave/cli/v2"
	"net/http"
	"os"
	"strconv"
)

var (
	db *bun.DB
)

func main() {
	app := cli.NewApp()
	app.Name = "Go TODO"
	app.Action = action
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "dsnDb", Value: "postgresql://go_todo:go_todo@localhost:5432/go_todo?sslmode=disable", EnvVars: []string{"DSN_DB"}},
	}
	app.Before = before
	if err := app.Run(os.Args); err != nil {
		logrus.Fatalf("error app run: %v", err)
	}
}

func before(c *cli.Context) error {
	db = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.String("dsnDb")))), pgdialect.New())
	return nil
}

func action(c *cli.Context) error {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/tasks/", tasksGet)
	e.POST("/tasks/", tasksCreate)
	e.GET("/tasks/:id/", tasksRetrieve)
	e.PUT("/tasks/:id/", tasksUpdate)
	e.DELETE("/tasks/:id/", tasksDelete)
	return e.Start(":8000")
}

func tasksGet(c echo.Context) (err error) {
	ctx := c.Request().Context()
	totalCount := 0
	tasks := make([]Task, 0)
	if totalCount, err = db.NewSelect().Model(&tasks).ScanAndCount(ctx); err != nil {
		return fmt.Errorf("tasks get error: %v", err)
	}
	return c.JSON(http.StatusOK, struct {
		TotalCount int    `json:"total_count"`
		Items      []Task `json:"items"`
	}{
		TotalCount: totalCount,
		Items:      tasks,
	})
}

func tasksCreate(c echo.Context) (err error) {
	ctx := c.Request().Context()
	task := Task{}
	if err = c.Bind(&task); err != nil {
		return fmt.Errorf("tasks create error bind: %v", err)
	}
	if _, err = db.NewInsert().Model(&task).Exec(ctx); err != nil {
		return fmt.Errorf("tasks create error: %v", err)
	}
	return c.JSON(http.StatusCreated, task)
}

func tasksRetrieve(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := Task{ID: id}
	if err := db.NewSelect().Model(&task).WherePK("id").Scan(ctx); err != nil {
		return fmt.Errorf("tasks retrieve error: %v", err)
	}
	return c.JSON(http.StatusOK, task)
}

func tasksUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := Task{ID: id}
	if err := db.NewSelect().Model(&task).WherePK().Scan(ctx); err != nil {
		return fmt.Errorf("tasks update error: %v", err)
	}
	if err = c.Bind(&task); err != nil {
		return fmt.Errorf("tasks update error bind: %v", err)
	}
	if _, err := db.NewUpdate().Model(&task).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("tasks update error: %v", err)
	}
	return c.JSON(http.StatusOK, task)
}

func tasksDelete(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusNotFound, "task not found")
	}
	task := Task{ID: id}
	if err := db.NewSelect().Model(&task).WherePK("id").Scan(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	} else if _, err := db.NewDelete().Model(&task).WherePK().Exec(ctx); err != nil {
		return fmt.Errorf("tasks delete error: %v", err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
