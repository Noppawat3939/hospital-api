package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB) *Server {
	r := gin.New()

	r.Use(CORS())

	return nil
}
