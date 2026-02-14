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
	handler.POST("/login", login(usersService))
	handler.POST("/login-with-google", loginWithGoogle(usersService))
}

func registration(usersService services.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.RegistrationRequest

		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		token, err := usersService.RegisterUser(c, models.CreateUserInput{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": token})
	}
}

func login(usersService services.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest

		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		token, err := usersService.LoginUser(c, models.LoginUserInput{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func loginWithGoogle(usersService services.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginWithGoogleRequest

		err := c.BindJSON(&req)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
			return
		}

		if req.AccessToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "access_token is required"})
			return
		}

		token, err := usersService.LoginUserWithGoogle(c, models.LoginUserWithGoogleInput{
			AccessToken: req.AccessToken,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
