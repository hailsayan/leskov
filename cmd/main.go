package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/hailsayan/woland/cmd/api"
	configs "github.com/hailsayan/woland/config"
	"github.com/hailsayan/woland/db"
	"github.com/hailsayan/woland/store"
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

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initStorage(db)
	store := store.NewStorage(db)

	server := api.NewServer(fmt.Sprintf(":%s", configs.Envs.Port), db, store)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
