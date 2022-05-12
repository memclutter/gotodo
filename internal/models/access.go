package models

import "github.com/uptrace/bun"

type Access struct {
	bun.BaseModel `bun:"table:access,alias:a"`

	ID        int64  `bun:"id,pk,autoincrement" json:"-"`
	UserID    int64  `json:"userId"`
	GroupID   int64  `bun:",nullzero" json:"-"`
	ProjectID int64  `bun:",nullzero" json:"-"`
	Role      string `json:"role"`

	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}

//goland:noinspection GoUnusedConst
const (
	AccessRoleMember = "member"
	AccessRoleAdmin  = "admin"
)
