package cli

import (
	"github.com/memclutter/gotodo/internal/cli/commands"
	"github.com/memclutter/gotodo/internal/cli/hooks"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:    "gotodo",
		Version: "0.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "dsnDb", Value: "postgresql://go_todo:go_todo@localhost:5432/go_todo?sslmode=disable", EnvVars: []string{"DSN_DB"}},
			&cli.StringFlag{Name: "secret", Value: "random-secret", EnvVars: []string{"SECRET"}},
			&cli.BoolFlag{Name: "debug", Value: false, EnvVars: []string{"DEBUG"}},
			&cli.StringFlag{Name: "logLevel", Value: logrus.InfoLevel.String(), EnvVars: []string{"LOG_LEVEL"}},
		},
		Before: hooks.Before,
		After:  hooks.After,
		Commands: cli.Commands{
			&cli.Command{
				Name:   "api",
				Action: commands.Api,
			},
		},
	}
}
