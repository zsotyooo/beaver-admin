package repositories

import (
	TodoModel "api/internal/todo/models"
	"api/pkg/database"

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

func (todoRepository *TodoRepository) List(limit int) (todos []TodoModel.Todo, err error) {
	err = todoRepository.DB.Limit(limit).Order("created_at DESC").Find(&todos).Error
	return
}

func (TodoRepository *TodoRepository) Find(id uint) (todo TodoModel.Todo, err error) {
	err = TodoRepository.DB.First(&todo, id).Error
	return
}

func (todoRepository *TodoRepository) Create(todo TodoModel.Todo) (newTodo TodoModel.Todo, err error) {
	err = todoRepository.DB.Create(&todo).Scan(&newTodo).Error
	return
}

func (todoRepository *TodoRepository) Update(todo TodoModel.Todo, fields map[string]interface{}) (TodoModel.Todo, error) {
	err := todoRepository.DB.Model(&todo).Updates(fields).Error
	return todo, err
}

func (todoRepository *TodoRepository) Delete(id uint) error {
	err := todoRepository.DB.Delete(&TodoModel.Todo{}, id).Error
	return err
}
