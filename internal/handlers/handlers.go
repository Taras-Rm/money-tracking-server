package handlers

import (
	"github.com/Taras-Rm/money-tracker-server/internal/services"
	"github.com/Taras-Rm/money-tracker-server/pkg/token"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services     *services.Services
	tokenManager *token.TokenManager
}

func NewHandlers(services *services.Services, tokenManager *token.TokenManager) *Handlers {
	return &Handlers{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handlers) InitHandlers() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger(), corsMiddleware)

	api := router.Group("/api")

	InjectAuthHandlers(api, h.services.Users, h.tokenManager)

	return router
}
