package tasks

import (
	"gotodo/internal/mocks"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEndpoint(t *testing.T) {
	db := &mocks.BunDB{}
	endpoint := NewEndpoint(db)
	require.NotNil(t, endpoint, "must be not nil")
}
