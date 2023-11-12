package api

import (
	"gotodo/internal/app/commands/api/server"
	"testing"
)

func TestModule(t *testing.T) {
	err := Action(server.NewServer())
	if err != nil {
		t.Fatalf("must be no error, but return error: %v", err)
	}
}
