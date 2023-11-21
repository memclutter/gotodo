package schemas

type Error struct {
	Message          string              `json:"message,omitempty"`
	Details          string              `json:"details,omitempty"`
	ValidationErrors map[string][]string `json:"validationErrors,omitempty"`
}
