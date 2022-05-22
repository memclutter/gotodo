package models

import "github.com/uptrace/bun"

type Participant struct {
	bun.BaseModel `bun:"table:participants,alias:p"`

	ID     int64 `bun:"id,pk,autoincrement" json:"id"`
	TaskID int64 `json:"taskId"`
	UserID int64 `json:"userId"`
}
