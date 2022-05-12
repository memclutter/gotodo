package schemas

import "github.com/memclutter/gotodo/internal/models"

type TasksListRequest struct {
	ProjectID int64 `query:"projectId"`
}

type TasksListResponse struct {
	TotalCount int           `json:"totalCount"`
	Items      []models.Task `json:"items"`
}
