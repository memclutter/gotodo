package api

import (
	_ "gotodo/internal/app/commands/api/docs"
	"gotodo/internal/app/commands/api/endpoints/tasks"
	"gotodo/internal/utils"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name: "api",
	Flags: utils.FlagsWithEnvs([]cli.Flag{
		&cli.StringFlag{
			Name:  "address",
			Value: ":9000",
			Usage: "The address at which the server will listen",
		},
	}),
	Action: utils.Invoke([]interface{}{
		echo.New,
		tasks.NewEndpoint,
		utils.NewSqlDriverConnector,
		utils.NewSqlDB,
		utils.NewBunDB,
		utils.NewBunIDB,
		utils.NewSchemaDialect,
	}, Action),
}

func Action(c *cli.Context, e *echo.Echo, tasksEndpoint *tasks.Endpoint) error {

	e.Debug = c.Bool("debug")
	e.HideBanner = !e.Debug

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))
	e.GET("/tasks", tasksEndpoint.List)
	e.POST("/tasks", tasksEndpoint.Create)

	return e.Start(c.String("address"))
}
