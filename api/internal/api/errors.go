package api

import (
	"fmt"
)

type HttpError struct {
	Code    int
	Message string
}

func (e *HttpError) Error() error {
	return fmt.Errorf("Error: %s", e.Message)
}
