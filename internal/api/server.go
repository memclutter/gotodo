//go:generate swag init --parseInternal --parseDependency --parseDepth 1 --instanceName api -g server.go
package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/memclutter/gotodo/internal/api/docs"
	"github.com/memclutter/gotodo/internal/api/endpoints"
	"github.com/memclutter/gotodo/internal/api/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

// RunServer godoc
//
// @title gotodo
// @version 1.0
// @description gotodo api
//
// @host localhost
// @BasePath /
//
// @securityDefinitions.apiKey ApiKeyHeader
// @in Header
// @name Authorization
//
// @tag.name auth
// @tag.description Auth endpoints
//
// @tag.name tasks
// @tag.description Tasks endpoint
func RunServer() error {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))

	authPublic := e.Group("/auth")
	authProtected := e.Group("/auth", middleware.User)
	{
		authPublic.POST("/login/", endpoints.AuthLogin)
		authPublic.POST("/refresh/", endpoints.AuthRefresh)
		authPublic.POST("/registration/", endpoints.AuthRegistration)
		authProtected.GET("/info/", endpoints.AuthInfo, middleware.User)
		authProtected.DELETE("/logout/", endpoints.AuthLogout, middleware.User)
	}

	tasks := e.Group("/tasks", middleware.User)
	{
		tasks.GET("/", endpoints.TasksList)
		tasks.POST("/", endpoints.TasksCreate)
		tasks.GET("/:id/", endpoints.TasksRetrieve)
		tasks.PUT("/:id/", endpoints.TasksUpdate)
		tasks.DELETE("/:id/", endpoints.TasksDelete)
	}

	e.GET("/_routes.json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})

	return e.Start(":8000")
}
