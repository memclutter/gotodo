package api

import (
	"gotodo/internal/app/commands/api/runner"
	"gotodo/internal/app/commands/api/server"
	"gotodo/internal/utils"

	"github.com/labstack/echo/v4"

	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name: "api",
	Action: utils.Invoke([]interface{}{
		runner.New,
		server.New,
		echo.New,
	}, Action),
}

func Action(r runner.Runner) error {
	return r.Run()
}
