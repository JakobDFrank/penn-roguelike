package driver

import (
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/graphql/graph"
	gqlmodel "github.com/JakobDFrank/penn-roguelike/internal/graphql/graph/model"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type GraphQLDriver struct {
	lc     *service.LevelService
	pc     *service.PlayerService
	logger *zap.Logger
	graph.ResolverRoot
}

func (gd *GraphQLDriver) Mutation() graph.MutationResolver {
	return gd
}

// InsertLevel is the resolver for the insertLevel field.
func (gd *GraphQLDriver) InsertLevel(ctx context.Context, level [][]int) (string, error) {
	cells := make([][]model.Cell, 0)

	for _, gRows := range level {
		row := make([]model.Cell, 0)

		for _, gCell := range gRows {
			cell, err := model.NewCell(gCell)

			if err != nil {
				return "", err
			}

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	id, err := gd.lc.SubmitLevel(cells)

	if err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(id), 10), nil
}

// MovePlayer is the resolver for the movePlayer field.
func (gd *GraphQLDriver) MovePlayer(ctx context.Context, id string, dir gqlmodel.Direction) (string, error) {
	var d model.Direction

	switch dir {
	case gqlmodel.DirectionLeft:
		d = model.Left
	case gqlmodel.DirectionRight:
		d = model.Right
	case gqlmodel.DirectionUp:
		d = model.Up
	case gqlmodel.DirectionDown:
		d = model.Down
	default:
		return "", &apperr.UnimplementedError{Message: "[MovePlayer] Direction"}
	}

	mapId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		return "", err
	}

	gameMap, err := gd.pc.MovePlayer(uint(mapId), d)

	if err != nil {
		return "", err
	}

	js, err := json.Marshal(gameMap)

	if err != nil {
		return "", err
	}

	return string(js), nil
}

var _ Driver = (*GraphQLDriver)(nil)

func NewGraphQLDriver(lc *service.LevelService, pc *service.PlayerService, logger *zap.Logger) (*GraphQLDriver, error) {

	if lc == nil {
		return nil, &apperr.NilArgumentError{Message: "lc"}
	}

	if pc == nil {
		return nil, &apperr.NilArgumentError{Message: "pc"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	wd := &GraphQLDriver{
		lc:     lc,
		pc:     pc,
		logger: logger,
	}

	return wd, nil
}

func (gd *GraphQLDriver) Serve() error {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gd}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	return http.ListenAndServe(":9091", nil)
}
