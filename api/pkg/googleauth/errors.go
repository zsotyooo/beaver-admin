package googleauth

import "errors"

var (
	ErrorValidatingToken = errors.New("Google oauth: Error validating token")
	ErrorEmailNotFound   = errors.New("Google oauth: Email not found")
	ErrorNameNotFound    = errors.New("Google oauth: Name not found")
	ErrorDecodingClaims  = errors.New("Google oauth: Error decoding claims")
)
