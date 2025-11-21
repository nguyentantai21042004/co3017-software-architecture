package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"learner-model-service/config"
	"learner-model-service/internal/consumer"
	"learner-model-service/internal/handler"
	"learner-model-service/internal/repository"
	"learner-model-service/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}

	db, err := connectDB(cfg.Postgres)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Printf("✅ Connected to database: %s:%d/%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DBName)

	masteryRepo := repository.NewMasteryRepository(db)
	learnerService := service.NewLearnerService(masteryRepo)
	learnerHandler := handler.NewLearnerHandler(learnerService)

	rabbitMQURL := cfg.RabbitMQ.URL
	if rabbitMQURL == "" {
		log.Fatal("❌ RABBITMQ_URL is not set")
	}

	eventConsumer, err := consumer.NewRabbitMQConsumer(rabbitMQURL, learnerService)
	if err != nil {
		log.Fatalf("❌ Failed to connect to RabbitMQ: %v", err)
	}
	defer eventConsumer.Close()

	err = eventConsumer.Start()
	if err != nil {
		log.Fatalf("❌ Failed to start consumer: %v", err)
	}

	if cfg.HTTPServer.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.GET("/health", learnerHandler.Health)

	internal := router.Group("/internal/learner")
	{
		internal.GET("/:user_id/mastery", learnerHandler.GetMastery)
	}

	go handleShutdown(eventConsumer, db)

	addr := fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Printf("🚀 Learner Model Service starting on %s", addr)

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
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func handleShutdown(eventConsumer consumer.EventConsumer, db *sql.DB) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("🛑 Shutting down gracefully...")
	if eventConsumer != nil {
		eventConsumer.Close()
	}
	if db != nil {
		db.Close()
	}
	log.Println("✅ Cleanup completed")
	os.Exit(0)
}
