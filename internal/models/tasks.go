package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks,alias:t"`

	ID             int64     `bun:"id,pk,autoincrement" json:"id"`
	Title          string    `json:"title" validate:"required"`
	Description    string    `json:"description"`
	DateCreated    time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"dateCreated"`
	Status         string    `bun:",type:task_status" json:"status" validate:"required,task_status"`
	CustomStatusID int64     `json:"customStatusId"`
	ProjectID      int64     `json:"projectId"`

	Project      Project       `bun:"rel:belongs-to,join:project_id=id" json:"project"`
	User         User          `bun:"rel:belongs-to,join:user_id=id" json:"user"`
	CustomStatus ProjectStatus `bun:"rel:belongs-to,join:custom_status_id=id" json:"customStatus"`
}

//goland:noinspection GoUnusedConst
const (
	TaskStatusActive  = "active"
	TaskStatusArchive = "archive"
	TaskStatusDelete  = "delete"
)
