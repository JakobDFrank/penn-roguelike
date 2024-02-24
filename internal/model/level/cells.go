package level

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
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
