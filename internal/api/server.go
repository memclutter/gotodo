//go:generate swag init --parseInternal --parseDependency --parseDepth 1 --instanceName api -g server.go
package api

import (
	"github.com/labstack/echo/v4"
	_ "github.com/memclutter/gotodo/internal/api/docs"
	"github.com/memclutter/gotodo/internal/api/endpoints"
	"github.com/memclutter/gotodo/internal/api/middleware"
	"github.com/memclutter/gotodo/internal/api/plugins"
	"github.com/memclutter/gotodo/internal/config"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// RunServer godoc
//
// @title gotodo
// @version 1.0
// @description gotodo api
//
// @host localhost
// @BasePath /api/
//
// @securityDefinitions.apiKey ApiKeyHeader
// @in header
// @name Authorization
//
// @tag.name auth
// @tag.description Auth endpoints
//
// @tag.name groups
// @tag.description Groups endpoints
//
// @tag.name tasks
// @tag.description Tasks endpoint
//
// @tag.name profile
// @tag.description Profile endpoint
func RunServer() error {
	e := echo.New()
	e.Debug = config.Config.Debug

	corsMiddleware := middleware.NewCors()
	authMiddleware := middleware.NewAuth()

	e.Use(corsMiddleware)

	e.Validator = plugins.NewValidator()
	e.HTTPErrorHandler = plugins.NewErrorHandler(e)

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))

	e.POST("/auth/registration/", endpoints.AuthRegistration)
	e.POST("/auth/confirmation/", endpoints.AuthConfirmation)
	e.POST("/auth/login/", endpoints.AuthLogin)
	e.POST("/auth/refresh/", endpoints.AuthRefresh)

	tasks := e.Group("/tasks", authMiddleware)
	{
		tasks.GET("/", endpoints.TasksList)
		tasks.POST("/", endpoints.TasksCreate)
		tasks.GET("/:id/", endpoints.TasksRetrieve)
		tasks.PUT("/:id/", endpoints.TasksUpdate)
		tasks.DELETE("/:id/", endpoints.TasksDelete)
	}

	groups := e.Group("/groups", authMiddleware)
	{
		groups.GET("/", endpoints.GroupsList)
		groups.POST("/", endpoints.GroupsCreate)
		groups.GET("/:id/", endpoints.GroupsRetrieve)
		groups.PUT("/:id/", endpoints.GroupsUpdate)
		groups.DELETE("/:id/", endpoints.GroupsDelete)
	}

	profile := e.Group("/profile", authMiddleware)
	{
		profile.GET("/", endpoints.ProfileRetrieve)
		profile.PUT("/", endpoints.ProfileUpdate)
		profile.PUT("/password/", endpoints.ProfilePasswordUpdate)
	}

	return e.Start(":8000")
}
