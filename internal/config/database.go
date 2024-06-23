package config

import (
	"database/sql"
	"log/slog"

	"github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {

	slog.Info("starting database configuration")

	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "mydb",
	}

	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		slog.Error("could not open sql database", "err", err.Error())
		panic(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		slog.Error("could not ping database", "err", pingErr.Error())
		panic(pingErr)
	}

	slog.Info("database connection established successfully")

	return db
}
