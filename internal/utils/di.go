package utils

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

func Invoke(provides []interface{}, invoke interface{}) cli.ActionFunc {
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
}
