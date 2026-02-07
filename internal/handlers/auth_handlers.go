package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InjectAuthHandlers(gr *gin.RouterGroup, usersService interface{}) {
	handler := gr.Group("/auth")

	handler.POST("/register", registration(usersService))
}

func registration(usersService interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "registered"})
	}
}
