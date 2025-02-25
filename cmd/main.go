package main

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-sql-driver/mysql"
	"github.com/hailsayan/leskov/cmd/api"
	configs "github.com/hailsayan/leskov/internal/config"
	"github.com/hailsayan/leskov/internal/db"
	"github.com/hailsayan/leskov/internal/store"
	"github.com/hailsayan/leskov/internal/store/cache"
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

	var rdb *redis.Client
	if configs.Envs.Enabled {
		rdb = cache.NewRedisClient(configs.Envs.Addr, configs.Envs.Pw, int(configs.Envs.Db))
		logger.Info("redis cache connection established")

		defer rdb.Close()
	}

	initStorage(db, logger)
	store := store.NewStorage(db)
	cacheStorage := cache.NewRedisStorage(rdb)

	server := api.NewServer(fmt.Sprintf(":%s", configs.Envs.Port), db, store, logger, cacheStorage)
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
