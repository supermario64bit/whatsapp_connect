package main

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/whatsapp_connect/config"
)

func main() {
	config.LoadEnvFile()

	r := gin.Default()
	r.Run()
}
