package services

import (
	TodoRequest "api/internal/todo/requests"
	TodoResponse "api/internal/todo/responses"
)

type TodoServiceInterface interface {
	GetTodos(limit int) (TodoResponse.TodosResponse, error)
	FindTodo(id uint) (TodoResponse.TodoResponse, error)
	CreateTodo(payload TodoRequest.TodoCreatePayload) (TodoResponse.TodoResponse, error)
	UpdateTodo(id uint, payload TodoRequest.TodoUpdatePayload) (TodoResponse.TodoResponse, error)
	DeleteTodo(id uint) error
}
