package server

import (
	"github.com/labstack/echo/v4"
	"testing"
)

func TestNewServer(t *testing.T) {
	server := New(echo.New())
	if server == nil {
		t.Fatal("must be return not nil")
	}
}
