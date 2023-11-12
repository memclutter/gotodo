package api

import (
	"fmt"
	"gotodo/internal/app/commands/api/server"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

var Command = &cli.Command{
	Name: "api",
	Action: func(provides []interface{}, invoke interface{}) func(c *cli.Context) error {
		return func(c *cli.Context) error {
			container := dig.New()

			provides = append(
				provides,
				func() *cli.Context { return c },
				func() *dig.Container { return container },
			)

			for _, provide := range provides {
				if err := container.Provide(provide); err != nil {
					return fmt.Errorf("error provide dependency: %v", err)
				}
			}

			return container.Invoke(invoke)
		}
	}(
		[]interface{}{
			server.NewServer,
		},
		Action,
	),
}

func Action(s server.Server) error {
	return s.Run()
}
