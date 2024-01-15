package googleauth

import (
	"context"
	"os"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/idtoken"
)

type GoogleUser struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func VerifyToken(token string) (user GoogleUser, err error) {
	payload, err := idtoken.Validate(context.Background(), token, os.Getenv("GOOGLE_CLIENT_ID"))

	if err != nil {
		return
	}

	claims := payload.Claims

	if _, ok := claims["email"]; !ok {
		err = ErrorEmailNotFound
	}

	if _, ok := claims["name"]; !ok {
		err = ErrorNameNotFound
	}

	err = mapstructure.Decode(claims, &user)
	if err != nil {
		err = ErrorDecodingClaims
	}

	return
}
