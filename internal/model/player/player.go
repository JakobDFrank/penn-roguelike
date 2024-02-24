package player

import (
	"gorm.io/gorm"
)

const (
	_startingHitpoints = 4
)

type Player struct {
	gorm.Model
	Hitpoints int

	LevelID uint
}

func NewPlayer(levelID uint) *Player {
	return &Player{
		Hitpoints: _startingHitpoints,
		LevelID:   levelID,
	}
}
