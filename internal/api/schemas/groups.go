package schemas

import "github.com/memclutter/gotodo/internal/models"

type GroupsListResponse struct {
	TotalCount int            `json:"totalCount"`
	Items      []models.Group `json:"items"`
}

type AccessMember struct {
	UserID int64  `json:"userId" example:"1"`
	Role   string `json:"role" enums:"admin,member"`
}

type GroupsCreateRequest struct {
	Name    string         `json:"name"`
	Members []AccessMember `json:"members"`
}
