package api

import "github.com/urfave/cli/v2"

var Command = &cli.Command{
	Name:   "api",
	Action: Action,
}

func Action(c *cli.Context) error {
	return nil
}
