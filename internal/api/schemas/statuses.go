package schemas

type StatusesRequest struct {
	ProjectID int64  `json:"projectId" validate:"required"`
	Name      string `json:"name" validate:"required"`
	IsFinal   bool   `json:"isFinal"`
	SortOrder int    `json:"sort_order"`
}
