package models

import "github.com/uptrace/bun"

type TaskParticipant struct {
	bun.BaseModel `bun:"table:task_participants,alias:tp"`

	ID     int64 `bun:"id,pk,autoincrement" json:"id"`
	TaskID int64 `json:"taskId"`
	UserID int64 `json:"userId"`
}
