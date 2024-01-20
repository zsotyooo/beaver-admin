package controllers

// UserResponse represents a user in the response
// @name UserResponse
type UserResponse struct {
	ID        uint   `json:"id" example:"1"`
	Email     string `json:"email" example:"john.doe@gmail.com"`
	Name      string `json:"name" example:"John Doe"`
	Role      string `json:"role" example:"user" enums:"user,moderator,admin"`
	CreatedAt string `json:"createdAt" example:"2024-01-09T11:59:57Z"`
}

// UsersResponse represents a list of users in the response
// @name UsersResponse
type UsersResponse struct {
	Data  []UserResponse `json:"data"`
	Total int64          `json:"total" example:"100"`
}
