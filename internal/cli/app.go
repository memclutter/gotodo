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
			&cli.StringFlag{Name: "dsnDb", Value: "postgresql://gotodo:gotodo@localhost:5432/gotodo?sslmode=disable", EnvVars: []string{"DSN_DB"}},
			&cli.StringFlag{Name: "dsnMail", EnvVars: []string{"DSN_MAIL"}},
			&cli.StringFlag{Name: "defaultFromMail", EnvVars: []string{"DEFAULT_FROM_MAIL"}},
			&cli.StringFlag{Name: "urlBase", Value: "http://localhost:3000/", EnvVars: []string{"URL_BASE"}},
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
