package runner

import (
	"gotodo/internal/app/commands/api/server"

	"github.com/urfave/cli/v2"
)

func New(s server.Server, c *cli.Context) Runner {
	return &runnerStruct{
		address: c.String("address"),
		server:  s,
	}
}
