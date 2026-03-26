package server

import (
	"hospital-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	r := s.router
	// db := s.db

	r.NoRoute(func(c *gin.Context) {
		response.Error(c, http.StatusNotFound, nil)
		c.Abort()
	})

}
