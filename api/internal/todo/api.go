package todo

import (
	"api/internal/user"
	"errors"
)

// TodoListFilterQuery is used to filter the todo list based on certain criteria.
// swagger:model TodoListFilter
type TodoListFilter struct {
	UserID *uint `form:"user_id" binding:"omitempty" example:"1"`
	Done   *bool `form:"done" binding:"omitempty" example:"false"`
}

func (f *TodoListFilter) Validate(authUser user.User) error {
	if f.UserID == nil {
		return nil
	}

	if authUser.IsSuperUser() {
		return nil
	}

	if *f.UserID == authUser.ID {
		return nil
	}

	return errors.New("You are not authorized to view this users todos")
}

// TodoCreatePayload represents the payload for creating a todo
// @name TodoCreatePayload
// @param title body string true "Title of the todo"
// @param done body bool false "Done status of the todo"
type TodoCreatePayload struct {
	UserID uint   `json:"user_id" binding:"required" example:"1"`
	Title  string `json:"title" binding:"required" example:"New todo title"`
	Done   bool   `json:"done" example:"false"`
}

func (p *TodoCreatePayload) Validate(authUser user.User) error {
	if authUser.IsSuperUser() {
		return nil
	}

	if p.UserID == authUser.ID {
		return nil
	}

	return errors.New("You are not authorized to create this users todos")
}

// TodoUpdatePayload represents the payload for updating a todo
// @name TodoUpdatePayload
// @param title body string false "Title of the todo"
// @param done body bool false "Done status of the todo"
type TodoUpdatePayload struct {
	Title *string `json:"title,omitempty" example:"New title"`
	Done  *bool   `json:"done,omitempty" example:"false"`
}

// TodoResponse represents a todo in the response
// @name TodoResponse
type TodoResponse struct {
	ID        uint              `json:"id" example:"1"`
	Title     string            `json:"title" example:"Todo title"`
	Done      bool              `json:"done" example:"true"`
	CreatedAt string            `json:"createdAt" example:"2024-01-09T11:59:57Z"`
	User      user.UserResponse `json:"user"`
}

// TodosResponse represents a list of todos in the response
// @name TodosResponse
type TodosResponse struct {
	Data  []TodoResponse `json:"data"`
	Total int64          `json:"total" example:"100"`
}
