package tasks

import (
	"gotodo/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEndpoint_List(t *testing.T) {
	db := &mocks.BunDB{}
	c := &mocks.EchoContext{}
	endpoint := NewEndpoint(db)
	require.NoError(t, endpoint.List(c))
}
