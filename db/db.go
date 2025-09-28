package db

import (
	"database/sql"
	"fmt"
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
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err = sql.Open("postgres", dsn)

	if err != nil {
		logger.HighlightedDanger("Unable to connect to db. Error: " + err.Error())
		panic(err)
	}

	return db
}
