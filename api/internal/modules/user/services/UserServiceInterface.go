package services

import (
	UserRequest "api/internal/modules/user/requests"
	UserResponse "api/internal/modules/user/responses"
)

type UserServiceInterface interface {
	GetUsers(limit int) (UserResponse.UsersResponse, error)
	FindUser(id uint) (UserResponse.UserResponse, error)
	FindUserByEmail(email string) (UserResponse.UserResponse, error)
	CreateUser(payload UserRequest.UserCreatePayload) (UserResponse.UserResponse, error)
	UpdateUser(id uint, payload UserRequest.UserUpdatePayload) (UserResponse.UserResponse, error)
	DeleteUser(id uint) error
	UserExists(email string) (bool, error)
}
