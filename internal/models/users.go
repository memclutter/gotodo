package models

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID           int64  `bun:"id,pk,autoincrement" json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Password     string `bun:"-" json:"password,omitempty"`
	Status       string `bun:",type:user_status" json:"status"`
}

//goland:noinspection GoUnusedConst
const (
	UserStatusPending = "pending"
	UserStatusActive  = "active"
	UserStatusDelete  = "delete"
)
