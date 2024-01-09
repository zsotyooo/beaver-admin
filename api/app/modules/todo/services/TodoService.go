package services

import (
	TodoModel "api/app/modules/todo/models"
	TodoRepository "api/app/modules/todo/repositories"
	TodoRequest "api/app/modules/todo/requests"
	TodoResponse "api/app/modules/todo/responses"
	"errors"
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

	return TodoResponse.ToTodos(todos), nil

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

	return TodoResponse.ToTodo(todo), nil
}

func (todoService *TodoService) CreateTodo(request TodoRequest.TodoCreateRequest) (TodoResponse.TodoResponse, error) {
	var todo TodoModel.Todo
	var response TodoResponse.TodoResponse

	todo.Title = request.Title
	todo.Done = request.Done

	newTodo, err := todoService.todoRepository.Create(todo)

	if err != nil {
		return response, err
	}

	if newTodo.ID == 0 {
		return response, errors.New("Error in creating the todo!")
	}

	return TodoResponse.ToTodo(newTodo), nil
}

func (todoService *TodoService) UpdateTodo(id uint, request TodoRequest.TodoUpdateRequest) (TodoResponse.TodoResponse, error) {
	var response TodoResponse.TodoResponse
	todo, err := todoService.todoRepository.Find(id)

	if err != nil {
		return response, err
	}

	if todo.ID == 0 {
		return response, errors.New("Todo not found!")
	}

	newTodo := TodoRequest.UpdateRequestToTodo(request)
	updatedTodo, err := todoService.todoRepository.Update(todo, newTodo)

	if err != nil {
		return response, err
	}

	return TodoResponse.ToTodo(updatedTodo), nil
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
