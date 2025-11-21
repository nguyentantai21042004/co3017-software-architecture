# Changelog - Scoring Service

All notable changes to the Scoring Service will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2025-11-22

### Added
- Initial implementation of Scoring Service with Clean Architecture
- Module-First architecture structure following Go best practices
- REST API for answer submission and scoring
- Integration with Content Service for question retrieval
- RabbitMQ publisher for submission events
- PostgreSQL database for submission history
- Answer validation and scoring logic
- Vietnamese feedback messages
- Health check endpoint
- Swagger API documentation

### Architecture
- **Structure**: Module-First + Clean Architecture
- **Layers**:
  - `internal/scoring/delivery/http/` - HTTP handlers and routing
  - `internal/scoring/usecase/` - Business logic for scoring submissions
  - `internal/scoring/repository/postgre/` - PostgreSQL data access
  - `internal/scoring/error.go` - Module-specific error definitions
  - `internal/scoring/type.go` - Input/Output types
  - `internal/scoring/interface.go` - Repository and UseCase interfaces
- **Database**: PostgreSQL with `submissions` table

### Technical Stack
- Go 1.25.4
- Gin Web Framework
- PostgreSQL database
- RabbitMQ message queue
- database/sql with pq driver
- HTTP client for service-to-service communication

### API Endpoints
- `POST /api/scoring/submit` - Submit answer for scoring
- `GET /api/scoring/health` - Health check

### Scoring Logic
1. Fetch question from Content Service by question_id
2. Compare submitted answer with correct answer
3. Calculate score (100 for correct, 0 for incorrect)
4. Store submission in database
5. Publish `SubmissionCompleted` event to RabbitMQ

### Event Publishing
- **Exchange**: `learner_exchange` (topic)
- **Routing Key**: `submission.completed`
- **Queue**: `learner.updates`
- **Event Type**: `SubmissionCompleted`
- **Payload**:
  ```json
  {
    "event": "SubmissionCompleted",
    "user_id": "string",
    "skill_tag": "string",
    "score_obtained": 0-100,
    "timestamp": "ISO8601"
  }
  ```

### Database Schema
```sql
CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    question_id INTEGER NOT NULL,
    answer VARCHAR(10) NOT NULL,
    is_correct BOOLEAN NOT NULL,
    score INTEGER NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

### Configuration
- `APP_PORT` - Service port (default: 8082)
- `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` - Database connection
- `POSTGRES_SSLMODE` - SSL mode
- `RABBITMQ_URL` - RabbitMQ connection string
- `CONTENT_SERVICE_URL` - Content Service URL
- `API_MODE` - Server mode (debug/release)
- Logger configuration (level, mode, encoding)

### Dependencies
- Content Service (for question retrieval)
- PostgreSQL database
- RabbitMQ message broker
- Learner Model Service (consumes published events)
