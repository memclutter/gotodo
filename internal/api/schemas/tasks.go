package schemas

import "github.com/memclutter/gotodo/internal/models"

type TasksListResponse struct {
	TotalCount int           `json:"total_count"`
	Items      []models.Task `json:"items"`
}
