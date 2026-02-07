package app

import (
	"github.com/Taras-Rm/money-tracker-server/internal/config"
	"github.com/Taras-Rm/money-tracker-server/internal/setup"
	"github.com/gin-gonic/gin"
)

func Run() {
	config := config.Config

	router := gin.Default()

	server := setup.NewServer(config.ServerConfig.Port, router)

	server.Start()
}
