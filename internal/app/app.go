package app

import (
	"github.com/urfave/cli/v2"
	"gotodo/internal/app/commands/api"
)

var App = &cli.App{
	Name:  "gotodo",
	Usage: "Golang app",
	Commands: cli.Commands{
		api.Command,
	},
}
