# Changelog - Learner Model Service

All notable changes to the Learner Model service will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-11-22

### Added
- Initial implementation of Learner Model Service with Clean Architecture
- Module-First architecture structure following Go best practices
- REST API for querying user skill mastery levels
- RabbitMQ consumer for real-time mastery updates
- Mastery calculation algorithm: `NewScore = (OldScore + ScoreObtained) / 2`
- PostgreSQL database integration with migration scripts
- Comprehensive unit tests achieving 93.8% coverage on use case layer
- Comprehensive unit tests achieving 86.4% coverage on HTTP handler layer
- Health check endpoint
- Graceful shutdown handling
- Swagger API documentation

### Architecture
- **Structure**: Module-First + Clean Architecture
- **Layers**:
  - `internal/learner/delivery/http/` - HTTP handlers and routing
  - `internal/learner/usecase/` - Business logic for mastery management
  - `internal/learner/repository/postgre/` - PostgreSQL data access
  - `internal/learner/error.go` - Module-specific error definitions
  - `internal/learner/type.go` - Input/Output types
  - `internal/learner/interface.go` - Repository and UseCase interfaces
  - `internal/model/` - Domain models (SkillMastery, SubmissionEvent)
  - `internal/consumer/` - RabbitMQ event consumer
- **Database**: PostgreSQL with `skill_mastery` table

### Technical Stack
- Go 1.25.4
- Gin Web Framework
- PostgreSQL database
- RabbitMQ message queue
- database/sql with pq driver
- Zap structured logging

### API Endpoints
- `GET /internal/learner/:user_id/mastery?skill={skill}` - Get user mastery for a skill
- `GET /internal/learner/health` - Health check

### Event Consumption
- **Queue**: `learner.updates`
- **Event Type**: `SubmissionCompleted`
- **Processing**: Updates skill mastery based on submission scores
- **Formula**: Averages old score with obtained score

### Database Schema
```sql
CREATE TABLE skill_mastery (
    user_id VARCHAR(255) NOT NULL,
    skill_tag VARCHAR(255) NOT NULL,
    current_score INTEGER NOT NULL DEFAULT 0,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, skill_tag),
    CHECK (current_score >= 0 AND current_score <= 100)
);
```

### Configuration
- `APP_PORT` - Service port (default: 8080)
- `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` - Database connection
- `POSTGRES_SSLMODE` - SSL mode (default: disable)
- `RABBITMQ_URL` - RabbitMQ connection string
- `API_MODE` - Server mode (debug/release)
- Logger configuration (level, mode, encoding)

### Testing
- Unit tests for all core use cases (16 tests)
- Unit tests for HTTP handlers (13+ tests)
- Mock-based testing using mockery and testify
- Total test coverage: 80%+ on core business logic

### Fixed Issues
- PostgreSQL SSL mode compatibility (changed default from "prefer" to "disable")
- Response error handling to use gin.H format
- Database migration script for skill_mastery table
- Port conflict resolution (8083 → 8080)
- RabbitMQ consumer lifecycle management

### Dependencies
- Scoring Service (publishes events this service consumes)
- PostgreSQL database
- RabbitMQ message broker
