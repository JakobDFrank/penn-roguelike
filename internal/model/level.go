package model

import (
	"database/sql/driver"
	"encoding/json"
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
