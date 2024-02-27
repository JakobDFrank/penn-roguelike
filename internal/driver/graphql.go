package driver

import (
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/JakobDFrank/penn-roguelike/api/graphql/graph"
	gqlmodel "github.com/JakobDFrank/penn-roguelike/api/graphql/graph/model"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//--------------------------------------------------------------------------------
// GraphQLDriver
//--------------------------------------------------------------------------------

const (
	_graphQLPort     = 9091
	_graphQLEndpoint = "/query"
)

// GraphQLDriver handles GraphQL calls.
type GraphQLDriver struct {
	levelService  *service.LevelService
	playerService *service.PlayerService
	logger        *zap.Logger
	graph.ResolverRoot
}

// NewGraphQLDriver creates a new instance of GraphQLDriver.
func NewGraphQLDriver(levelService *service.LevelService, playerService *service.PlayerService, logger *zap.Logger) (*GraphQLDriver, error) {

	if levelService == nil {
		return nil, &apperr.NilArgumentError{Message: "levelService"}
	}

	if playerService == nil {
		return nil, &apperr.NilArgumentError{Message: "playerService"}
	}

	if logger == nil {
		return nil, &apperr.NilArgumentError{Message: "logger"}
	}

	gd := &GraphQLDriver{
		levelService:  levelService,
		playerService: playerService,
		logger:        logger,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gd}))

	http.Handle("/", playground.Handler("GraphQL playground", _graphQLEndpoint))
	http.Handle(_graphQLEndpoint, srv)

	return gd, nil
}

// InsertLevel is the resolver for the insertLevel field in the GraphQL schema.
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

	id, err := gd.levelService.SubmitLevel(cells)

	if err != nil {
		return "", err
	}

	return strconv.FormatUint(uint64(id), 10), nil
}

// MovePlayer is the resolver for the movePlayer field in the GraphQL schema.
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

	gameMap, err := gd.playerService.MovePlayer(uint(mapId), d)

	if err != nil {
		return "", err
	}

	js, err := json.Marshal(gameMap)

	if err != nil {
		return "", err
	}

	return string(js), nil
}

func (gd *GraphQLDriver) Serve(onExitCtx context.Context) error {
	gd.logger.Info("graphql_server_start_listening")

	return httpGracefulServe(_graphQLPort, onExitCtx, gd.logger)
}

func (gd *GraphQLDriver) Mutation() graph.MutationResolver {
	return gd
}

var _ Driver = (*GraphQLDriver)(nil)
