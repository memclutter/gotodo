package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	DateCreated time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"dateCreated"`
	Status      string    `bun:",type:task_status" json:"status" validate:"required,task_status"`

	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}

const (
	TaskStatusTodo       = "todo"
	TaskStatusInProgress = "inProgress"
	TaskStatusClosed     = "closed"
)

var TaskStatuses = []string{
	TaskStatusTodo,
	TaskStatusInProgress,
	TaskStatusClosed,
}
