package main

import (
	"fmt"
	"log"

	"adaptive-engine-service/internal/handler"
	"adaptive-engine-service/internal/service"

	_ "adaptive-engine-service/docs"

	"github.com/caarlos0/env/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Config struct {
	Port              int    `env:"APP_PORT" envDefault:"8084"`
	Mode              string `env:"API_MODE" envDefault:"debug"`
	LearnerServiceURL string `env:"LEARNER_SERVICE_URL" envDefault:"http://localhost:8083"`
	ContentServiceURL string `env:"CONTENT_SERVICE_URL" envDefault:"http://localhost:8081"`
}

func main() {
	_ = godotenv.Load()
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	log.Printf("✅ Config loaded: Learner=%s, Content=%s", cfg.LearnerServiceURL, cfg.ContentServiceURL)

	adaptiveService := service.NewAdaptiveService(cfg.LearnerServiceURL, cfg.ContentServiceURL)
	adaptiveHandler := handler.NewAdaptiveHandler(adaptiveService)

	if cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/health", adaptiveHandler.Health)

	// Swagger
	router.GET("/adaptive-engine/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/adaptive")
	{
		api.POST("/next-lesson", adaptiveHandler.NextLesson)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("🚀 Adaptive Engine starting on %s", addr)
	log.Printf("📍 POST http://localhost:%d/api/adaptive/next-lesson", cfg.Port)

	if err := router.Run(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
