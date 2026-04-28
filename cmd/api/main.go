package main

import (
	"backend-app/config"
	"backend-app/internal/core/cache"
	"backend-app/internal/core/database"
	"backend-app/internal/core/middleware"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/auth"
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

	// 3.1 Initialize Redis (Optional)
	rdb, err := database.NewRedisClient(cfg)
	if err != nil {
		logrus.Warnf("failed to initialize redis (using in-memory cache instead): %v", err)
	} else if rdb != nil {
		cache.GetCache().SetRedisClient(rdb)
	}

	// 4. Initialize Master Router using Wire
	masterRouter := master.InitializeMasterRouter(cfg, db)
	authRouter := auth.InitializeAuthRouter(cfg, db)

	// 5. Setup Gin
	r := gin.New()
	r.Use(gin.Recovery())

	// 6. Global Middleware
	r.Use(middleware.ResponseIDMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(response.GlobalErrorHandler())

	// 6.1 Handle 404 & 405
	r.NoRoute(func(c *gin.Context) {
		response.SendError(c, response.PathNotFound, "The requested path was not found on this server")
	})
	r.NoMethod(func(c *gin.Context) {
		response.SendError(c, response.MethodNotAllowed, "The requested method is not allowed for this path")
	})

	// 7. Register Routes
	api := r.Group("/api/v1")
	masterRouter.RegisterRoutes(api)
	authRouter.RegisterRoutes(api)

	// 8. Start Server
	logrus.Infof("Starting %s on port %d...", cfg.App.Name, cfg.App.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		logrus.Fatalf("failed to start server: %v", err)
	}
}
