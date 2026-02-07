package setup

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	router *gin.Engine
}

func NewServer(port string, router *gin.Engine) *Server {
	return &Server{
		port:   port,
		router: router,
	}
}

func (s *Server) Start() {
	s.router.Run(":" + s.port)

	fmt.Printf("Server working on port: %s", s.port)
}
