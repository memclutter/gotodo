package models

import "github.com/uptrace/bun"

type Group struct {
	bun.BaseModel `bun:"table:groups,alias:g"`

	ID     int64  `bun:"id,pk,autoincrement" json:"id"`
	Name   string `json:"name" validate:"required"`
	Status string `bun:",type:group_status" json:"status"`

	Projects []*Project `bun:"rel:has-many,join:id=group_id" json:"projects"`
	Members  []*Access  `bun:"rel:has-many,join:id=group_id" json:"members"`
}

//goland:noinspection GoUnusedConst
const (
	GroupStatusActive  = "active"
	GroupStatusArchive = "archive"
	GroupStatusDeleted = "delete"
)
