package api

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestModule(t *testing.T) {
	err := Action(cli.NewContext(cli.NewApp(), nil, nil))
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
