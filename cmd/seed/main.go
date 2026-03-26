package main

import (
	"fmt"
	"hospital-api/config"
	"hospital-api/internal/database"
	"hospital-api/internal/seed"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	if err := seed.SeedHospital(db); err != nil {
		log.Fatalf("Failed to seed: %v", err)
	}

	fmt.Println("Seed hospital")
}
