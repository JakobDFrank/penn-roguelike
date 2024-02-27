package model

import (
	"testing"
)

func TestValidateMapSize(t *testing.T) {

	t.Run("test_empty_map", func(t *testing.T) {
		emptyMap := make(GameMap, 0)

		if err := validateMapSize(emptyMap); err == nil {
			t.Error("validateMapSize on an empty map did not return an error")
		}
	})

	t.Run("test_large_map", func(t *testing.T) {

		largeMap := make(GameMap, MaxLevelSize+1)

		if err := validateMapSize(largeMap); err == nil {
			t.Error("validateMapSize on a map too large did not return an error")
		}
	})
}

func TestValidateMapRectangular(t *testing.T) {

	t.Run("test_jagged_map", func(t *testing.T) {
		jaggedMap := make(GameMap, 10)

		for i := range jaggedMap {
			jaggedMap[i] = make([]Cell, i+1)
		}

		if err := validateMapRectangular(jaggedMap); err == nil {
			t.Error("validateMapRectangular on a jagged map did not return an error")
		}
	})

	t.Run("test_square_map", func(t *testing.T) {

		squareMap := make(GameMap, 10)

		for i := range squareMap {
			squareMap[i] = make([]Cell, len(squareMap))
		}

		if err := validateMapRectangular(squareMap); err != nil {
			t.Errorf("validateMapRectangular on a square map returned an error: %v", err)
		}
	})

	t.Run("test_valid_map", func(t *testing.T) {

		rectMap := make(GameMap, 10)

		for i := range rectMap {
			rectMap[i] = make([]Cell, len(rectMap)/2)
		}

		if err := validateMapRectangular(rectMap); err != nil {
			t.Errorf("validateMapRectangular on a rectangular map returned an error: %v", err)
		}
	})
}

func TestValidateCells(t *testing.T) {

	max := int(CellPlayer)
	lvl := make(GameMap, 10)

	t.Run("test_valid_map", func(t *testing.T) {

		current := 0

		for i := range lvl {
			lvl[i] = make([]Cell, 10)

			for j := range lvl[i] {
				c, err := NewCell(current)

				if err != nil {
					t.Errorf("could not create cell. value: %d | row: %d | col: %d", current, i, j)
				}

				lvl[i][j] = c

				current += 1
				current %= max
			}
		}

		lvl[0][0] = CellPlayer // one player on map

		if _, err := validateCells(lvl); err != nil {
			t.Errorf("validateCells on a valid map returned an error: %v", err)
		}
	})

	t.Run("test_invalid_map", func(t *testing.T) {

		lvl[0][0] = Cell(max + 1) // invalid level.Cell

		if _, err := validateCells(lvl); err == nil {
			t.Error("validateCells on a invalid map did not return an error")
		}
	})
}
