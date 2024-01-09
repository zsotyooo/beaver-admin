package responses

import (
	TodoModel "api/internal/modules/todo/models"
	"time"

	"github.com/thoas/go-funk"
)

type TodoResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"createdAt"`
}

type TodosResponse struct {
	Data []TodoResponse `json:"data"`
}

func ToTodo(todo TodoModel.Todo) TodoResponse {
	return TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
	}
}

func ToTodos(todos []TodoModel.Todo) TodosResponse {
	return TodosResponse{
		Data: funk.Map(todos, func(todo TodoModel.Todo) TodoResponse {
			return ToTodo(todo)
		}).([]TodoResponse),
	}
}
