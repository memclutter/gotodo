package runner

import "gotodo/internal/app/commands/api/server"

type runnerStruct struct {
	address string
	server  server.Server
}
