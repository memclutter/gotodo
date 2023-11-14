package server

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/labstack/echo/v4"
)

func TestNew(t *testing.T) {
	server := New(echo.New())
	require.NotNil(t, server, "must be return not nil Server interface")
}
