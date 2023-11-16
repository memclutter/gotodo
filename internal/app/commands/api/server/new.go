package server

import (
	"gotodo/internal/app/commands/api/endpoints/tasks"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"

	_ "gotodo/internal/app/commands/api/docs"
)

func New(c *cli.Context, e *echo.Echo, tasksEndpoint *tasks.Endpoint) Server {
	// @TODO some echo setup
	e.Debug = c.Bool("debug")
	e.HideBanner = !e.Debug

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))
	e.GET("/tasks", tasksEndpoint.List)
	return e
}
