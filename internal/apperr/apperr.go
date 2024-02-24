package apperr

import (
	"errors"
	"fmt"
)

// https://github.com/uber-go/guide/blob/master/style.md#errors

type ArgumentNilError struct {
	Message string
}

func (e *ArgumentNilError) Error() string {
	return fmt.Sprintf("argument nil: %s", e.Message)
}

var (
	ErrInvalidCast = errors.New("type_cast_failure")
)
