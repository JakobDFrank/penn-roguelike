package level

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

type Cell int

const (
	Open Cell = iota
	Wall
	Pit
	Arrow
	Player
)

func NewCell(cell int) (Cell, error) {
	c := Cell(cell)

	if c.IsValid() {
		return c, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("cell: %d", cell)}
}

func (c *Cell) IsValid() bool {
	switch *c {
	case Open, Wall, Pit, Arrow, Player:
		return true
	default:
		return false
	}
}
