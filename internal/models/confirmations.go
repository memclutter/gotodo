package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Confirmation struct {
	bun.BaseModel `bun:"table:confirmations,alias:c"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	UserID      int64     `json:"user_id"`
	Token       string    `json:"token"`
	DateCreated time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"dateCreated"`
	DateExpired time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"dateExpired"`

	User User `bun:"rel:belongs-to,join:user_id=id" json:"user"`
}

func (c Confirmation) IsExpired() bool {
	return c.DateExpired.Before(time.Now().UTC())
}
