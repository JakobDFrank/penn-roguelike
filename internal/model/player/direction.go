package player

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func NewDirection(dir int) (Direction, error) {
	d := Direction(dir)

	if d.IsValid() {
		return d, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("direction: %d", dir)}
}

func (d *Direction) IsValid() bool {
	switch *d {
	case Left, Right, Up, Down:
		return true
	default:
		return false
	}
}

func (d *Direction) String() string {
	switch *d {
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	default:
		return "unimplemented"
	}
}
