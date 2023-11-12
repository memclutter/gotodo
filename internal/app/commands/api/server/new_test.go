package server

import "testing"

func TestNewServer(t *testing.T) {
	server := NewServer()
	if server == nil {
		t.Fatal("must be return not nil")
	}
}
