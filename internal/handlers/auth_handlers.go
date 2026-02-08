package handlers

import (
	"net/http"

	"github.com/Taras-Rm/money-tracker-server/internal/dto"
	"github.com/Taras-Rm/money-tracker-server/internal/services"
	"github.com/Taras-Rm/money-tracker-server/internal/services/models"
	"github.com/gin-gonic/gin"
)

func InjectAuthHandlers(gr *gin.RouterGroup, usersService services.Users) {
	handler := gr.Group("/auth")

	handler.POST("/register", registration(usersService))
}

func registration(usersService services.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegistrationRequest

		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		user, err := usersService.CreateUser(c, models.CreateUserInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}
