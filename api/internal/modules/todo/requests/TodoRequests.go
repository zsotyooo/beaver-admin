package todo

type TodoCreatePayload struct {
	Title string `json:"title" binding:"required"`
	Done  bool   `json:"done"`
}

type TodoUpdatePayload struct {
	Title *string `json:"title,omitempty"`
	Done  *bool   `json:"done,omitempty"`
}
