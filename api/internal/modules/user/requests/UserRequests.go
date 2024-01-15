package requests

import (
	UserModel "api/internal/modules/user/models"
)

type UserCreatePayload struct {
	Name  string `json:"name" binding:"required" example:"John Doe"`
	Email string `json:"email" binding:"required" example:"john.doe@gmail.com"`
	Role  string `json:"role" binding:"required" example:"user" enums:"user,moderator,admin"`
}

type UserUpdatePayload struct {
	Name  string             `json:"name,omitempty" example:"John Doe"`
	Email string             `json:"email,omitempty" example:"john.doe@gmail.com"`
	Role  UserModel.UserRole `json:"role,omitempty" example:"user" enums:"user,moderator,admin"`
}
