package app

import (
	"github.com/urfave/cli/v2"
	"gotodo/internal/app/commands/api"
)

var App = &cli.App{
	Name:   "gotodo",
	Usage:  "Golang app",
	Action: Action,
	Commands: cli.Commands{
		api.Command,
	},
}

func Action(c *cli.Context) error {
	return nil
}
