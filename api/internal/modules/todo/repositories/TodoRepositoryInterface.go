package repositories

import (
	TodoModel "api/internal/modules/todo/models"
)

type TodoRepositoryInterface interface {
	List(limit int) ([]TodoModel.Todo, error)
	Find(id uint) (TodoModel.Todo, error)
	Create(todo TodoModel.Todo) (TodoModel.Todo, error)
	Update(todo, newTodo TodoModel.Todo) (TodoModel.Todo, error)
	Delete(id uint) error
}
