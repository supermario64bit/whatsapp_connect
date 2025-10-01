package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/supermario64bit/whatsapp_connect/db"
	"github.com/supermario64bit/whatsapp_connect/pkg/logger"
)

func Run() {
	createDBIfNotExists()
	db := db.New()

	migrationScriptsDir := "migrations/scripts/"

	files, err := filepath.Glob(filepath.Join(migrationScriptsDir, "*.sql"))
	logger.Info(strconv.Itoa(len(files)) + " migration scripts found.")
	if err != nil {
		logger.HighlightedDanger("Unable to load migration scripts. Error : " + err.Error())
	}

	for _, file := range files {
		logger.Info("Running migration: " + file)
		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			logger.HighlightedDanger("Unable to read migration script for the file " + file + ". Error : " + err.Error())
		}

		sqlScript := string(sqlBytes)
		_, err = db.Exec(sqlScript)

		if err != nil {
			logger.HighlightedDanger("Unable to run migration script for the file " + file + ". Error : " + err.Error())
		}

	}

	logger.Success("All migrations applied successfully!")

}

func createDBIfNotExists() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	masterDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", user, password, host, port)
	masterDB, err := sql.Open("postgres", masterDSN)
	if err != nil {
		logger.HighlightedDanger("Unable to connect postgres master DB. Error : " + err.Error())
		panic(err)
	}

	defer masterDB.Close()

	_, err = masterDB.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			logger.Info("Database already exists.")
		} else {
			logger.HighlightedDanger("Unable to create database. Error : " + err.Error())
			panic(err)
		}
	} else {
		logger.Success("Database created successfully!")
	}

	_, err = masterDB.Exec("GRANT ALL PRIVILEGES ON DATABASE " + dbName + " TO " + user + ";")
	if err != nil {
		logger.HighlightedDanger("Unable to create database. Error : " + err.Error())
		panic(err)
	}
}
