package user

import (
	UserModel "api/internal/modules/user/models"
)

type UserCreatePayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Role  string `json:"role" binding:"required"`
}

type UserUpdatePayload struct {
	Name  string             `json:"name,omitempty"`
	Email string             `json:"email,omitempty"`
	Role  UserModel.UserRole `json:"role,omitempty"`
}
