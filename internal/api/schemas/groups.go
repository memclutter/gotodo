package schemas

import "github.com/memclutter/gotodo/internal/models"

type GroupsListResponse struct {
	TotalCount int            `json:"totalCount"`
	Items      []models.Group `json:"items"`
}
