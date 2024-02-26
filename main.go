package main

import (
	"database/sql"
	"fmt"
	"github.com/JakobDFrank/penn-roguelike/internal/driver"
	"github.com/JakobDFrank/penn-roguelike/internal/model"
	"github.com/JakobDFrank/penn-roguelike/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	_ "google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

	driver, err := setupHandlers(logger, db)
	if err != nil {
		logger.Fatal("setup_server", zap.Error(err))
	}

	if err := driver.Serve(); err != nil {
		logger.Fatal("serve", zap.Error(err))
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

// JF - note: we could create interfaces here for zap.Logger and gorm.DB to abide by dependency inversion
// however, it will increase complexity. trade-offs.
func setupHandlers(logger *zap.Logger, db *gorm.DB) (driver.Driver, error) {
	lc, err := service.NewLevelController(logger, db)

	if err != nil {
		return nil, err
	}

	pc, err := service.NewPlayerController(logger, db)

	if err != nil {
		return nil, err
	}

	wd, err := driver.NewWebDriver(lc, pc, logger)

	if err != nil {
		return nil, err
	}

	//gd, err := driver.NewGrpcDriver(lc, pc, logger)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gd, nil
	return wd, nil
}
