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

type IDTokenValidator interface {
	Validate(ctx context.Context, rawToken string, audience string) (*idtoken.Payload, error)
}

type IDTokenWrapper struct{}

func (w IDTokenWrapper) Validate(ctx context.Context, rawToken string, audience string) (*idtoken.Payload, error) {
	return idtoken.Validate(ctx, rawToken, audience)
}

var TokenValidator IDTokenValidator = IDTokenWrapper{}

func VerifyToken(token string) (user GoogleUser, err error) {
	payload, err := TokenValidator.Validate(context.Background(), token, os.Getenv("GOOGLE_CLIENT_ID"))

	if err != nil {
		err = Error(ErrorValidatingToken)
		return
	}

	claims := payload.Claims

	if _, ok := claims["email"]; !ok {
		err = Error(ErrorEmailNotFoundInToken)
		return
	}

	if _, ok := claims["name"]; !ok {
		err = Error(ErrorNameNotFoundInToken)
		return
	}

	err = mapstructure.Decode(claims, &user)
	if err != nil {
		err = Error(ErrorDecodingClaims)
	}

	return
}
