package models

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	PwHash    string `json:"-"`
	Password  string `bun:"-" json:"password,omitempty"`
	Status    string `bun:",type:user_status" json:"status"`
}

const (
	UserStatusPending = "pending"
	UserStatusActive  = "active"
	UserStatusBlock   = "block"
)
