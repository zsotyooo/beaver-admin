package requests

// LoginPayload represents the payload for logging in via google oauth token
// @name LoginPayload
type LoginPayload struct {
	Token string `json:"token" binding:"required" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
