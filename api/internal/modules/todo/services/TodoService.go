package services

import (
	TodoModel "api/internal/modules/todo/models"
	TodoRepository "api/internal/modules/todo/repositories"
	TodoRequest "api/internal/modules/todo/requests"
	TodoResponse "api/internal/modules/todo/responses"
	"errors"

	"api/packages/converters"
)

type TodoService struct {
	todoRepository TodoRepository.TodoRepositoryInterface
}

func New() *TodoService {
	return &TodoService{
		todoRepository: TodoRepository.New(),
	}
}

func (todoService *TodoService) GetTodos(limit int) (TodoResponse.TodosResponse, error) {
	todos, err := todoService.todoRepository.List(limit)

	if err != nil {
		return TodoResponse.TodosResponse{}, err
	}

	return TodoResponse.ConvertModelsToResponse(todos), nil
}

func (todoService *TodoService) FindTodo(id uint) (TodoResponse.TodoResponse, error) {
	var response TodoResponse.TodoResponse

	todo, err := todoService.todoRepository.Find(id)

	if err != nil {
		return response, err
	}

	if todo.ID == 0 {
		return response, errors.New("Todo not found!")
	}

	return TodoResponse.ConvertModelToResponse(todo), nil
}

func (todoService *TodoService) CreateTodo(payload TodoRequest.TodoCreatePayload) (TodoResponse.TodoResponse, error) {
	var todo TodoModel.Todo
	var response TodoResponse.TodoResponse

	todo.Title = payload.Title
	todo.Done = payload.Done

	newTodo, err := todoService.todoRepository.Create(todo)

	if err != nil {
		return response, err
	}

	if newTodo.ID == 0 {
		return response, errors.New("Error in creating the todo!")
	}

	return TodoResponse.ConvertModelToResponse(newTodo), nil
}

func (todoService *TodoService) UpdateTodo(id uint, payload TodoRequest.TodoUpdatePayload) (TodoResponse.TodoResponse, error) {
	var response TodoResponse.TodoResponse
	todo, err := todoService.todoRepository.Find(id)

	if err != nil {
		return response, err
	}

	if todo.ID == 0 {
		return response, errors.New("Todo not found!")
	}

	fields, err := converters.StructToMap(payload)

	if err != nil {
		return response, err
	}

	updatedTodo, err := todoService.todoRepository.Update(todo, fields)

	if err != nil {
		return response, err
	}

	return TodoResponse.ConvertModelToResponse(updatedTodo), nil
}

func (todoService *TodoService) DeleteTodo(id uint) error {
	todo, err := todoService.todoRepository.Find(id)

	if err != nil {
		return err
	}

	if todo.ID == 0 {
		return errors.New("Todo not found!")
	}

	todoService.todoRepository.Delete(id)

	return nil
}
