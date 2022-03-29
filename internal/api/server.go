//go:generate swag init --parseInternal --parseDependency --parseDepth 1 --instanceName api -g server.go
package api

import (
	"github.com/labstack/echo/v4"
	_ "github.com/memclutter/gotodo/internal/api/docs"
	"github.com/memclutter/gotodo/internal/api/endpoints"
	"github.com/memclutter/gotodo/internal/api/middleware"
	"github.com/memclutter/gotodo/internal/config"
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
// @BasePath /api/
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
	e.Debug = config.Config.Debug

	corsMiddleware := middleware.NewCors()
	authMiddleware := middleware.NewAuth()

	e.Use(corsMiddleware)

	e.GET("/docs/*", echoSwagger.EchoWrapHandler(echoSwagger.InstanceName("api")))

	e.POST("/auth/login/", endpoints.AuthLogin)
	e.POST("/auth/refresh/", endpoints.AuthRefresh)
	e.POST("/auth/registration/", endpoints.AuthRegistration)
	e.GET("/auth/info/", endpoints.AuthInfo, authMiddleware)
	e.POST("/auth/logout/", endpoints.AuthLogout, authMiddleware)

	tasks := e.Group("/tasks", authMiddleware)
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
