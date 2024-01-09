package repositories

import (
	TodoModel "api/app/modules/todo/models"
	"api/packages/database"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func New() *TodoRepository {
	return &TodoRepository{
		DB: database.Connection(),
	}
}

func (todoRepository *TodoRepository) List(limit int) ([]TodoModel.Todo, error) {
	var todos []TodoModel.Todo

	if err := todoRepository.DB.Limit(limit).Order("created_at DESC").Find(&todos).Error; err != nil {
		return todos, err
	}

	return todos, nil
}

func (TodoRepository *TodoRepository) Find(id uint) (TodoModel.Todo, error) {
	var todo TodoModel.Todo

	TodoRepository.DB.First(&todo, id)

	return todo, nil
}

func (todoRepository *TodoRepository) Create(todo TodoModel.Todo) (TodoModel.Todo, error) {
	var newTodo TodoModel.Todo

	todoRepository.DB.Create(&todo).Scan(&newTodo)

	return newTodo, nil
}

func (todoRepository *TodoRepository) Update(todo, newTodo TodoModel.Todo) (TodoModel.Todo, error) {
	todoRepository.DB.Model(&todo).Updates(&newTodo)

	return todo, nil
}

func (todoRepository *TodoRepository) Delete(id uint) error {
	todoRepository.DB.Delete(&TodoModel.Todo{}, id)
	return nil
}
