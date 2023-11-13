package server

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestNew(t *testing.T) {
	server := New(echo.New())
	if server == nil {
		t.Fatal("must be return not nil Server interface")
	}
}
