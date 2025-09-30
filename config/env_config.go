package config

import (
	"github.com/joho/godotenv"
	"github.com/supermario64bit/whatsapp_connect/pkg/logger"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.HighlightedDanger("Error loading .env file. Error: " + err.Error())
		panic(err)
	}
}
