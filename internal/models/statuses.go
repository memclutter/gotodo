package models

import "github.com/uptrace/bun"

type Status struct {
	bun.BaseModel `bun:"table:statuses,alias:s"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id"`
	ProjectID int64  `json:"projectId"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
	IsFinal   bool   `json:"isFinal"`
}
