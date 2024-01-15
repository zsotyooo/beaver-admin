package requests

// TodoCreatePayload represents the payload for creating a todo
// @name TodoCreatePayload
// @param title body string true "Title of the todo"
// @param done body bool false "Done status of the todo"
type TodoCreatePayload struct {
	Title string `json:"title" binding:"required" example:"New todo title"`
	Done  bool   `json:"done" example:"false"`
}

// TodoUpdatePayload represents the payload for updating a todo
// @name TodoUpdatePayload
// @param title body string false "Title of the todo"
// @param done body bool false "Done status of the todo"
type TodoUpdatePayload struct {
	Title *string `json:"title,omitempty" example:"New title"`
	Done  *bool   `json:"done,omitempty" example:"false"`
}
