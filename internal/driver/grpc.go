package driver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/api/rpc"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

//--------------------------------------------------------------------------------
// GrpcDriver
//--------------------------------------------------------------------------------

const (
	_grpcPort = 9090
)

// GrpcDriver handles gRPC calls.
type GrpcDriver struct {
	levelService  *service.LevelService
	playerService *service.PlayerService

	rpc.UnimplementedLevelServiceServer
	rpc.UnimplementedPlayerServiceServer

	logger *zap.Logger
}

// NewGrpcDriver creates a new instance of GrpcDriver.
func NewGrpcDriver(levelService *service.LevelService, playerService *service.PlayerService, logger *zap.Logger) (*GrpcDriver, error) {
	if levelService == nil {
		return nil, &apperr.NilArgumentError{Message: "levelService"}
	}

	if playerService == nil {
		return nil, &apperr.NilArgumentError{Message: "playerService"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	gd := &GrpcDriver{
		levelService:  levelService,
		playerService: playerService,
		logger:        logger,
	}

	return gd, nil
}

func (gd *GrpcDriver) CreateLevel(_ context.Context, req *rpc.CreateLevelRequest) (*rpc.CreateLevelResponse, error) {

	cells := make([][]model.Cell, 0)

	for _, gRows := range req.Level {
		row := make([]model.Cell, 0)

		for _, gCell := range gRows.Cells {
			cell, err := model.NewCell(gCell)

			if err != nil {
				return nil, err
			}

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	id, err := gd.levelService.SubmitLevel(cells)

	if err != nil {
		return nil, err
	}

	resp := &rpc.CreateLevelResponse{Id: uint32(id)}

	return resp, nil
}

func (gd *GrpcDriver) MovePlayer(_ context.Context, req *rpc.MovePlayerRequest) (*rpc.MovePlayerResponse, error) {

	var dir model.Direction

	switch req.Direction {
	case rpc.Direction_LEFT:
		dir = model.Left
	case rpc.Direction_UP:
		dir = model.Up
	case rpc.Direction_RIGHT:
		dir = model.Right
	case rpc.Direction_DOWN:
		dir = model.Down
	default:
		return nil, &apperr.UnimplementedError{Message: "[MovePlayer] Direction"}
	}

	gameMap, err := gd.playerService.MovePlayer(req.Id, dir)

	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(gameMap)

	if err != nil {
		return nil, err
	}

	resp := &rpc.MovePlayerResponse{
		Map: string(js),
	}

	return resp, nil
}

func (gd *GrpcDriver) Serve(onExitCtx context.Context) error {

	addr := fmt.Sprintf(":%d", _grpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		gd.logger.Error("failed to listen", zap.Error(err))
		return err
	}

	s := grpc.NewServer()

	rpc.RegisterLevelServiceServer(s, gd)
	rpc.RegisterPlayerServiceServer(s, gd)

	reflection.Register(s)

	gd.logger.Info("grpc_server_listening", zap.Any("addr", lis.Addr()))

	go func() {
		<-onExitCtx.Done()
		gd.logger.Warn("shutting_down_server")

		s.GracefulStop()

		gd.logger.Warn("server_exit")
	}()

	if err := s.Serve(lis); err != nil {
		gd.logger.Error("failed to serve", zap.Error(err))
		return err
	}

	return nil
}

var _ Driver = (*GrpcDriver)(nil)
