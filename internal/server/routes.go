package server

import (
	"fmt"
	"hospital-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerRoutes() {
	r := s.router

	r.NoRoute(func(c *gin.Context) {
		msg := fmt.Sprintf("path %s not found", c.Request.URL)
		response.Error(c, http.StatusNotFound, msg)
	})

	RegisterStaffRoutes(s.router, s.db)
}
