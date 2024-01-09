package services

import (
	TodoRequest "api/app/modules/todo/requests"
	TodoResponse "api/app/modules/todo/responses"
)

type TodoServiceInterface interface {
	GetTodos(limit int) (TodoResponse.TodosResponse, error)
	FindTodo(id uint) (TodoResponse.TodoResponse, error)
	CreateTodo(request TodoRequest.TodoCreateRequest) (TodoResponse.TodoResponse, error)
	UpdateTodo(id uint, request TodoRequest.TodoUpdateRequest) (TodoResponse.TodoResponse, error)
	DeleteTodo(id uint) error
}
