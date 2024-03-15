// Package apperr defines custom error types
// see: https://github.com/uber-go/guide/blob/master/style.md#errors
package apperr

import (
	"errors"
	"fmt"
)

// NilArgumentError indicates that a function argument was unexpectedly nil.
type NilArgumentError struct {
	Message string
}

func (e *NilArgumentError) Error() string {
	return fmt.Sprintf("argument nil: %s", e.Message)
}

// InvalidArgumentError indicates that a function argument
// received a value that is outside the acceptable range
type InvalidArgumentError struct {
	Message string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid argument: %s", e.Message)
}

// InvalidCellTypeError indicates that an unexpected cell type was encountered.
type InvalidCellTypeError struct {
	Message string
}

func (e *InvalidCellTypeError) Error() string {
	return fmt.Sprintf("invalid cell specified: %s", e.Message)
}

// UnimplementedError indicates that the functionality still needs to be implemented.
type UnimplementedError struct {
	Message string
}

func (e *UnimplementedError) Error() string {
	return fmt.Sprintf("unimplemented: %s", e.Message)
}

// InvalidOperationError indicates a method call is invalid.
type InvalidOperationError struct {
	Message string
}

func (e *InvalidOperationError) Error() string {
	return fmt.Sprintf("invalid operatoin: %s", e.Message)
}

// ErrInvalidCast indicates a type cast failure.
var ErrInvalidCast = errors.New("type cast failure")

// ErrMapNotRectangular indicates that a map is not rectangular.
var ErrMapNotRectangular = errors.New("map not rectangular")

// ErrMapTooLarge indicates that a map is too large.
var ErrMapTooLarge = errors.New("map width and map height must not exceed 100 units")

// ErrEmptyMap indicates that a map is empty.
var ErrEmptyMap = errors.New("the map must have at least one cell")
