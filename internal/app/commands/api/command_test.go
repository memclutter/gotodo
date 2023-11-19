package api

import (
	"context"
	"flag"
	"gotodo/internal/app/commands/api/endpoints/tasks"
	"gotodo/internal/mocks"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/labstack/echo/v4"
	"github.com/urfave/cli/v2"
)

func TestAction(t *testing.T) {

	// Arrange
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "address"},
		&cli.BoolFlag{Name: "debug"},
	}
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.String("address", ":9100", "Address usage")
	flagSet.Bool("debug", false, "Debug usage")
	c := cli.NewContext(app, flagSet, nil)
	e := echo.New()
	db := &mocks.BunDB{}
	tasksEndpoint := tasks.NewEndpoint(db)

	// Act
	errCh := make(chan error)
	go func() {
		errCh <- Action(c, e, tasksEndpoint)
	}()

	// Assert
	var err error
	time.Sleep(100 * time.Millisecond) // @TODO need more beautiful solution)
	conn, err := net.Dial("tcp", ":9100")
	require.NoError(t, err, "must be server listen on :9100")
	require.NoError(t, conn.Close(), "must be conn close without error")
	require.NoError(t, e.Shutdown(context.Background()), "must be shutdown without error")
	err = <-errCh
	require.ErrorIs(t, http.ErrServerClosed, err, "must be run without error")
}
