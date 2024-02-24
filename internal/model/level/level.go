package level

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"gorm.io/gorm"
)

const (
	MaxLevelSize = 100
)

type Level struct {
	gorm.Model
	Cells  Cells `gorm:"type:jsonb"`
	XSpawn int
	YSpawn int
}

func NewLevel(cells Cells) (*Level, error) {
	pos, err := validateMap(cells)

	if err != nil {
		return nil, err
	}

	lvl := &Level{
		Cells:  cells,
		XSpawn: pos.X,
		YSpawn: pos.Y,
	}

	return lvl, nil
}

func validateMap(cells Cells) (*cellPos, error) {

	// validate map size, don't want to iterate over potentially massive array
	if err := validateMapSize(cells); err != nil {
		return nil, err
	}

	// validate rectangular map
	if err := validateMapRectangular(cells); err != nil {
		return nil, err
	}

	// validate cells after ensuring map is rectangular, ensure only one player position
	pos, err := validateCells(cells)

	if err != nil {
		return nil, err
	}

	return pos, nil
}

func validateMapSize(cells Cells) error {
	rowCount := len(cells)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	if rowCount > 100 {
		return apperr.ErrMapTooLarge
	}

	expectedColCount := len(cells[0])

	if expectedColCount > 100 {
		return apperr.ErrMapTooLarge
	}

	return nil
}

func validateMapRectangular(cells Cells) error {

	rowCount := len(cells)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	expectedColCount := len(cells[0])

	for _, row := range cells[1:] {
		colCount := len(row)

		if colCount != expectedColCount {
			return apperr.ErrMapNotRectangular
		}
	}

	return nil
}

type cellPos struct {
	X int
	Y int
}

func validateCells(cells Cells) (*cellPos, error) {

	playerCount := 0

	pos := cellPos{}

	for i, row := range cells {
		for j, cell := range row {
			if !cell.IsValid() {
				return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("cell value: %d | row: %d | col: %d", cell, i, j)}
			}

			if cell == StartPosition {
				playerCount += 1

				pos.X = i
				pos.Y = j
				if playerCount > 1 {
					return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("more than one player in map | row: %d | col: %d", i, j)}
				}
			}
		}
	}

	if playerCount != 1 {
		return nil, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("no player in map")}
	}

	return &pos, nil
}
