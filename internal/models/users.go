package models

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID       int64  `bun:"id,pk,autoincrement" json:"id"`
	Email    string `json:"email"`
	PwHash   string `json:"-"`
	Password string `bun:"-" json:"password,omitempty"`
	Status   int    `json:"status"`

	Sessions []Session `bun:"rel:has-many,join:id=user_id" json:"sessions"`
}

const (
	UserStatusPending = iota
	UserStatusActive
	UserStatusBlock
)
