package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	Cells Cells `gorm:"type:jsonb"`
}

type Cell int

const (
	Open Cell = iota
	Wall
	Pit
	Arrow
	StartPosition
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
	case Open, Wall, Pit, Arrow, StartPosition:
		return true
	default:
		return false
	}
}

type Cells [][]Cell

func (m Cells) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Cells) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return apperr.ErrInvalidCast
	}
	return json.Unmarshal(b, &m)
}
