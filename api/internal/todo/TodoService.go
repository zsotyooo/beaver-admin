package todo

import (
	"api/pkg/converters"
	"errors"
)

type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{
		todoRepository: *NewTodoRepository(),
	}
}

func (todoService *TodoService) GetTodos(filter TodoListFilter, limit int) (todos []Todo, total int64, err error) {
	todosChan := make(chan []Todo)
	totalChan := make(chan int64)
	errChan := make(chan error, 2)

	go func() {
		var err error
		todos, err = todoService.todoRepository.FindAll(filter, limit)
		todosChan <- todos
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		var err error
		total, err = todoService.todoRepository.Total(filter)
		totalChan <- total
		if err != nil {
			errChan <- err
		}
	}()

	todos = <-todosChan
	total = <-totalChan
	close(errChan)

	for e := range errChan {
		if e != nil {
			err = e
			break
		}
	}

	return
}

func (todoService *TodoService) FindTodo(id uint) (todo Todo, err error) {
	todo, err = todoService.todoRepository.FindById(id)

	if err != nil {
		return
	}

	if todo.ID == 0 {
		err = errors.New("Todo not found!")
	}

	return
}

func (todoService *TodoService) CreateTodo(data TodoFullData) (newTodo Todo, err error) {
	var todo Todo

	todo.UserID = data.UserID
	todo.Title = data.Title
	todo.Done = data.Done

	newTodo, err = todoService.todoRepository.Create(todo)

	if err != nil {
		return
	}

	if newTodo.ID == 0 {
		err = errors.New("Error in creating the todo!")
	}

	return
}

func (todoService *TodoService) UpdateTodo(id uint, data TodoUpdatableData) (updatedTodo Todo, err error) {
	todo, err := todoService.FindTodo(id)

	if err != nil {
		return
	}

	fields, err := converters.StructToMap(data)

	if err != nil {
		return
	}

	updatedTodo, err = todoService.todoRepository.Update(todo, fields)

	return
}

func (todoService *TodoService) DeleteTodo(id uint) (err error) {
	_, err = todoService.FindTodo(id)

	if err != nil {
		return
	}

	err = todoService.todoRepository.Delete(id)

	return
}
