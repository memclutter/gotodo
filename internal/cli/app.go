package cli

import (
	"github.com/memclutter/gotodo/internal/cli/commands"
	"github.com/memclutter/gotodo/internal/cli/hooks"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:    "gotodo",
		Version: "0.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "dsnDb", Value: "postgresql://go_todo:go_todo@localhost:5432/go_todo?sslmode=disable", EnvVars: []string{"DSN_DB"}},
		},
		Before: hooks.Before,
		Commands: cli.Commands{
			&cli.Command{
				Name:   "api",
				Action: commands.Api,
			},
		},
	}
}
