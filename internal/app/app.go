package app

import (
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:   "gotodo",
	Usage:  "Golang app",
	Action: Action,
}

func Action(c *cli.Context) error {
	return nil
}
