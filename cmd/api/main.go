package main

import (
	"backend-app/config"
	"backend-app/internal/core/database"
	"backend-app/internal/core/middleware"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/master"
	"fmt"

	"github.com/gin-gonic/gin"
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

	// 4. Initialize Master Router using Wire
	masterRouter := master.InitializeMasterRouter(cfg, db)

	// 5. Setup Gin
	r := gin.New()
	r.Use(gin.Recovery())

	// 6. Global Middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(response.GlobalErrorHandler())

	// 7. Register Routes
	api := r.Group("/api/v1")
	masterRouter.RegisterRoutes(api)

	// 8. Start Server
	logrus.Infof("Starting %s on port %d...", cfg.App.Name, cfg.App.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		logrus.Fatalf("failed to start server: %v", err)
	}
}
