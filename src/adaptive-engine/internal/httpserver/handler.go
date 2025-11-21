package httpserver

import (
	"adaptive-engine/internal/middleware"
	projecthttp "adaptive-engine/internal/project/delivery/http"
	projectrepository "adaptive-engine/internal/project/repository/postgre"
	projectusecase "adaptive-engine/internal/project/usecase"
	"adaptive-engine/pkg/i18n"
	"adaptive-engine/pkg/scope"

	// Import this to execute the init function in docs.go which setups the Swagger docs.
	// Uncomment after running: make swagger
	_ "adaptive-engine/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const apiPrefix = "/project"

func (srv HTTPServer) mapHandlers() error {
	srv.registerMiddlewares()
	srv.registerSystemRoutes()

	scopeManager := scope.New(srv.jwtSecretKey)
	mw := middleware.New(srv.l, scopeManager)

	i18n.Init()

	// Initialize project repository
	projectRepo := projectrepository.New(srv.postgresDB, srv.l)

	// Initialize project usecase
	projectUC := projectusecase.New(srv.l, projectRepo)

	// Initialize project HTTP handler
	projectHandler := projecthttp.New(srv.l, projectUC)

	// Map routes
	api := srv.gin.Group(apiPrefix)
	projecthttp.MapProjectRoutes(api.Group("/projects"), projectHandler, mw)

	return nil
}

func (srv HTTPServer) registerMiddlewares() {
	srv.gin.Use(middleware.Recovery(srv.l, srv.discord))

	corsConfig := middleware.DefaultCORSConfig()
	srv.gin.Use(middleware.CORS(corsConfig))
}

func (srv HTTPServer) registerSystemRoutes() {
	api := srv.gin.Group(apiPrefix)
	api.GET("/health", srv.healthCheck)
	api.GET("/ready", srv.readyCheck)
	api.GET("/live", srv.liveCheck)

	srv.gin.GET("/project/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/project/swagger/doc.json"),
	))
}
