package controllers

import userControllers "api/internal/user/controller"

// LoginPayload represents the payload for logging in via google oauth token
// @name LoginPayload
type LoginPayload struct {
	Token string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// LoginResponse represents the response for logging in
// @name LoginResponse
type LoginResponse struct {
	Token string                       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  userControllers.UserResponse `json:"user"`
}
