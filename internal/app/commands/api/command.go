package api

import (
	"github.com/urfave/cli/v2"
	"gotodo/internal/app/commands/api/server"
)

var Command = &cli.Command{
	Name:   "api",
	Action: Action,
}

func Action(c *cli.Context) error {
	return server.NewServer().Run()
}
