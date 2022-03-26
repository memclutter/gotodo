package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/memclutter/gotodo/internal/api/endpoints"
)

func RunServer() error {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/tasks/", endpoints.TasksList)
	e.POST("/tasks/", endpoints.TasksCreate)
	e.GET("/tasks/:id/", endpoints.TasksRetrieve)
	e.PUT("/tasks/:id/", endpoints.TasksUpdate)
	e.DELETE("/tasks/:id/", endpoints.TasksDelete)
	return e.Start(":8000")
}
