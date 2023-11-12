package app

import (
	"gotodo/internal/app/commands/api"

	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "gotodo",
	Usage: "Golang app",
	Commands: cli.Commands{
		api.Command,
	},
}
