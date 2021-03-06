package models

import "github.com/uptrace/bun"

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	ID      int64  `bun:"id,pk,autoincrement" json:"id"`
	GroupID int64  `json:"groupId"`
	Name    string `json:"name" validate:"required"`
	Status  string `bun:",type:project_status" json:"status"`

	Group    Group     `bun:"rel:belongs-to,join:group_id=id" json:"group" validate:"-"`
	Members  []*Access `bun:"rel:has-many,join:id=project_id" json:"members"`
	Statuses []Status  `bun:"rel:has-many,join:id=project_id" json:"statuses" validate:"-"`
}

//goland:noinspection GoUnusedConst
const (
	ProjectStatusActive  = "active"
	ProjectStatusArchive = "archive"
	ProjectStatusDelete  = "delete"
)
