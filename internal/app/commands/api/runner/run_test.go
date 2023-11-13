package runner

import (
	"testing"
)

func TestRunnerStruct_Run(t *testing.T) {
	runner := New(&serverMock{})
	err := runner.Run()
	if err != nil {
		t.Fatalf("must be run without error, but error returned: %v", err)
	}
}
