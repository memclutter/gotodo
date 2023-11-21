package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:tsk"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	IsDone      bool      `bun:"is_done" json:"isDone"`
	DateCreated time.Time `bun:"date_created,nullzero,notnull,default:current_timestamp" json:"dateCreated"`
}
