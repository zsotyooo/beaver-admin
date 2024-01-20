package user

import (
	"errors"
	"fmt"
)

var (
	ErrorUserNotFound   = errors.New("User not found")
	ErrorUserNotCreated = errors.New("Error in creating the user")
)

func Error(err error) error {
	return fmt.Errorf("Error in user: %w", err)
}
