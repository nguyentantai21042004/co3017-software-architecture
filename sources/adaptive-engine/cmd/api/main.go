package main

import (
	"context"
	"fmt"

	"adaptive-engine/config"
	adaptivehttp "adaptive-engine/internal/adaptive/delivery/http"
	adaptiveusecase "adaptive-engine/internal/adaptive/usecase"
	"adaptive-engine/pkg/curl"
	pkglog "adaptive-engine/pkg/log"

	_ "adaptive-engine/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Adaptive Engine API
// @description Adaptive Engine API documentation.
// @version     1
// @host        localhost:8084
// @schemes     http
// @BasePath    /api/adaptive
func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Logger
	log := pkglog.Init(pkglog.ZapConfig{
		Level:    cfg.LoggerLevel,
		Mode:     cfg.LoggerMode,
		Encoding: cfg.LoggerEncoding,
	})

	// HTTP clients for external services
	lrCl := curl.NewLearnerServiceClient(cfg.LearnerServiceURL)
	ctCl := curl.NewContentServiceClient(cfg.ContentServiceURL)

	log.Infof(ctx, "cmd.api.main: Config loaded successfully")

	// Use case
	uc := adaptiveusecase.New(log, lrCl, ctCl)

	// HTTP handler
	h := adaptivehttp.New(log, uc)

	if cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger
	r.GET("/adaptive-engine/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api/adaptive")
	adaptivehttp.MapAdaptiveRoutes(api, h)

	// Health endpoint
	api.GET("/health", func(c *gin.Context) {
		h.Health(c)
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Infof(ctx, "cmd.api.main: Adaptive Engine starting on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to start server: %v", err)
	}
}
