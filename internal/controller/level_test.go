package controller

import (
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"testing"
)

func TestValidateMapSize(t *testing.T) {
	emptyMap := make(model.Cells, 0)

	if err := validateMapSize(emptyMap); err == nil {
		t.Error("validateMapSize on an empty map did not return an error")
	}

	largeMap := make(model.Cells, MaxLevelSize+1)

	if err := validateMapSize(largeMap); err == nil {
		t.Error("validateMapSize on a map too large did not return an error")
	}
}

func TestValidateMapRectangular(t *testing.T) {
	jaggedMap := make(model.Cells, 10)

	for i := range jaggedMap {
		jaggedMap[i] = make([]model.Cell, i+1)
	}

	if err := validateMapRectangular(jaggedMap); err == nil {
		t.Error("validateMapRectangular on a jagged map did not return an error")
	}

	squareMap := make(model.Cells, 10)

	for i := range squareMap {
		squareMap[i] = make([]model.Cell, len(squareMap))
	}

	if err := validateMapRectangular(squareMap); err != nil {
		t.Errorf("validateMapRectangular on a square map returned an error: %v", err)
	}

	rectMap := make(model.Cells, 10)

	for i := range rectMap {
		rectMap[i] = make([]model.Cell, len(rectMap)/2)
	}

	if err := validateMapRectangular(rectMap); err != nil {
		t.Errorf("validateMapRectangular on a rectangular map returned an error: %v", err)
	}
}

func TestValidateCells(t *testing.T) {
	max := int(model.StartPosition) + 1

	level := make(model.Cells, 10)

	current := 0

	for i := range level {
		level[i] = make([]model.Cell, 10)

		for j := range level[i] {
			c, err := model.NewCell(current)

			if err != nil {
				t.Errorf("could not create cell. value: %d | row: %d | col: %d", current, i, j)
			}

			level[i][j] = c

			current += 1
			current %= max
		}
	}

	if err := validateCells(level); err != nil {
		t.Errorf("validateCells on a valid map returned an error: %v", err)
	}

	level[0][0] = model.Cell(max)

	if err := validateCells(level); err == nil {
		t.Error("validateCells on a invalid map did not return an error")
	}
}
