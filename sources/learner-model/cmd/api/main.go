package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"learner-model-service/config"
	"learner-model-service/internal/consumer"
	learnerhttp "learner-model-service/internal/learner/delivery/http"
	learnerrepo "learner-model-service/internal/learner/repository/postgre"
	learnerusecase "learner-model-service/internal/learner/usecase"
	pkglog "learner-model-service/pkg/log"

	_ "learner-model-service/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title       Learner Model Service API
// @description Learner Model Service API documentation.
// @version     1
// @host        localhost:8083
// @schemes     http
// @BasePath    /internal/learner
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

	// Database connection
	db, err := connectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to connect to database | error=%v", err)
	}
	defer db.Close()

	log.Infof(ctx, "cmd.api.main: Connected to database | host=%s | port=%d | db=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	// Initialize repository
	learnerRepo := learnerrepo.New(db, log)

	// Initialize use case
	learnerUC := learnerusecase.New(log, learnerRepo)

	// Initialize HTTP handler
	learnerHandler := learnerhttp.New(log, learnerUC)

	// RabbitMQ Consumer
	rabbitMQURL := cfg.RabbitMQ.URL
	if rabbitMQURL == "" {
		log.Fatalf(ctx, "cmd.api.main: RABBITMQ_URL is not set")
	}

	eventConsumer, err := consumer.NewRabbitMQConsumer(rabbitMQURL, learnerUC, log)
	if err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to connect to RabbitMQ | error=%v", err)
	}
	defer eventConsumer.Close()

	err = eventConsumer.Start()
	if err != nil {
		log.Fatalf(ctx, "cmd.api.main: Failed to start consumer | error=%v", err)
	}

	if cfg.HTTPServer.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Swagger
	router.GET("/learner-model/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	internal := router.Group("/internal/learner")
	learnerhttp.MapLearnerRoutes(internal, learnerHandler)

	// Health endpoint
	internal.GET("/health", func(c *gin.Context) {
		learnerHandler.Health(c)
	})

	go handleShutdown(eventConsumer, db, log)

	addr := fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Infof(ctx, "cmd.api.main: Learner Model Service starting on %s", addr)

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

func handleShutdown(eventConsumer consumer.EventConsumer, db *sql.DB, log pkglog.Logger) {
	ctx := context.Background()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Infof(ctx, "cmd.api.main: Shutting down gracefully...")
	if eventConsumer != nil {
		eventConsumer.Close()
	}
	if db != nil {
		db.Close()
	}
	log.Infof(ctx, "cmd.api.main: Cleanup completed")
	os.Exit(0)
}
