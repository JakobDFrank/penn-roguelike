package model

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

//--------------------------------------------------------------------------------
// Cell
//--------------------------------------------------------------------------------

// Cell represents different types of objects found within a game level.
type Cell int32

const (
	CellOpen   Cell = iota // An open tile
	CellWall               // An impassable barrier
	CellPit                // A hazard that does one damage
	CellArrow              // A hazard that does two damage
	CellPlayer             // The player's character
)

// NewCell creates an implementation of Cell.
func NewCell(cell int32) (Cell, error) {
	c := Cell(cell)

	if c.IsValid() {
		return c, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("cell: %d", cell)}
}

// IsValid verifies if the Cell instance is a valid Cell member.
func (c *Cell) IsValid() bool {
	switch *c {
	case CellOpen, CellWall, CellPit, CellArrow, CellPlayer:
		return true
	default:
		return false
	}
}
