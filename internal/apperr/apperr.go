package apperr

import (
	"errors"
	"fmt"
)

// https://github.com/uber-go/guide/blob/master/style.md#errors

type NilArgumentError struct {
	Message string
}

func (e *NilArgumentError) Error() string {
	return fmt.Sprintf("argument nil: %s", e.Message)
}

type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid argument: %s", e.Message)
}

type InvalidCellTypeError struct {
	Message string
}

func (e *InvalidCellTypeError) Error() string {
	return fmt.Sprintf("invalid cell specified: %s", e.Message)
}

var (
	ErrInvalidCast       = errors.New("type cast failure")
	ErrMapNotRectangular = errors.New("map not rectangular")
	ErrMapTooLarge       = errors.New("map width and map height must not exceed 100 units")
	ErrEmptyMap          = errors.New("the map must have at least one cell")
)
