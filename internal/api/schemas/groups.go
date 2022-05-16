package schemas

import "github.com/memclutter/gotodo/internal/models"

type GroupsListResponse struct {
	TotalCount int            `json:"totalCount"`
	Items      []models.Group `json:"items"`
}

type GroupsRequest struct {
	Name    string         `json:"name" validate:"required"`
	Members []AccessMember `json:"members"`
}
