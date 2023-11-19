package app

import (
	"gotodo/internal/app/commands/api"
	"gotodo/internal/utils"

	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "gotodo",
	Usage: "Golang app",
	Flags: utils.FlagsWithEnvs([]cli.Flag{
		&cli.StringFlag{Name: "dbDsn", Value: "postgresql://user:secret@localhost:5432/db?sslmode=disable"},
	}),
	Commands: cli.Commands{
		api.Command,
	},
}
