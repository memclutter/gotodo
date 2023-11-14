package runner

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func TestNew(t *testing.T) {
	address := "9001"
	app := cli.NewApp()
	app.Flags = []cli.Flag{&cli.StringFlag{Name: "address"}}
	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.String("address", address, "")
	c := cli.NewContext(app, flagSet, nil)
	runner := New(nil, c)
	require.NotNil(t, runner, "must be runner not nil")
	require.Equal(t, address, runner.(*runnerStruct).address, "must be address set")
}
