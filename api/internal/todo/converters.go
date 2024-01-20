package todo

import (
	"api/internal/user"
	"time"

	"github.com/thoas/go-funk"
)

func ConvertModelToResponse(todo Todo) TodoResponse {
	return TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt.Format(time.RFC3339),
		User:      user.ConvertModelToResponse(todo.User),
	}
}

func ConvertModelsToResponse(todos []Todo, total int64) TodosResponse {
	return TodosResponse{
		Data: funk.Map(todos, func(todo Todo) TodoResponse {
			return ConvertModelToResponse(todo)
		}).([]TodoResponse),
		Total: total,
	}
}
