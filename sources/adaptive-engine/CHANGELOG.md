# Changelog - Adaptive Engine Service

All notable changes to the Adaptive Engine service will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-11-22

### Added
- Initial implementation of Adaptive Engine with Clean Architecture
- Module-First architecture structure following Go best practices
- REST API endpoint for personalized learning recommendations
- Integration with Learner Model Service for mastery score retrieval
- Integration with Content Service for content recommendations
- Adaptive logic: Remedial content for mastery < 50%, standard content otherwise
- Comprehensive logging for debugging and monitoring
- Health check endpoint
- Swagger API documentation

### Architecture
- **Structure**: Module-First + Clean Architecture
- **Layers**:
  - `internal/adaptive/delivery/http/` - HTTP handlers and routing
  - `internal/adaptive/usecase/` - Business logic and orchestration
  - `internal/adaptive/error.go` - Module-specific error definitions
  - `internal/adaptive/type.go` - Input/Output types
  - `internal/adaptive/interface.go` - UseCase interface
- **External Dependencies**:
  - Learner Model Service (port 8083)
  - Content Service (port 8081)
  - HTTP client utilities (`pkg/curl`)

### Technical Stack
- Go 1.25.4
- Gin Web Framework
- RESTful API design
- HTTP client for service-to-service communication

### API Endpoints
- `POST /api/adaptive/next-lesson` - Get personalized lesson recommendation

### Configuration
- `APP_PORT` - Service port (default: 8084)
- `LEARNER_SERVICE_URL` - Learner Model Service URL
- `CONTENT_SERVICE_URL` - Content Service URL
- `API_MODE` - Server mode (debug/release)
- Logger configuration (level, mode, encoding)

### Notes
- Service is stateless (no database required)
- Implements adaptive learning logic based on user mastery scores
- Acts as orchestrator between Learner Model and Content services
