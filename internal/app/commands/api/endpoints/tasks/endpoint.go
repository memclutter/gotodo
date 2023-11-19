package tasks

import "github.com/uptrace/bun"

type Endpoint struct {
	db bun.IDB
}

func NewEndpoint(bunIDB bun.IDB) *Endpoint {
	return &Endpoint{db: bunIDB}
}
