package main

import (
	"context"
	"database/sql"
	"fmt"

	"scoring/config"
	"scoring/internal/publisher"
	scoringhttp "scoring/internal/scoring/delivery/http"
	scoringrepo "scoring/internal/scoring/repository/postgre"
	scoringusecase "scoring/internal/scoring/usecase"
	"scoring/pkg/curl"
	pkglog "scoring/pkg/log"

	_ "scoring/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Scoring Service API
// @description Scoring Service API documentation.
// @version     1
// @host        localhost:8082
// @schemes     http
// @BasePath    /api/scoring
func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Logger
	log := pkglog.Init(pkglog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	// Database
	db, err := connectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to connect to database | error=%v", err)
	}
	defer db.Close()

	log.Infof(ctx, "cmd.api.main: Connected to database | host=%s | port=%d | db=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	// Initialize RabbitMQ Publisher
	eventPublisher, err := publisher.NewRabbitMQPublisher(cfg.RabbitMQ.URL, log)
	if err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to connect to RabbitMQ | error=%v", err)
	}
	defer eventPublisher.Close()

	log.Infof(ctx, "cmd.api.main: Connected to RabbitMQ | url=%s", cfg.RabbitMQ.URL)

	// HTTP Client for Content Service
	contentClient := curl.NewContentServiceClient(cfg.ContentServiceURL)

	// Initialize layers (Module-First approach)
	submissionRepo := scoringrepo.New(db, log)
	scoringUC := scoringusecase.New(log, submissionRepo, eventPublisher, contentClient)
	scoringHandler := scoringhttp.New(log, scoringUC)

	if cfg.HTTPServer.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS Middleware
	router.Use(func(c *gin.Context) {
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

	// Health check (global)
	router.GET("/health", scoringHandler.Health)

	// Swagger
	router.GET("/scoring/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := router.Group("/api/scoring")
	scoringhttp.MapScoringRoutes(api, scoringHandler)

	addr := fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Infof(ctx, "cmd.api.main: Scoring Service starting on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to start server | error=%v", err)
	}
}

func connectDB(cfg config.PostgresConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
