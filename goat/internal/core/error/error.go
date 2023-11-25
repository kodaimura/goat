package error

import (
	"fmt"
)


type InvalidArgumentError struct {
	Message string
}

func NewInvalidArgumentError(msg string) *InvalidArgumentError {
	return &InvalidArgumentError{Message: msg}
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("InvalidArgumentError: %s", e.Message)
}