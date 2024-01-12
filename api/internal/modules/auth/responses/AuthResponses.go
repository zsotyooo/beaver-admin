package responses

import (
	UserResponses "api/internal/modules/user/responses"
)

type LoginResponse struct {
	Token string                     `json:"token"`
	User  UserResponses.UserResponse `json:"user"`
}

func CreateLoginResponse(token string, user UserResponses.UserResponse) LoginResponse {
	return LoginResponse{
		Token: token,
		User:  user,
	}
}
