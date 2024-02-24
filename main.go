package main

import (
	"database/sql"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/controller"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

// environment variables
const (
	DbHostEnvVar = "DB_HOST"
	DbUserEnvVar = "DB_USER"
	DbPassEnvVar = "DB_PASSWORD"
	DbNameEnvVar = "DB_NAME"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	db, err := setupDatabase(logger)

	if err != nil {
		logger.Fatal("setup_db", zap.Error(err))
	}

	if err := setupHandlers(logger, db); err != nil {
		logger.Fatal("setup_server", zap.Error(err))
	}

	logger.Debug("listening")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("listen_and_serve", zap.Error(err))
	}
}

func setupDatabase(logger *zap.Logger) (*gorm.DB, error) {
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

	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		logger.Error("gorm_open", zap.Error(err))
		return nil, err
	}

	if err := gormDb.AutoMigrate(&model.Level{}, &model.Player{}); err != nil {
		return nil, err
	}

	return gormDb, nil
}

func setupHandlers(logger *zap.Logger, db *gorm.DB) error {

	lc, err := controller.NewLevelController(logger, db)

	if err != nil {
		return err
	}

	pc, err := controller.NewPlayerController(logger, db)

	if err != nil {
		return err
	}

	http.HandleFunc("/level/submit", lc.SubmitLevel)
	http.HandleFunc("/player/move", pc.MovePlayer)

	return nil
}
