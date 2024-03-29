// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Mutation struct {
}

type Query struct {
}

type Direction string

const (
	DirectionLeft  Direction = "LEFT"
	DirectionRight Direction = "RIGHT"
	DirectionUp    Direction = "UP"
	DirectionDown  Direction = "DOWN"
)

var AllDirection = []Direction{
	DirectionLeft,
	DirectionRight,
	DirectionUp,
	DirectionDown,
}

func (e Direction) IsValid() bool {
	switch e {
	case DirectionLeft, DirectionRight, DirectionUp, DirectionDown:
		return true
	}
	return false
}

func (e Direction) String() string {
	return string(e)
}

func (e *Direction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Direction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Direction", str)
	}
	return nil
}

func (e Direction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
