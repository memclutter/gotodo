package api

import (
	"testing"
)

type runnerMock struct {
}

func (m runnerMock) Run() error {
	return nil
}

func TestModule(t *testing.T) {
	err := Action(&runnerMock{})
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
