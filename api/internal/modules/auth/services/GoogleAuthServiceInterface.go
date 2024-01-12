package services

type GoogleAuthServiceInterface interface {
	VerifyToken(token string) (GoogleUser, error)
}

type GoogleUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
