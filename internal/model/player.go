package model

import (
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"gorm.io/gorm"
)

//--------------------------------------------------------------------------------
// Player
//--------------------------------------------------------------------------------

const (
	_startingHitpoints = 4
)

// Player is an entity in the database that holds information on player data
type Player struct {
	gorm.Model
	Hitpoints int
	RowIdx    int
	ColIdx    int

	LevelID uint
}

func NewPlayer(levelID uint, startRowIdx, startColIdx int) *Player {
	return &Player{
		Hitpoints: _startingHitpoints,
		RowIdx:    startRowIdx,
		ColIdx:    startColIdx,

		LevelID: levelID,
	}
}

func (p *Player) ResetHitpoints() {
	p.Hitpoints = _startingHitpoints
}

//--------------------------------------------------------------------------------
// Direction
//--------------------------------------------------------------------------------

// Direction represents different moves the CellPlayer can make on the map
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

func NewDirection(dir int) (Direction, error) {
	d := Direction(dir)

	if d.IsValid() {
		return d, nil
	}

	return 0, &apperr.InvalidArgumentError{Message: fmt.Sprintf("direction: %d", dir)}
}

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
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	default:
		return "unimplemented"
	}
}
