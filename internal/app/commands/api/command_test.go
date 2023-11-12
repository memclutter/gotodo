package api

import (
	"github.com/urfave/cli/v2"
	"gotodo/internal/app"
	"testing"
)

func TestModule(t *testing.T) {
	err := Action(cli.NewContext(app.App, nil, nil))
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
