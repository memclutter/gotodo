package models

import "github.com/uptrace/bun"

type Access struct {
	bun.BaseModel `bun:"table:access,alias:a"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id"`
	UserID    int64  `json:"userId"`
	GroupID   *int64 `json:"groupId"`
	ProjectID *int64 `json:"projectId"`
	Role      string `json:"role"`
}

//goland:noinspection GoUnusedConst
const (
	AccessRoleMember = "member"
	AccessRoleAdmin  = "admin"
)
