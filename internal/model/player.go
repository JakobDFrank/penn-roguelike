package model

import (
	"gorm.io/gorm"
)

const (
	_startingHitpoints = 4
)

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
