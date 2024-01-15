package responses

import (
	UserResponses "api/internal/user/responses"
)

// LoginResponse represents the response for logging in
// @name LoginResponse
type LoginResponse struct {
	Token string                     `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponses.UserResponse `json:"user"`
}

func CreateLoginResponse(token string, user UserResponses.UserResponse) LoginResponse {
	return LoginResponse{
		Token: token,
		User:  user,
	}
}
