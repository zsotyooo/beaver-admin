package controllers

import (
	"api/internal/todo"
	userControllers "api/internal/user/controller"
	"time"

	"github.com/thoas/go-funk"
)

func ConvertTodoModelToResponse(todoItem todo.Todo) TodoResponse {
	return TodoResponse{
		ID:        todoItem.ID,
		Title:     todoItem.Title,
		Done:      todoItem.Done,
		CreatedAt: todoItem.CreatedAt.Format(time.RFC3339),
		User:      userControllers.ConvertUserModelToResponse(todoItem.User),
	}
}

func ConvertTodoModelsToResponse(todos []todo.Todo, total int64) TodosResponse {
	return TodosResponse{
		Data: funk.Map(todos, func(todoItem todo.Todo) TodoResponse {
			return ConvertTodoModelToResponse(todoItem)
		}).([]TodoResponse),
		Total: total,
	}
}
