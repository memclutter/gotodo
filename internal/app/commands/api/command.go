package api

import (
	"gotodo/internal/app/commands/api/endpoints/tasks"
	"gotodo/internal/app/commands/api/runner"
	"gotodo/internal/app/commands/api/server"
	"gotodo/internal/utils"

	"github.com/labstack/echo/v4"

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
		runner.New,
		server.New,
		echo.New,
		tasks.NewEndpoint,
	}, Action),
}

func Action(r runner.Runner) error {
	return r.Run()
}
