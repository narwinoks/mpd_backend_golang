package main

import (
	"backend-app/config"
	"backend-app/db/seeders"
	"backend-app/internal/core/database"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// 2. Setup Logging
	config.SetupLog(&cfg.Log)

	// 3. Initialize Database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	// 4. Run Seeders
	if err := seeders.SeedAll(db); err != nil {
		logrus.Fatalf("failed to run seeders: %v", err)
	}

	logrus.Info("Seeders completed successfully")
}
