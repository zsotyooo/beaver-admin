package repositories

import (
	UserModel "api/internal/modules/user/models"
)

type UserRepositoryInterface interface {
	List(limit int) ([]UserModel.User, error)
	Find(id uint) (UserModel.User, error)
	FindByEmail(email string) (UserModel.User, error)
	Create(user UserModel.User) (UserModel.User, error)
	Update(user UserModel.User, fields map[string]interface{}) (UserModel.User, error)
	Delete(id uint) error
}
