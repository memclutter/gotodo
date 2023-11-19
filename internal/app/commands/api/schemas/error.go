package schemas

type Error struct {
	Message          string
	ValidationErrors interface{}
}
