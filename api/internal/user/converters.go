package user

import (
	"time"

	"github.com/thoas/go-funk"
)

func ConvertModelToResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func ConvertModelsToResponse(users []User, total int64) UsersResponse {
	return UsersResponse{
		Data: funk.Map(users, func(user User) UserResponse {
			return ConvertModelToResponse(user)
		}).([]UserResponse),
		Total: total,
	}
}
