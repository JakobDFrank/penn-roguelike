package controller

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type LevelController struct {
	db     *gorm.DB
	logger *zap.Logger
}

type InsertLevelResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewLevelController(logger *zap.Logger, db *gorm.DB) (*LevelController, error) {
	if db == nil {
		return nil, &apperr.NilArgumentError{Message: "db"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	lc := &LevelController{
		db:     db,
		logger: logger,
	}

	return lc, nil
}

func (lc *LevelController) SubmitLevel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(lc.logger, w, err)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			lc.logger.Error("close_body", zap.Error(err))
		}
	}()

	cells := make([][]model.Cell, 0)
	if err := json.Unmarshal(body, &cells); err != nil {
		handleError(lc.logger, w, err)
		return
	}

	lc.logger.Debug("unmarshalled_level", zap.Any("cells", cells))

	level := &model.Level{Cells: cells}

	if err := validateLevel(level); err != nil {
		handleError(lc.logger, w, err)
		return
	}

	res := lc.db.FirstOrCreate(level, &model.Level{Cells: cells})

	if res.Error != nil {
		handleError(lc.logger, w, err)
		return
	}

	lc.logger.Debug("submit_level", zap.Int64("rows_affected", res.RowsAffected))

	msg := fmt.Sprintf("rows_affected: %d", res.RowsAffected)
	resp := InsertLevelResponse{
		Id:      level.ID,
		Message: msg,
		Status:  http.StatusOK,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func validateLevel(level *model.Level) error {
	cells := level.Cells

	rowCount := len(cells)
	if rowCount == 0 {
		return apperr.ErrEmptyMap
	}

	// validate map size, don't want to iterate over potentially massive array

	if rowCount > 100 {
		return apperr.ErrMapTooLarge
	}

	// validate rectangular map
	expectedColCount := len(cells[0])

	if expectedColCount > 100 {
		return apperr.ErrMapTooLarge
	}

	for _, row := range cells[1:] {
		colCount := len(row)

		if colCount != expectedColCount {
			return apperr.ErrMapNotRectangular
		}
	}

	// validate cells after ensuring map is rectangular
	for i, row := range cells {
		for j, cell := range row {
			if !cell.IsValid() {
				return &apperr.InvalidCellTypeError{Message: fmt.Sprintf("cell value: %d | row: %d | col: %d", cell, i, j)}
			}
		}
	}

	return nil
}

func handleError(logger *zap.Logger, w http.ResponseWriter, err error) {

	logger.Error("handling_error", zap.Error(err))

	resp := InsertLevelResponse{
		Id:      0,
		Message: err.Error(),
		Status:  http.StatusBadRequest,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		logger.Error("marshal_error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
