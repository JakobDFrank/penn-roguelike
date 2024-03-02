package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/analytics"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"go.uber.org/zap"
	"net/http"
)

const (
	_httpPort = 8080
	Kibibyte  = 1024
	Mebibyte  = Kibibyte * Kibibyte
)

//--------------------------------------------------------------------------------
// WebDriver
//--------------------------------------------------------------------------------

// WebDriver handles HTTP requests.
type WebDriver struct {
	levelService  *service.LevelService
	playerService *service.PlayerService
	logger        *zap.Logger
	obs           analytics.Collector
}

// NewWebDriver creates a new instance of WebDriver.
func NewWebDriver(logger *zap.Logger, obs analytics.Collector, levelService *service.LevelService, playerService *service.PlayerService) (*WebDriver, error) {

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	if obs == nil {
		return nil, &apperr.NilArgumentError{Message: "obs"}
	}

	if levelService == nil {
		return nil, &apperr.NilArgumentError{Message: "levelService"}
	}

	if playerService == nil {
		return nil, &apperr.NilArgumentError{Message: "playerService"}
	}

	wd := &WebDriver{
		levelService:  levelService,
		playerService: playerService,
		logger:        logger,
		obs:           obs,
	}

	return wd, nil
}

func (wd *WebDriver) SubmitLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cells := make([][]model.Cell, 0)

	if err := deserializePostRequest(w, r, &cells); err != nil {
		wd.handleError(w, err)
		return
	}

	id, err := wd.levelService.SubmitLevel(cells)

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

func (wd *WebDriver) MovePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	moveRequest := MovePlayerRequest{}

	if err := deserializePostRequest(w, r, &moveRequest); err != nil {
		wd.handleError(w, err)
		return
	}

	gameMap, err := wd.playerService.MovePlayer(moveRequest.ID, moveRequest.Direction)

	if err != nil {
		wd.handleError(w, err)
		return
	}

	js, err := json.Marshal(gameMap)

	if err != nil {
		wd.handleError(w, err)
		return
	}

	wd.logger.Debug("unmarshal", zap.Any("move_request", moveRequest))

	resp := MovePlayerResponse{
		Id:     moveRequest.ID,
		Level:  string(js),
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

func (wd *WebDriver) Serve(onExitCtx context.Context) error {
	wd.logger.Info("http_server_start_listening")

	mux := http.NewServeMux()

	mux.HandleFunc("/level/submit", wd.SubmitLevel)
	mux.HandleFunc("/player/move", wd.MovePlayer)

	handler := analyticsMiddleware(wd.obs, mux)

	return httpGracefulServe(_httpPort, handler, onExitCtx, wd.logger)
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

var _ Driver = (*WebDriver)(nil)

//--------------------------------------------------------------------------------
// InsertLevelResponse
//--------------------------------------------------------------------------------

type InsertLevelResponse struct {
	Id      int32  `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

//--------------------------------------------------------------------------------
// MovePlayerRequest
//--------------------------------------------------------------------------------

type MovePlayerRequest struct {
	ID        int32           `json:"id"`
	Direction model.Direction `json:"direction"`
}

//--------------------------------------------------------------------------------
// MovePlayerResponse
//--------------------------------------------------------------------------------

type MovePlayerResponse struct {
	Id     int32  `json:"id"`
	Level  string `json:"level"`
	Status int    `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
