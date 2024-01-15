package jwtauth

import "errors"

var (
	ErrorCreatingToken         = errors.New("Couldn't create token")
	ErrorTokenExpired          = errors.New("Token has expired")
	ErrorInvalidToken          = errors.New("Invalid token")
	ErrorInvalidTokenSignature = errors.New("Invalid token signature")
	ErrorInvalidClaims         = errors.New("Couldn't parse claims")
)
