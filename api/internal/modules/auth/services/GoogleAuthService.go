package services

import (
	"context"
	"errors"
	"os"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

type GoogleAuthService struct {
}

func New() *GoogleAuthService {
	return &GoogleAuthService{}
}

func (googleAuthService *GoogleAuthService) VerifyToken(token string) (user GoogleUser, err error) {
	payload, err := idtoken.Validate(context.Background(), token, os.Getenv("GOOGLE_CLIENT_ID"))

	if err != nil {
		return
	}

	claims := payload.Claims

	if _, ok := claims["email"]; !ok {
		err = errors.New("Email not found!")
	}

	if _, ok := claims["name"]; !ok {
		err = errors.New("Name not found!")
	}

	err = mapstructure.Decode(claims, &user)

	return
}
