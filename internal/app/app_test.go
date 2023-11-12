package app

import (
	"github.com/urfave/cli/v2"
	"testing"
)

func TestAction(t *testing.T) {
	err := Action(cli.NewContext(App, nil, nil))
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
