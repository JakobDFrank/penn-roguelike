package service

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/analytics"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"go.uber.org/zap"
	"strings"
)

//--------------------------------------------------------------------------------
// LevelService
//--------------------------------------------------------------------------------

// LevelService handles level management.
type LevelService struct {
	levelRepo  model.LevelRepository
	playerRepo model.PlayerRepository

	logger *zap.Logger
	obs    analytics.Collector
}

// NewLevelService creates a new instance of LevelService.
func NewLevelService(logger *zap.Logger, obs analytics.Collector, levelRepo model.LevelRepository, playerRepo model.PlayerRepository) (*LevelService, error) {

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	if obs == nil {
		return nil, &apperr.NilArgumentError{Message: "obs"}
	}

	if levelRepo == nil {
		return nil, &apperr.NilArgumentError{Message: "levelRepo"}
	}

	if playerRepo == nil {
		return nil, &apperr.NilArgumentError{Message: "playerRepo"}
	}

	lc := &LevelService{
		levelRepo:  levelRepo,
		playerRepo: playerRepo,
		logger:     logger,
		obs:        obs,
	}

	return lc, nil
}

// SubmitLevel inserts levels that can be played.
// It returns the unique ID of the level or an error.
func (ls *LevelService) SubmitLevel(cells [][]model.Cell) (int32, error) {

	ls.logger.Debug("unmarshalled_level", zap.Any("cells", cells))

	x, y, err := validateMap(cells)

	if err != nil {
		return 0, err
	}

	id, err := ls.createMap(cells, x, y)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ls *LevelService) createMap(cells [][]model.Cell, rowStartIdx, colStartIdx int32) (int32, error) {

	// saving this within the player
	cells[rowStartIdx][colStartIdx] = model.CellOpen

	tx, err := ls.levelRepo.Begin()

	if err != nil {
		return 0, err
	}

	lvl, err := ls.levelRepo.CreateLevelWithTx(tx, cells)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	playr, err := ls.playerRepo.CreatePlayerWithTx(tx, model.CreatePlayerParams{
		LevelID:   lvl.ID,
		StartX:    colStartIdx,
		StartY:    rowStartIdx,
		CurrX:     colStartIdx,
		CurrY:     rowStartIdx,
		Hitpoints: _startingHitpoints,
	})

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	ls.logger.Debug("submit_level")

	PrintMap(&lvl, &playr, ls.logger)

	return lvl.ID, nil
}

//--------------------------------------------------------------------------------
// Level
//--------------------------------------------------------------------------------

const (
	MaxLevelSize         = 100
	_expectedPlayerCount = 1
)

// Level is an entity in the database that holds information on map data.

func validateMap(gameMap [][]model.Cell) (int32, int32, error) {

	// validate map size, don't want to iterate over potentially massive array
	if err := validateMapSize(gameMap); err != nil {
		return 0, 0, err
	}

	// validate rectangular map
	if err := validateMapRectangular(gameMap); err != nil {
		return 0, 0, err
	}

	// validate gameMap after ensuring map is rectangular, ensure only one player position
	x, y, err := validateCells(gameMap)

	if err != nil {
		return 0, 0, err
	}

	return x, y, nil
}

func validateMapSize(gameMap [][]model.Cell) error {
	rowCount := len(gameMap)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	if rowCount > MaxLevelSize {
		return apperr.ErrMapTooLarge
	}

	expectedColCount := len(gameMap[0])

	if expectedColCount > MaxLevelSize {
		return apperr.ErrMapTooLarge
	}

	return nil
}

func validateMapRectangular(gameMap [][]model.Cell) error {

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

func validateCells(gameMap [][]model.Cell) (int32, int32, error) {

	playerCount := 0

	var x, y int32

	for rowIdx, row := range gameMap {
		for colIdx, cell := range row {

			if !cell.IsValid() {
				return 0, 0, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("cell value: %d | row: %d | col: %d", cell, rowIdx, colIdx)}
			}

			if cell == model.CellPlayer {
				playerCount += 1

				x = int32(rowIdx)
				y = int32(colIdx)

				if playerCount > 1 {
					return 0, 0, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("more than one player in map | row: %d | col: %d", rowIdx, colIdx)}
				}
			}
		}
	}

	if playerCount != _expectedPlayerCount {
		return 0, 0, &apperr.InvalidCellTypeError{Message: fmt.Sprintf("unexpected player count: %d (expected: %d)", playerCount, _expectedPlayerCount)}
	}

	return x, y, nil
}

func SerializeCellsWithPlayer(cells [][]model.Cell, player *model.Player, logger *zap.Logger) ([]byte, error) {

	oldCell := cells[player.CurrY][player.CurrX]
	cells[player.CurrY][player.CurrX] = model.CellPlayer

	if logger.Level() == zap.DebugLevel {
		var sb strings.Builder
		for idx, row := range cells {
			for _, element := range row {
				sb.WriteString(fmt.Sprintf("%4d", element))
			}

			rowText := fmt.Sprintf("row_%d", idx)
			logger.Debug("print_level", zap.String(rowText, sb.String()))
			sb.Reset()
		}
	}

	jsonText, err := json.Marshal(&cells)

	if err != nil {
		return nil, err
	}

	cells[player.CurrY][player.CurrX] = oldCell

	return jsonText, nil
}

func PrintMap(lvl *model.Level, player *model.Player, logger *zap.Logger) {

	cells := make([][]model.Cell, 0)

	if err := json.Unmarshal(lvl.Map, &cells); err != nil {
		logger.Error("unmarshal", zap.Error(err))
		return
	}

	if _, err := SerializeCellsWithPlayer(cells, player, logger); err != nil {
		logger.Error("print_map", zap.Error(err))
	}
}
