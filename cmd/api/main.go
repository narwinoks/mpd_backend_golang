package main

import (
	"backend-app/config"
	"backend-app/internal/core/cache"
	"backend-app/internal/core/database"
	"backend-app/internal/core/middleware"
	"backend-app/internal/core/response"
	"backend-app/internal/modules/auth"
	"backend-app/internal/modules/master"
	"backend-app/pkg/validations"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// Setup Logging
	config.SetupLog(&cfg.Log)

	// Initialize Database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	// Initialize Redis (Optional)
	rdb, err := database.NewRedisClient(cfg)
	if err != nil {
		logrus.Warnf("failed to initialize redis (using in-memory cache instead): %v", err)
	} else if rdb != nil {
		cache.GetCache().SetRedisClient(rdb)
	}

	// Initialize Routers using Wire
	authRouter := auth.InitializeAuthRouter(cfg, db)
	masterRouter := master.InitializeMasterRouter(cfg, db)

	//  Setup Gin
	r := gin.New()
	r.Use(gin.Recovery())

	//  Global Middleware
	r.Use(middleware.CORSMiddleware(&cfg.CORS))
	r.Use(middleware.ResponseIDMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(response.GlobalErrorHandler())

	// Handle 404 & 405
	r.NoRoute(func(c *gin.Context) {
		response.SendError(c, response.PathNotFound, "The requested path was not found on this server")
	})
	r.NoMethod(func(c *gin.Context) {
		response.SendError(c, response.MethodNotAllowed, "The requested method is not allowed for this path")
	})

	// Register Routes
	api := r.Group("/api/v1")

	// Auth routes (some are public like /login)
	authRouter.RegisterRoutes(api)

	// Protected routes (Master module)
	protectedGroup := api.Group("")
	protectedGroup.Use(authRouter.AuthMiddleware.Handle())
	{
		masterRouter.RegisterRoutes(protectedGroup)
	}
	//custom validation handler
	validations.InitGinValidator(db)
	// Start Server
	logrus.Infof("Starting %s on port %d...", cfg.App.Name, cfg.App.Port)
	if err := r.Run(fmt.Sprintf(":%d", cfg.App.Port)); err != nil {
		logrus.Fatalf("failed to start server: %v", err)
	}
}
