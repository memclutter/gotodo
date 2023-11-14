package runner

import (
	"gotodo/internal/app/commands/api/server"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunnerStruct_Run(t *testing.T) {
	address := "9001"
	serverMock := &server.Mock{}
	serverMock.On("Start", address).Return(nil)
	runner := &runnerStruct{
		address: address,
		server:  serverMock,
	}

	require.NoError(t, runner.Run(), "must be run without error")
}
