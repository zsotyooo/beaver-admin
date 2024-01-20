package controllers

import (
	"api/internal/user"
	"time"

	"github.com/thoas/go-funk"
)

func ConvertUserModelToResponse(userItem user.User) UserResponse {
	return UserResponse{
		ID:        userItem.ID,
		Email:     userItem.Email,
		Name:      userItem.Name,
		Role:      string(userItem.Role),
		CreatedAt: userItem.CreatedAt.Format(time.RFC3339),
	}
}

func ConvertUserModelsToResponse(users []user.User, total int64) UsersResponse {
	return UsersResponse{
		Data: funk.Map(users, func(userItem user.User) UserResponse {
			return ConvertUserModelToResponse(userItem)
		}).([]UserResponse),
		Total: total,
	}
}
