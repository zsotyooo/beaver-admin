package repositories

import (
	UserModel "api/internal/modules/user/models"
	"api/packages/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (userRepository *UserRepository) List(limit int) (users []UserModel.User, err error) {
	err = userRepository.DB.Limit(limit).Order("created_at DESC").Find(&users).Error
	return
}

func (userRepository *UserRepository) Find(id uint) (user UserModel.User, err error) {
	err = userRepository.DB.First(&user, id).Error
	return
}

func (userRepository *UserRepository) FindByEmail(email string) (user UserModel.User, err error) {
	err = userRepository.DB.Where("email = ?", email).First(&user).Error
	return
}

func (userRepository *UserRepository) Create(user UserModel.User) (newUser UserModel.User, err error) {
	err = userRepository.DB.Create(&user).Scan(&newUser).Error
	return
}

func (userRepository *UserRepository) Update(user UserModel.User, fields map[string]interface{}) (UserModel.User, error) {
	err := userRepository.DB.Model(&user).Updates(fields).Error
	return user, err
}

func (userRepository *UserRepository) Delete(id uint) error {
	err := userRepository.DB.Delete(&UserModel.User{}, id).Error
	return err
}
