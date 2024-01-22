package user

import (
	database "api/pkg/database/postgresql"
	"sync"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var (
	instance *UserRepository
	once     sync.Once
)

func NewUserRepository() *UserRepository {
	once.Do(func() {
		instance = &UserRepository{
			db: database.DBConnection(),
		}
	})
	return instance
}

func (repo *UserRepository) FindAll(limit int) (users []User, err error) {
	err = repo.db.Limit(limit).Order("created_at DESC").Find(&users).Error
	return
}

func (repo *UserRepository) Total() (total int64, err error) {
	err = repo.db.Model(&User{}).Count(&total).Error
	return
}

func (repo *UserRepository) FindById(id uint) (user User, err error) {
	err = repo.db.First(&user, id).Error
	return
}

func (repo *UserRepository) FindByEmail(email string) (user User, err error) {
	err = repo.db.Where("email = ?", email).First(&user).Error
	return
}

func (repo *UserRepository) Create(user User) (newUser User, err error) {
	err = repo.db.Create(&user).Scan(&newUser).Error
	return
}

func (repo *UserRepository) Update(user User, fields map[string]interface{}) (User, error) {
	err := repo.db.Model(&user).Updates(fields).Error
	return user, err
}

func (repo *UserRepository) Delete(id uint) error {
	err := repo.db.Delete(&User{}, id).Error
	return err
}
