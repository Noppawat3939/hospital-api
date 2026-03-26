package main

import (
	"fmt"
	"hospital-api/config"
	"hospital-api/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	_, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	fmt.Println("Server started")
	r.Run(":" + cfg.Port)
}
