package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"go.uber.org/dig"

	"github.com/urfave/cli/v2"
)

func TestInvokeWithEmptyDependencies(t *testing.T) {
	action := Invoke([]interface{}{}, func(c *cli.Context, dc *dig.Container) {
		require.NotNil(t, c, "cannot be nil")
		require.NotNil(t, dc, "cannot be nil")
	})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	require.NoError(t, action(c), "cannot be error")
}

func TestInvokeWithDependencies(t *testing.T) {
	action := Invoke([]interface{}{
		func() int { return 10 },
	}, func(num int) {
		require.Equal(t, 10, num, "can be equal 10")
	})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	require.NoError(t, action(c), "must be run without error")
}

func TestInvokeFailProvide(t *testing.T) {
	action := Invoke([]interface{}{
		func(num1 int) int { return 10 },
		func(num2 int) int { return 11 },
	}, func(num int) {})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	require.Error(t, action(c), "must be run with error")
}
