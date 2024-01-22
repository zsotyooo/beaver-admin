package jwtauth

import (
	"errors"
	"fmt"
)

var (
	ErrorCreatingToken              = errors.New("Couldn't create token")
	ErrorTokenExpired               = errors.New("Token has expired")
	ErrorInvalidToken               = errors.New("Invalid token")
	ErrorInvalidTokenSignature      = errors.New("Invalid token signature")
	ErrorInvalidClaims              = errors.New("Couldn't parse claims")
	ErrorAuthTokenNotFoundInRequest = errors.New("No auth token found in request")
)

func Error(err error) error {
	return fmt.Errorf("Error in jwt auth: %w", err)
}
