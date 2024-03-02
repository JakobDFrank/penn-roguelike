// Package model contains entity definitions and basic entity management.
package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"strings"
)

//--------------------------------------------------------------------------------
// LevelRepository
//--------------------------------------------------------------------------------

// LevelRepository handles CRUD operations for model.Level.
type LevelRepository interface {
	Begin() (*sql.Tx, error)
	CreateLevelWithTx(tx *sql.Tx, cells [][]Cell) (Level, error) // CreateLevelWithTx creates a Level within the given transaction.
	GetFirst(id int32) (Level, error)                            // GetFirst retrieves a Level by its ID.
}

//--------------------------------------------------------------------------------
// PlayerRepository
//--------------------------------------------------------------------------------

// PlayerRepository handles CRUD operations for model.Player.
type PlayerRepository interface {
	Begin() (*sql.Tx, error)
	CreatePlayerWithTx(tx *sql.Tx, levelId CreatePlayerParams) (Player, error) // CreatePlayerWithTx creates a Player within the given transaction.
	GetPlayerByLevelIDAndLockWithTx(tx *sql.Tx, levelId int32) (Player, error) // GetPlayerByLevelIDAndLockWithTx retrieves a Player by level ID and locks the Player row, within the given transaction.
	UpdatePlayer(tx *sql.Tx, params UpdatePlayerParams) error                  // UpdatePlayer updates a Player within the given transaction.
}

//--------------------------------------------------------------------------------
// Cell
//--------------------------------------------------------------------------------

// Cell represents different types of objects found within a game level.
type Cell int32

const (
	CellOpen   Cell = iota // An open tile
	CellWall               // An impassable barrier
	CellPit                // A hazard that does one damage
	CellArrow              // A hazard that does two damage
	CellPlayer             // The player's character
)

// NewCell creates an implementation of Cell.
func NewCell(cell int32) (Cell, error) {
	c := Cell(cell)

	if c.IsValid() {
		return c, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("cell: %d", cell)}
}

// IsValid verifies if the Cell instance is a valid Cell member.
func (c *Cell) IsValid() bool {
	switch *c {
	case CellOpen, CellWall, CellPit, CellArrow, CellPlayer:
		return true
	default:
		return false
	}
}

//--------------------------------------------------------------------------------
// Direction
//--------------------------------------------------------------------------------

const (
	_left  = "left"
	_right = "right"
	_up    = "up"
	_down  = "down"
)

// Direction represents different moves the CellPlayer can make on the map
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

// NewDirection creates a new instance of Direction
func NewDirection(dir int32) (Direction, error) {
	d := Direction(dir)

	if d.IsValid() {
		return d, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("direction: %d", dir)}
}

// IsValid verifies if the Direction instance is a valid Direction member.
func (d *Direction) IsValid() bool {
	switch *d {
	case Left, Right, Up, Down:
		return true
	default:
		return false
	}
}

func (d *Direction) String() string {

	switch *d {
	case Left:
		return _left
	case Right:
		return _right
	case Up:
		return _up
	case Down:
		return _down
	default:
		// panic in debug
		return "unimplemented"
	}

}

func ParseDirection(name string) (Direction, error) {
	name = strings.ToLower(name)

	switch name {
	case _left:
		return Left, nil
	case _right:
		return Right, nil
	case _up:
		return Up, nil
	case _down:
		return Down, nil
	default:
		return 0, &apperr.InvalidArgumentError{Message: "name"}
	}
}

func (d *Direction) UnmarshalJSON(data []byte) error {

	var num int32

	// check for integer first
	if err := json.Unmarshal(data, &num); err == nil {
		dir, err := NewDirection(num)

		if err != nil {
			return err
		}

		*d = dir
		return nil
	}

	// if not an integer, check for string representation

	var str string

	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	val, err := ParseDirection(str)

	if err != nil {
		return err
	}

	*d = val
	return nil
}

var _ json.Unmarshaler = (*Direction)(nil)
