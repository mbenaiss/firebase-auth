package api

import (
	"fmt"

	"github.com/auth/api/middleware"
	"github.com/auth/service"
	"github.com/gin-gonic/gin"
)

// Server is the API server.
type Server struct {
	Port string
	svc  *service.Service
}

// NewServer creates a new API server.
func NewServer(port string, svc *service.Service) *Server {
	return &Server{
		Port: port,
	}
}

// Run runs the API server.
func (s *Server) Run() error {
	server := gin.Default()
	server.Use(middleware.Cors())

	server.POST("/signup", signup(s.svc))
	server.POST("/signin", signin(s.svc))
	server.POST("/password/reset", resetPassword(s.svc))
	server.POST("/token/validate", validateToken(s.svc))

	err := server.Run(s.Port)
	if err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	return nil
}
