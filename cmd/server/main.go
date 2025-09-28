package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/whatsapp_connect/config"
	"github.com/supermario64bit/whatsapp_connect/migrations"
)

func main() {
	config.LoadEnvFile()
	migrations.Run()

	r := gin.Default()
	r.Run()
}
