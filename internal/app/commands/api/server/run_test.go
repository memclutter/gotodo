package server

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestServer_Run(t *testing.T) {
	server := New(echo.New())
	err := server.Run()
	if err != nil {
		t.Fatalf("must be run without error, but error returned: %v", err)
	}
}
