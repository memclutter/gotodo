package schemas

type Error struct {
	Message          string
	Details          string
	ValidationErrors map[string][]string
}
