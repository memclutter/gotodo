package api

import (
	"gotodo/internal/app/commands/api/server"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestModule(t *testing.T) {
	err := Action(server.New(echo.New()))
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
