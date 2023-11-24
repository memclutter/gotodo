package api

import (
	_ "gotodo/internal/app/commands/api/docs"
	"gotodo/internal/app/commands/api/endpoints/tasks"
	"gotodo/internal/app/commands/api/extensions"
	"gotodo/internal/app/commands/api/validations"
	"gotodo/internal/utils"

	"github.com/go-playground/validator/v10"

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
		// third party dependencies
		echo.New,
		validator.New,

		// internal dependencies
		utils.NewSqlDriverConnector,
		utils.NewSqlDB,
		utils.NewBunDB,
		utils.NewBunIDB,
		utils.NewSchemaDialect,
		utils.NewValidator,
		utils.NewValidate,

		// command dependencies
		tasks.NewEndpoint,
		validations.NewEmail,
	}, Action),
}

func Action(
	c *cli.Context,
	e *echo.Echo,
	validator *utils.Validator,
	tasksEndpoint *tasks.Endpoint,
	_ *validations.Email,
) error {

	e.Debug = c.Bool("debug")
	e.HideBanner = !e.Debug
	e.HTTPErrorHandler = extensions.ErrorHandler
	e.Validator = validator

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))
	e.GET("/tasks", tasksEndpoint.List)
	e.POST("/tasks", tasksEndpoint.Create)
	e.PUT("/tasks/:id", tasksEndpoint.Update)
	e.DELETE("/tasks/:id", tasksEndpoint.Delete)

	return e.Start(c.String("address"))
}
