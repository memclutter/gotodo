package commands

import (
	"github.com/memclutter/gotodo/internal/api"
	"github.com/urfave/cli/v2"
)

func Api(c *cli.Context) error {
	return api.RunServer()
}
