package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/analytics"
	"github.com/JakobDFrank/penn-roguelike/internal/apperr"
	"github.com/JakobDFrank/penn-roguelike/internal/database/model"
	"github.com/JakobDFrank/penn-roguelike/internal/driver"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// environment variables
const (
	DbHostEnvVar = "DB_HOST"
	DbUserEnvVar = "DB_USER"
	DbPassEnvVar = "DB_PASSWORD"
	DbNameEnvVar = "DB_NAME"
)

type driverKindFlag struct {
	driver.DriverKind
}

func (s *driverKindFlag) Set(text string) error {
	svc, exists := driver.DriverNameToEnum[strings.ToLower(text)]
	if !exists {
		return &apperr.InvalidArgumentError{Message: fmt.Sprintf("invalid service: %s", text)}
	}
	s.DriverKind = svc
	return nil
}

func (s *driverKindFlag) String() string {
	for k, v := range driver.DriverNameToEnum {
		if v == s.DriverKind {
			return k
		}
	}
	return ""
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	ctx, cancel := setupGracefulExit(logger)
	defer cancel()

	var svc driverKindFlag
	flag.Var(&svc, "api", "Set the API to use (http, grpc, or graphql)")
	flag.Parse()

	db, err := setupDatabase(logger)

	if err != nil {
		logger.Fatal("setup_db", zap.Error(err))
	}

	defer db.Close()

	driver, err := setupHandlers(svc.DriverKind, logger, db)
	if err != nil {
		logger.Fatal("setup_server", zap.Error(err))
	}

	if err := driver.Serve(ctx); err != nil {
		logger.Fatal("serve", zap.Error(err))
	}
}

func setupGracefulExit(logger *zap.Logger) (context.Context, context.CancelFunc) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-ctx.Done()
		logger.Info("wait_for_cleanup")
		time.Sleep(time.Second * 3)
		logger.Warn("exiting...")
		os.Exit(0)
	}()

	return ctx, cancel
}

//go:embed internal/database/migrations/*.sql
var migrations embed.FS

func setupDatabase(logger *zap.Logger) (*sql.DB, error) {
	dbHost := os.Getenv(DbHostEnvVar)
	dbUser := os.Getenv(DbUserEnvVar)
	dbName := os.Getenv(DbNameEnvVar)
	dbPass := os.Getenv(DbPassEnvVar)

	logger.Info("db_env_variables", zap.String("host", dbHost), zap.String("user", dbUser), zap.String("name", dbName))
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbUser, dbName, dbPass)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		logger.Error("sql_open", zap.Error(err))
		return nil, err
	}

	attempt := 0

	for {
		logger.Debug("open_db", zap.Int("attempt", attempt+1))

		if err := db.Ping(); err == nil {
			break
		}

		if attempt >= 10 {
			logger.Error("db_not_active", zap.Error(err))
			return nil, err
		}

		attempt += 1
		time.Sleep(time.Second)
	}

	if err := handleSchemaMigration(db, dbName); err != nil {
		logger.Error("schema_migration", zap.Error(err))
		return nil, err
	}

	return db, nil
}

func handleSchemaMigration(db *sql.DB, dbName string) error {
	tempDir, err := os.MkdirTemp("", "schema")
	if err != nil {
		return err
	}

	// recreate migrations directory in temp folder - small size, should be fine
	if err := fs.WalkDir(migrations, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			content, err := migrations.ReadFile(path)
			if err != nil {
				return err
			}

			tempFilePath := filepath.Join(tempDir, filepath.Base(path))
			if err := os.WriteFile(tempFilePath, content, 0644); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", tempDir),
		dbName, driver)

	if err != nil {
		return err
	}

	// handle migration
	if err := m.Up(); err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

func setupHandlers(drivr driver.DriverKind, logger *zap.Logger, db *sql.DB) (driver.Driver, error) {

	logger.Info("starting_analytics")
	obs, err := analytics.NewPrometheus()

	if err != nil {
		return nil, err
	}

	levelRepo, err := model.NewSqlcLevelRepository(db, obs)

	if err != nil {
		return nil, err
	}

	playerRepo, err := model.NewSqlcPlayerRepository(db, obs)

	if err != nil {
		return nil, err
	}

	lc, err := service.NewLevelService(logger, obs, levelRepo, playerRepo)

	if err != nil {
		return nil, err
	}

	pc, err := service.NewPlayerService(logger, obs, levelRepo, playerRepo)

	if err != nil {
		return nil, err
	}

	var drvr driver.Driver

	switch drivr {
	case driver.Http:
		drvr, err = driver.NewWebDriver(logger, obs, lc, pc)
	case driver.Grpc:
		drvr, err = driver.NewGrpcDriver(logger, obs, lc, pc)
	case driver.GraphQL:
		drvr, err = driver.NewGraphQLDriver(logger, obs, lc, pc)
	default:
		return nil, &apperr.UnimplementedError{Message: "service"}
	}

	if err != nil {
		return nil, err
	}

	return drvr, nil
}
