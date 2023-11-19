package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.Model `bun:"table:tasks,alias:tsk"`

	ID          int64 `bun:"id,pk,autoincrement"`
	Title       string
	Body        string
	IsDone      bool
	DateCreated time.Time
}
