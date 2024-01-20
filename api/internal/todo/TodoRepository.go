package todo

import (
	"api/pkg/database"
	"sync"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

var (
	instance *TodoRepository
	once     sync.Once
)

func NewTodoRepository() *TodoRepository {
	once.Do(func() {
		instance = &TodoRepository{
			db: database.Connection(),
		}
	})
	return instance
}

func (repo *TodoRepository) FindAll(filter TodoListFilter, limit int) (todos []Todo, err error) {
	err = repo.db.Joins("User").Limit(limit).Where(filter).Order("created_at DESC").Find(&todos).Error
	return
}

func (repo *TodoRepository) Total(filter TodoListFilter) (total int64, err error) {
	err = repo.db.Model(&Todo{}).Where(filter).Count(&total).Error
	return
}

func (repo *TodoRepository) FindById(id uint) (todo Todo, err error) {
	err = repo.db.Joins("User").First(&todo, id).Error
	return
}

func (repo *TodoRepository) Create(todo Todo) (newTodo Todo, err error) {
	err = repo.db.Joins("User").Create(&todo).Scan(&newTodo).Error
	return
}

func (repo *TodoRepository) Update(todo Todo, fields map[string]interface{}) (Todo, error) {
	err := repo.db.Joins("User").Model(&todo).Updates(fields).Error
	return todo, err
}

func (repo *TodoRepository) Delete(id uint) error {
	err := repo.db.Delete(&Todo{}, id).Error
	return err
}
