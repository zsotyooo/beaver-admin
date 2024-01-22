package auth

import (
	"api/internal/user"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrorBadLoginCredentials = errors.New("Bad login credentials")
	ErrorNotAuthorized       = errors.New("User not authorized")
	ErrorForbiddenForUser    = errors.New("Forbidden for this user")
)

func Error(err error) error {
	return fmt.Errorf("Error in auth: %w", err)
}

func GetHttpStatusCode(err error) int {
	if errors.Is(err, ErrorBadLoginCredentials) {
		return http.StatusBadRequest
	}

	if errors.Is(err, ErrorNotAuthorized) {
		return http.StatusUnauthorized
	}

	if errors.Is(err, ErrorForbiddenForUser) {
		return http.StatusForbidden
	}

	if errors.Is(err, user.ErrorUserNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
