package model

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

//--------------------------------------------------------------------------------
// Direction
//--------------------------------------------------------------------------------

// Direction represents different moves the CellPlayer can make on the map
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

// NewDirection creates a new instance of Direction
func NewDirection(dir int) (Direction, error) {
	d := Direction(dir)

	if d.IsValid() {
		return d, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("direction: %d", dir)}
}

// IsValid verifies if the Direction instance is a valid Direction member.
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
