package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/supermario64bit/whatsapp_connect/pkg/logger"
)

var db *sql.DB

func New() *sql.DB {
	if db != nil {
		return db
	}
	var err error
	dsn := os.Getenv("DB_DSN")
	db, err = sql.Open("postgres", dsn)

	if err != nil {
		logger.HighlightedDanger("Unable to connect to db. Error: " + err.Error())
		panic(err)
	}

	return db
}
