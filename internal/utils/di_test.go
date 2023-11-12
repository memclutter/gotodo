package utils

import (
	"testing"

	"go.uber.org/dig"

	"github.com/urfave/cli/v2"
)

func TestInvokeWithEmptyDependencies(t *testing.T) {
	action := Invoke([]interface{}{}, func(c *cli.Context, dc *dig.Container) {
		if c == nil {
			t.Fatal("c *cli.Context cannot be nil")
		}
		if dc == nil {
			t.Fatalf("dc *dig.Container cannot be nil")
		}
	})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	err := action(c)
	if err != nil {
		t.Fatalf("cannot be return nil: %v", err)
	}
}

func TestInvokeWithDependencies(t *testing.T) {
	action := Invoke([]interface{}{
		func() int { return 10 },
	}, func(num int) {
		if num != 10 {
			t.Fatalf("num not equal 10")
		}
	})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	err := action(c)
	if err != nil {
		t.Fatalf("cannot be return nil: %v", err)
	}
}

func TestInvokeFailProvide(t *testing.T) {
	action := Invoke([]interface{}{
		func(num1 int) int { return 10 },
		func(num2 int) int { return 11 },
	}, func(num int) {})
	c := cli.NewContext(cli.NewApp(), nil, nil)
	err := action(c)
	if err == nil {
		t.Fatalf("can be return error, but nil return")
	}
}
