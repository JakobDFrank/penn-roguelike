package service

import (
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"testing"
)

func TestValidateMapSize(t *testing.T) {
	emptyMap := make([][]model.Cell, 0)

	if err := validateMapSize(emptyMap); err == nil {
		t.Error("validateMapSize on an empty map did not return an error")
	}

	largeMap := make([][]model.Cell, MaxLevelSize+1)

	if err := validateMapSize(largeMap); err == nil {
		t.Error("validateMapSize on a map too large did not return an error")
	}
}

func TestValidateMapRectangular(t *testing.T) {
	jaggedMap := make([][]model.Cell, 10)

	for i := range jaggedMap {
		jaggedMap[i] = make([]model.Cell, i+1)
	}

	if err := validateMapRectangular(jaggedMap); err == nil {
		t.Error("validateMapRectangular on a jagged map did not return an error")
	}

	squareMap := make([][]model.Cell, 10)

	for i := range squareMap {
		squareMap[i] = make([]model.Cell, len(squareMap))
	}

	if err := validateMapRectangular(squareMap); err != nil {
		t.Errorf("validateMapRectangular on a square map returned an error: %v", err)
	}

	rectMap := make([][]model.Cell, 10)

	for i := range rectMap {
		rectMap[i] = make([]model.Cell, len(rectMap)/2)
	}

	if err := validateMapRectangular(rectMap); err != nil {
		t.Errorf("validateMapRectangular on a rectangular map returned an error: %v", err)
	}
}

func TestValidateCells(t *testing.T) {
	max := int(model.CellPlayer)

	lvl := make([][]model.Cell, 10)

	current := 0

	for i := range lvl {
		lvl[i] = make([]model.Cell, 10)

		for j := range lvl[i] {
			c := int32(current)
			cell, err := model.NewCell(c)

			if err != nil {
				t.Errorf("could not create cell. value: %d | row: %d | col: %d", current, i, j)
			}

			lvl[i][j] = cell

			current += 1
			current %= max
		}
	}

	lvl[0][0] = model.CellPlayer // one player on map

	if _, _, err := validateCells(lvl); err != nil {
		t.Errorf("validateCells on a valid map returned an error: %v", err)
	}

	lvl[0][0] = model.Cell(max + 1) // invalid level.Cell

	if _, _, err := validateCells(lvl); err == nil {
		t.Error("validateCells on a invalid map did not return an error")
	}
}
