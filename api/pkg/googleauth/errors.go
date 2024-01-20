package googleauth

import (
	"errors"
	"fmt"
)

var (
	ErrorValidatingToken      = errors.New("Error validating token")
	ErrorEmailNotFoundInToken = errors.New("Email not found in claims")
	ErrorNameNotFoundInToken  = errors.New("Name not found in claims")
	ErrorDecodingClaims       = errors.New("Error decoding claims")
)

func Error(err error) error {
	return fmt.Errorf("Error in google auth: %w", err)
}
