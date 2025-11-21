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
	learnerrepo "learner-model-service/internal/learner/repository/postgre"
	learnerusecase "learner-model-service/internal/learner/usecase"
	pkglog "learner-model-service/pkg/log"

	_ "github.com/lib/pq"
)

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

	log.Infof(ctx, "cmd.consumer.main: Starting Learner Model Consumer Service")

	// Database connection
	db, err := connectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf(ctx, "cmd.consumer.main: Failed to connect to database | error=%v", err)
	}
	defer db.Close()

	log.Infof(ctx, "cmd.consumer.main: Connected to database | host=%s | port=%d | db=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	// Initialize repository
	learnerRepo := learnerrepo.New(db, log)

	// Initialize use case
	learnerUC := learnerusecase.New(log, learnerRepo)

	// RabbitMQ Consumer
	rabbitMQURL := cfg.RabbitMQ.URL
	if rabbitMQURL == "" {
		log.Fatalf(ctx, "cmd.consumer.main: RABBITMQ_URL is not set")
	}

	eventConsumer, err := consumer.NewRabbitMQConsumer(rabbitMQURL, learnerUC, log)
	if err != nil {
		log.Fatalf(ctx, "cmd.consumer.main: Failed to connect to RabbitMQ | error=%v", err)
	}
	defer eventConsumer.Close()

	log.Infof(ctx, "cmd.consumer.main: Connected to RabbitMQ | url=%s", rabbitMQURL)

	// Start consumer
	err = eventConsumer.Start()
	if err != nil {
		log.Fatalf(ctx, "cmd.consumer.main: Failed to start consumer | error=%v", err)
	}

	log.Infof(ctx, "cmd.consumer.main: Consumer started successfully, waiting for messages...")

	// Wait for interrupt signal to gracefully shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Infof(ctx, "cmd.consumer.main: Shutting down consumer...")

	if err := eventConsumer.Close(); err != nil {
		log.Errorf(ctx, "cmd.consumer.main: Error closing consumer | error=%v", err)
	}

	if err := db.Close(); err != nil {
		log.Errorf(ctx, "cmd.consumer.main: Error closing database | error=%v", err)
	}

	log.Infof(ctx, "cmd.consumer.main: Consumer shut down successfully")
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
