package model

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
)

type Cell int

const (
	CellOpen Cell = iota
	CellWall
	CellPit
	CellArrow
	CellPlayer
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
	case CellOpen, CellWall, CellPit, CellArrow, CellPlayer:
		return true
	default:
		return false
	}
}
