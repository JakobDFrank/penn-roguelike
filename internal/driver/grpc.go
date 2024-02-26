package driver

import (
	"context"
	"encoding/json"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"github.com/JakobDFrank/penn-roguelike/internal/rpc"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// GrpcDriver handles gRPC calls.
type GrpcDriver struct {
	lc *service.LevelService
	pc *service.PlayerService

	rpc.UnimplementedLevelServiceServer
	rpc.UnimplementedPlayerServiceServer

	logger *zap.Logger
}

var _ Driver = (*GrpcDriver)(nil)

func NewGrpcDriver(lc *service.LevelService, pc *service.PlayerService, logger *zap.Logger) (*GrpcDriver, error) {
	if lc == nil {
		return nil, &apperr.NilArgumentError{Message: "lc"}
	}

	if pc == nil {
		return nil, &apperr.NilArgumentError{Message: "pc"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	gd := &GrpcDriver{
		lc:     lc,
		pc:     pc,
		logger: logger,
	}

	return gd, nil
}

func (gd *GrpcDriver) CreateLevel(_ context.Context, req *rpc.CreateLevelRequest) (*rpc.CreateLevelResponse, error) {

	cells := make([][]model.Cell, 0)

	for _, gRows := range req.Level {
		row := make([]model.Cell, 0)

		for _, gCell := range gRows.Cells {
			cell, err := model.NewCell(int(gCell))

			if err != nil {
				return nil, err
			}

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	id, err := gd.lc.SubmitLevel(cells)

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

	gameMap, err := gd.pc.MovePlayer(uint(req.Id), dir)

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

func (gd *GrpcDriver) Serve() error {

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		gd.logger.Error("failed to listen", zap.Error(err))
		return err
	}

	s := grpc.NewServer()

	rpc.RegisterLevelServiceServer(s, gd)
	rpc.RegisterPlayerServiceServer(s, gd)

	reflection.Register(s)

	gd.logger.Info("grpc_server_listening", zap.Any("addr", lis.Addr()))

	if err := s.Serve(lis); err != nil {
		gd.logger.Error("failed to serve", zap.Error(err))
		return err
	}

	return nil
}
