package main

import (
	"hospital-api/config"
	"hospital-api/internal/database"
	"hospital-api/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(db)

	if err := srv.Start(cfg.Port); err != nil {
		log.Fatal(err)
	}
}
