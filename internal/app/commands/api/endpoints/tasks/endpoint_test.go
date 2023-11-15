package tasks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEndpoint(t *testing.T) {
	endpoint := NewEndpoint()
	require.NotNil(t, endpoint, "must be not nil")
}
