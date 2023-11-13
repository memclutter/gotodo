package runner

import (
	"testing"
)

type serverMock struct {
}

func (m serverMock) Start(address string) error {
	return nil
}

func TestNew(t *testing.T) {
	runner := New(&serverMock{})
	if runner == nil {
		t.Fatal("must be return not nil")
	}
}
