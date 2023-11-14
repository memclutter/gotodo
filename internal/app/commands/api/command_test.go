package api

import (
	"gotodo/internal/app/commands/api/runner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModule(t *testing.T) {
	runnerMock := &runner.Mock{}
	runnerMock.On("Run").Return(nil)
	require.NoError(t, Action(runnerMock), "must be run not error")
	runnerMock.AssertExpectations(t)
}
