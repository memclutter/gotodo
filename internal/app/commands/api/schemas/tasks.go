package schemas

type TaskListRequest struct {
	Offset int
	Limit  int
}

type TaskForm struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body"`
}
