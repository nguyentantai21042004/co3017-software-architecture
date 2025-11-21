package main

import (
	"database/sql"
	"fmt"
	"log"
	"scoring-serviceinternal/handler"
	"scoring-serviceinternal/publisher"
	"scoring-serviceinternal/repository"
	"scoring-serviceinternal/service"
	"scoring/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title Scoring Service API
// @description ITS Scoring Service - Handles answer submissions and scoring
// @version 1.0
// @host localhost:8082
// @BasePath /api
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	db, err := connectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("✅ Connected to database: %s:%d/%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	// Initialize RabbitMQ Publisher
	rabbitMQURL := cfg.RabbitMQ.URL
	if rabbitMQURL == "" {
		log.Fatal("❌ RABBITMQ_URL is not set")
	}

	eventPublisher, err := publisher.NewRabbitMQPublisher(rabbitMQURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	defer eventPublisher.Close()

	// Initialize layers
	submissionRepo := repository.NewSubmissionRepository(db)
	scoringService := service.NewScoringService(submissionRepo, eventPublisher)
	scoringHandler := handler.NewScoringHandler(scoringService)

	// Setup Gin router
	if cfg.HTTPServer.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Health check
	router.GET("/health", scoringHandler.Health)

	// API routes
	api := router.Group("/api/scoring")
	{
		api.POST("/submit", scoringHandler.SubmitAnswer)
	}

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Printf("🚀 Scoring Service starting on %s", addr)
	log.Printf("📍 POST http://localhost:%d/api/scoring-servicesubmit", cfg.HTTPServer.Port)

	if err := router.Run(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
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

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
