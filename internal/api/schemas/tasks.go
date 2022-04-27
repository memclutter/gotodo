package schemas

import "github.com/memclutter/gotodo/internal/models"

type TasksListResponse struct {
	TotalCount int           `json:"totalCount"`
	Items      []models.Task `json:"items"`
}
