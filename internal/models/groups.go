package models

import "github.com/uptrace/bun"

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`

	ID     int64  `bun:"id,pk,autoincrement" json:"id"`
	Name   string `json:"name"`
	Status string `bun:",type:group_status" json:"status"`
}

//goland:noinspection GoUnusedConst
const (
	GroupStatusActive  = "active"
	GroupStatusArchive = "archive"
	GroupStatusDeleted = "delete"
)
