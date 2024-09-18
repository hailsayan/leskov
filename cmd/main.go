package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/hailsayan/woland/cmd/api"
	configs "github.com/hailsayan/woland/internal/config"
	"github.com/hailsayan/woland/internal/db"
	"github.com/hailsayan/woland/internal/store"
	"go.uber.org/zap"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		DBName:               configs.Envs.DBName,
		Addr:                 configs.Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	initStorage(db, logger)
	store := store.NewStorage(db)

	server := api.NewServer(fmt.Sprintf(":%s", configs.Envs.Port), db, store, logger)
	if err := server.Run(); err != nil {
		logger.Fatal(err)
	}
}

func initStorage(db *sql.DB, logger *zap.SugaredLogger) {
	err := db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("DB: Successfully connected!")
}
