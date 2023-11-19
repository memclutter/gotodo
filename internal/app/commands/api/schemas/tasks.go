package schemas

type TaskListRequest struct {
	Offset int
	Limit  int
}

type TaskForm struct {
	Title string
	Body  string
}
