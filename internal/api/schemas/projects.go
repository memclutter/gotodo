package schemas

import "github.com/memclutter/gotodo/internal/models"

type ProjectsListRequest struct {
	GroupID int64 `query:"groupId"`
}

type ProjectsListResponse struct {
	TotalCount int              `json:"totalCount"`
	Items      []models.Project `json:"items"`
}

type ProjectsRequest struct {
	Name    string         `json:"name" validate:"required"`
	GroupID int64          `json:"groupId" validate:"required"`
	Members []AccessMember `json:"members"`
}
