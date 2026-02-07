package app

import (
	"github.com/Taras-Rm/money-tracker-server/internal/setup"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	server := setup.NewServer("8080", router)

	server.Start()
}
