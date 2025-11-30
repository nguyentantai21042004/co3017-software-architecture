package main

import (
	"context"
	"database/sql"
	"fmt"

	"learner-model-service/config"
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
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
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

	if cfg.HTTPServer.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// CORS Middleware
	router.Use(CORSMiddleware())

	// Health check (global)
	router.GET("/health", learnerHandler.Health)

	// Swagger
	router.GET("/learner-model/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	internal := router.Group("/internal/learner")
	learnerhttp.MapLearnerRoutes(internal, learnerHandler)

	addr := fmt.Sprintf("%s:%d", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	log.Infof(ctx, "cmd.api.main: Learner Model API Service starting on %s", addr)

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

// CORSMiddleware handles CORS configuration
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}
