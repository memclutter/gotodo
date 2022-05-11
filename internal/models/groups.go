package models

import (
	"errors"
	"github.com/memclutter/gocore/pkg/coreslices"
	"github.com/uptrace/bun"
)

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

func (g Group) Can(userID int64, roles []string) (bool, error) {
	if g.Members == nil {
		return false, errors.New("check group access error: nil members")
	}
	for _, member := range g.Members {
		if member.UserID == userID && (len(roles) == 0 || coreslices.StringIn(member.Role, roles)) {
			return true, nil
		}
	}
	return false, nil
}
