package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"gorm.io/gorm"
	"strings"
)

//--------------------------------------------------------------------------------
// Level
//--------------------------------------------------------------------------------

const (
	MaxLevelSize = 100
)

// Level is an entity in the database that holds information on map data
type Level struct {
	gorm.Model
	Map         GameMap `gorm:"type:jsonb"`
	RowStartIdx int     // RowStartIdx is the player's starting row index. It is used to restore the player on death.
	ColStartIdx int     // ColStartIdx is the player's starting column index. It is used to restore the player on death.
}

func NewLevel(gameMap GameMap) (*Level, error) {
	pos, err := validateMap(gameMap)

	if err != nil {
		return nil, err
	}

	lvl := &Level{
		Map:         gameMap,
		RowStartIdx: pos.RowIdx,
		ColStartIdx: pos.ColIdx,
	}

	return lvl, nil
}

func validateMap(gameMap GameMap) (*cellPos, error) {

	// validate map size, don't want to iterate over potentially massive array
	if err := validateMapSize(gameMap); err != nil {
		return nil, err
	}

	// validate rectangular map
	if err := validateMapRectangular(gameMap); err != nil {
		return nil, err
	}

	// validate gameMap after ensuring map is rectangular, ensure only one player position
	pos, err := validateCells(gameMap)

	if err != nil {
		return nil, err
	}

	return pos, nil
}

func validateMapSize(gameMap GameMap) error {
	rowCount := len(gameMap)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	if rowCount > 100 {
		return apperr.ErrMapTooLarge
	}

	expectedColCount := len(gameMap[0])

	if expectedColCount > 100 {
		return apperr.ErrMapTooLarge
	}

	return nil
}

func validateMapRectangular(gameMap GameMap) error {

	rowCount := len(gameMap)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	expectedColCount := len(gameMap[0])

	for _, row := range gameMap[1:] {
		colCount := len(row)

		if colCount != expectedColCount {
			return apperr.ErrMapNotRectangular
		}
	}

	return nil
}

type cellPos struct {
	RowIdx int
	ColIdx int
}

func validateCells(gameMap GameMap) (*cellPos, error) {

	playerCount := 0

	pos := &cellPos{}

	for rowIdx, row := range gameMap {
		for colIdx, cell := range row {
			if !cell.IsValid() {
				return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("cell value: %d | row: %d | col: %d", cell, rowIdx, colIdx)}
			}

			if cell == CellPlayer {
				playerCount += 1

				pos.RowIdx = rowIdx
				pos.ColIdx = colIdx

				if playerCount > 1 {
					return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("more than one player in map | row: %d | col: %d", rowIdx, colIdx)}
				}
			}
		}
	}

	if playerCount != 1 {
		return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("no player in map")}
	}

	return pos, nil
}

//--------------------------------------------------------------------------------
// GameMap
//--------------------------------------------------------------------------------

// GameMap represents a game map
type GameMap [][]Cell

func (c GameMap) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *GameMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return apperr.ErrInvalidCast
	}
	return json.Unmarshal(b, &c)
}

func (c *GameMap) String() string {

	var sb strings.Builder
	for _, row := range *c {
		for _, element := range row {
			sb.WriteString(fmt.Sprintf("%4d", element))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

//--------------------------------------------------------------------------------
// Cell
//--------------------------------------------------------------------------------

// Cell represents different types of objects found within a game level.
type Cell int

const (
	CellOpen   Cell = iota // An open tile
	CellWall               // An impassable barrier
	CellPit                // A hazard that does one damage
	CellArrow              // A hazard that does two damage
	CellPlayer             // The player's character
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
