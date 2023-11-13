package runner

import (
	"gotodo/internal/app/commands/api/server"
)

func New(s server.Server) Runner { return &runnerStruct{server: s} }
