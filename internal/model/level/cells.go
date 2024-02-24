package level

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"strings"
)

type Cells [][]Cell

func (c Cells) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *Cells) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return apperr.ErrInvalidCast
	}
	return json.Unmarshal(b, &c)
}

func (c *Cells) String() string {

	var sb strings.Builder
	for _, row := range *c {
		for _, element := range row {
			sb.WriteString(fmt.Sprintf("%4d", element))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}
