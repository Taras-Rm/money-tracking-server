package handlers

import (
	"net/http"

	"github.com/Taras-Rm/money-tracker-server/internal/dto"
	"github.com/Taras-Rm/money-tracker-server/internal/services"
	"github.com/Taras-Rm/money-tracker-server/internal/services/models"
	"github.com/Taras-Rm/money-tracker-server/pkg/token"
	"github.com/gin-gonic/gin"
)

func InjectAuthHandlers(gr *gin.RouterGroup, usersService services.Users, tokenManager *token.TokenManager) {
	handler := gr.Group("/auth")

	handler.POST("/register", registration(usersService))
	handler.POST("/login", login(usersService))
	handler.POST("/login-with-google", loginWithGoogle(usersService))

	handler.GET("/me", authMiddleware(tokenManager), me(usersService))
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

func me(usersService services.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by authMiddleware)
		userId, exists := c.Get(UserIDKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "user ID not found in context"})
			return
		}

		userIdInt64, ok := userId.(int64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "invalid user ID type"})
			return
		}

		user, err := usersService.GetUserByID(c, userIdInt64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
