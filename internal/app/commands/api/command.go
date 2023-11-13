package api

import (
	"gotodo/internal/app/commands/api/server"
	"gotodo/internal/utils"

	"github.com/labstack/echo/v4"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name: "api",
	Action: utils.Invoke([]interface{}{
		echo.New,
		server.New,
	}, Action),
}

func Action(s server.Server) error {
	return s.Run()
}
