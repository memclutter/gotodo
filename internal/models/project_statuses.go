package models

import "github.com/uptrace/bun"

type ProjectStatus struct {
	bun.BaseModel `bun:"table:project_statuses,alias:ps"`

	ID        int64  `bun:"id,pk,autoincrement" json:"id"`
	ProjectID int64  `json:"projectId"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
	IsFinal   bool   `json:"isFinal"`
}
