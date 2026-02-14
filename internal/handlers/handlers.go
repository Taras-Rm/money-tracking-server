package handlers

import (
	"github.com/Taras-Rm/money-tracker-server/internal/services"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *services.Services
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		services,
	}
}

func (h *Handlers) InitHandlers() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger(), corsMiddleware)

	api := router.Group("/api")

	InjectAuthHandlers(api, h.services.Users)

	return router
}
