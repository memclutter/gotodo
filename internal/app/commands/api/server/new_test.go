package server

import (
	"flag"
	"gotodo/internal/app/commands/api/endpoints/tasks"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/require"

	"github.com/labstack/echo/v4"
)

func TestNew(t *testing.T) {
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.Bool("debug", true, "")
	app := cli.NewApp()
	app.Flags = []cli.Flag{&cli.BoolFlag{Name: "debug"}}
	c := cli.NewContext(app, flagSet, nil)
	server := New(c, echo.New(), tasks.NewEndpoint())
	require.IsType(t, &echo.Echo{}, server, "must be type of *echo.Echo")
	require.Equal(t, true, server.(*echo.Echo).Debug, "must be debug is true")
	require.Equal(t, false, server.(*echo.Echo).HideBanner, "must be hide banner is false")
}
