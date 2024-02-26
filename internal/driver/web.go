package driver

import (
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/controller"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"go.uber.org/zap"
	"net/http"
)

const (
	Kibibyte = 1024
	Mebibyte = Kibibyte * Kibibyte
)

// WebDriver handles HTTP requests
type WebDriver struct {
	lc     *controller.LevelController
	pc     *controller.PlayerController
	logger *zap.Logger
}

var _ Driver = (*WebDriver)(nil)

func NewWebDriver(lc *controller.LevelController, pc *controller.PlayerController, logger *zap.Logger) (*WebDriver, error) {

	if lc == nil {
		return nil, &apperr.NilArgumentError{Message: "lc"}
	}

	if pc == nil {
		return nil, &apperr.NilArgumentError{Message: "pc"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	wd := &WebDriver{
		lc:     lc,
		pc:     pc,
		logger: logger,
	}

	http.HandleFunc("/level/submit", wd.SubmitLevel)
	http.HandleFunc("/player/move", wd.MovePlayer)

	return wd, nil
}

// SubmitLevel handles HTTP requests to insert levels that can be played.
// It returns the unique ID of the level or an error.
func (wd *WebDriver) SubmitLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cells := make([][]model.Cell, 0)

	if err := deserializePostRequest(w, r, &cells); err != nil {
		wd.handleError(w, err)
		return
	}

	id, err := wd.lc.SubmitLevel(cells)

	if err != nil {
		wd.handleError(w, err)
		return
	}

	resp := InsertLevelResponse{
		Id:      id,
		Message: "",
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

// MovePlayer handles HTTP requests to move a player within the game.
// It returns the new game state or an error.
func (wd *WebDriver) MovePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	moveRequest := MovePlayerRequest{}

	if err := deserializePostRequest(w, r, &moveRequest); err != nil {
		wd.handleError(w, err)
		return
	}

	dir, err := model.NewDirection(moveRequest.Direction)

	if err != nil {
		wd.handleError(w, err)
		return
	}

	cellJson, err := wd.pc.MovePlayer(moveRequest.ID, dir)

	if err != nil {
		wd.handleError(w, err)
		return
	}

	wd.logger.Debug("unmarshal", zap.Any("move_request", moveRequest))

	resp := MovePlayerResponse{
		Id:     moveRequest.ID,
		Level:  cellJson,
		Status: http.StatusOK,
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

func (wd *WebDriver) Serve() error {
	wd.logger.Debug("Listening...")

	return http.ListenAndServe(":8080", nil)
}

func deserializePostRequest(w http.ResponseWriter, r *http.Request, value any) error {
	if r.Method != http.MethodPost {
		return &apperr.InvalidArgumentError{Message: fmt.Sprintf("invalid_method: %s", r.Method)}
	}

	// debug only? reflection a bit slow
	//if reflect.ValueOf(value).Kind() != reflect.Ptr {
	//	return &apperr.InvalidArgumentError{Message: "value must be a pointer"}
	//}

	// limit request size
	r.Body = http.MaxBytesReader(w, r.Body, Mebibyte)

	err := json.NewDecoder(r.Body).Decode(value)
	if err != nil {
		return err
	}

	// only return error in debug? log otherwise
	return r.Body.Close()
}

func (wd *WebDriver) handleError(w http.ResponseWriter, err error) {

	wd.logger.Error("handling_error", zap.Error(err))

	resp := ErrorResponse{
		Message: err.Error(),
		Status:  http.StatusBadRequest,
	}

	jsonData, err := json.Marshal(resp)

	if err != nil {
		wd.logger.Error("marshal_error", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//--------------------------------------------------------------------------------
// InsertLevelResponse
//--------------------------------------------------------------------------------

type InsertLevelResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//--------------------------------------------------------------------------------
// MovePlayerRequest
//--------------------------------------------------------------------------------

type MovePlayerRequest struct {
	ID        uint `json:"id"`
	Direction int  `json:"direction"`
}

//--------------------------------------------------------------------------------
// MovePlayerResponse
//--------------------------------------------------------------------------------

type MovePlayerResponse struct {
	Id     uint   `json:"id"`
	Level  string `json:"level"`
	Status int    `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
