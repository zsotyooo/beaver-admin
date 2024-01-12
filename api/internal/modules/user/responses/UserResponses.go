package responses

import (
	UserModel "api/internal/modules/user/models"
	"time"

	"github.com/thoas/go-funk"
)

type UserResponse struct {
	ID        uint               `json:"id"`
	Email     string             `json:"email"`
	Name      string             `json:"name"`
	Role      UserModel.UserRole `json:"role"`
	CreatedAt string             `json:"createdAt"`
}

type UsersResponse struct {
	Data []UserResponse `json:"data"`
}

func ConvertModelToResponse(user UserModel.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}
}

func ConvertModelsToResponse(users []UserModel.User) UsersResponse {
	return UsersResponse{
		Data: funk.Map(users, func(user UserModel.User) UserResponse {
			return ConvertModelToResponse(user)
		}).([]UserResponse),
	}
}
