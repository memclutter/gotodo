package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments,alias:cm"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	TaskID      int64     `json:"taskId"`
	UserID      int64     `json:"userId"`
	DateCreated time.Time `json:"dateCreated"`
	Title       string    `json:"title"`
	Message     string    `json:"message"`

	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}
