package server

import "testing"

func TestServer_Run(t *testing.T) {
	server := NewServer()
	err := server.Run()
	if err != nil {
		t.Fatalf("must be run without error, but error returned: %v", err)
	}
}
