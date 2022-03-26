package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateCreated time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"dateCreated"`
	Status      int       `json:"status"`
}
