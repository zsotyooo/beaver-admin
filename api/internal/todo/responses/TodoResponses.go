package responses

import (
	TodoModel "api/internal/todo/models"
	"time"

	"github.com/thoas/go-funk"
)

// TodoResponse represents a todo in the response
// @name TodoResponse
type TodoResponse struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"Todo title"`
	Done      bool   `json:"done" example:"true"`
	CreatedAt string `json:"createdAt" example:"2024-01-09T11:59:57Z"`
}

// TodosResponse represents a list of todos in the response
// @name TodosResponse
type TodosResponse struct {
	Data []TodoResponse `json:"data"`
}

func ConvertModelToResponse(todo TodoModel.Todo) TodoResponse {
	return TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
	}
}

func ConvertModelsToResponse(todos []TodoModel.Todo) TodosResponse {
	return TodosResponse{
		Data: funk.Map(todos, func(todo TodoModel.Todo) TodoResponse {
			return ConvertModelToResponse(todo)
		}).([]TodoResponse),
	}
}
