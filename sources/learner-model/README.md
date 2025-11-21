# Learner Model Service

**Port:** 8080 (API), Consumer (Background)
**Database:** PostgreSQL (learner_db)
**Technology:** Go 1.25.4, Gin, PostgreSQL, RabbitMQ
**Architecture:** Module-First + Clean Architecture

## Overview

Learner Model Service tracks and manages user skill mastery levels. It consists of two components:

1. **REST API** - Query mastery scores for adaptive learning
2. **Event Consumer** - Process submission events to update mastery in real-time

### Key Responsibilities
- Store and retrieve user skill mastery scores
- Process submission events from RabbitMQ
- Calculate updated mastery using adaptive algorithm
- Provide mastery data to Adaptive Engine

### Mastery Update Algorithm
```
NewMastery = (CurrentMastery + ScoreObtained) / 2
```

Example:
- Current: 60, Score: 0 (wrong) → New: 30
- Current: 30, Score: 100 (correct) → New: 65

## Architecture

### Module-First + Clean Architecture

```
learner-model/
├── cmd/api/main.go                    # Entry point
├── internal/
│   ├── learner/                       # Learner module
│   │   ├── delivery/http/             # HTTP handlers
│   │   ├── usecase/                   # Business logic
│   │   ├── repository/postgre/        # Data access
│   │   ├── error.go
│   │   ├── type.go
│   │   └── interface.go
│   ├── model/                         # Domain models
│   │   └── skill_mastery.go
│   └── consumer/                      # RabbitMQ consumer
│       └── rabbitmq_consumer.go
├── migration/                         # Database migrations
└── pkg/                               # Shared utilities
```

### Test Coverage
- **Use Case Tests**: 16 tests, 93.8% coverage
- **HTTP Handler Tests**: 13+ tests, 86.4% coverage
- **Total**: 80%+ coverage on core business logic

## Setup

### Prerequisites
- Go 1.25.4+
- PostgreSQL 13+
- RabbitMQ 3.8+

### Database Setup

```bash
# Run migration script
psql -U postgres -h localhost -p 5432 -f migration/01_create_skill_mastery_table.sql

# Verify table
psql -U postgres -d postgres -c "\d skill_mastery"
```

### Installation

```bash
cd sources/learner-model

# Install dependencies
go mod tidy

# Set environment variables
export APP_PORT=8080
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_DB=postgres
export POSTGRES_SSLMODE=disable
export RABBITMQ_URL=amqp://admintest:adminTest2025@localhost:5672/
export LOGGER_LEVEL=debug

# Run service
go run cmd/api/main.go
```

Service starts on **http://localhost:8080**

## API Reference

### GET /internal/learner/:user_id/mastery

Get user's mastery score for a specific skill.

**Request:**
```bash
curl "http://localhost:8080/internal/learner/user123/mastery?skill=math_algebra"
```

**Response (Success):**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "user_id": "user123",
    "skill_tag": "math_algebra",
    "current_score": 75,
    "last_updated": "2025-11-22T10:30:00Z"
  }
}
```

**Response (Not Found):**
```json
{
  "error_code": 404,
  "message": "Mastery record not found",
  "data": null
}
```

### GET /internal/learner/health

Health check endpoint.

```bash
curl http://localhost:8080/internal/learner/health
```

## Event Consumer

### RabbitMQ Configuration
- **Queue**: `learner.updates`
- **Exchange**: `learner_exchange` (topic)
- **Routing Key**: `submission.completed`

### Event Schema

```json
{
  "event": "SubmissionCompleted",
  "user_id": "user123",
  "skill_tag": "math_algebra",
  "score_obtained": 0,
  "timestamp": "2025-11-22T10:30:00Z"
}
```

### Processing Flow

```
1. Receive event from RabbitMQ
   ↓
2. Fetch current mastery from database
   ↓
3. Calculate new mastery: (old + obtained) / 2
   ↓
4. Update database with new score
   ↓
5. Log success and acknowledge message
```

## Testing

### Unit Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./internal/learner/...

# Generate coverage report
go test -coverprofile=coverage.out ./internal/learner/...
go tool cover -html=coverage.out
```

### Integration Testing

**Test 1: Check Initial Mastery**
```bash
curl "http://localhost:8080/internal/learner/user123/mastery?skill=math_algebra"
# Expected: mastery_score varies based on history
```

**Test 2: Trigger Update via Scoring Service**
```bash
# Submit wrong answer (score=0)
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "question_id": 1,
    "answer": "C"
  }'

# Wait 1-2 seconds for event processing

# Verify mastery decreased
curl "http://localhost:8080/internal/learner/user123/mastery?skill=math_algebra"
```

**Test 3: Submit Correct Answer**
```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "question_id": 1,
    "answer": "A"
  }'

# Verify mastery increased
curl "http://localhost:8080/internal/learner/user123/mastery?skill=math_algebra"
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_PORT` | API server port | `8080` | No |
| `POSTGRES_HOST` | Database host | `localhost` | Yes |
| `POSTGRES_PORT` | Database port | `5432` | Yes |
| `POSTGRES_USER` | Database user | `postgres` | Yes |
| `POSTGRES_PASSWORD` | Database password | - | Yes |
| `POSTGRES_DB` | Database name | `postgres` | Yes |
| `POSTGRES_SSLMODE` | SSL mode | `disable` | No |
| `RABBITMQ_URL` | RabbitMQ connection URL | - | Yes |
| `API_MODE` | Server mode | `debug` | No |
| `LOGGER_LEVEL` | Log level | `debug` | No |

### Sample .env File

```env
APP_PORT=8080
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_SSLMODE=disable
RABBITMQ_URL=amqp://admintest:adminTest2025@localhost:5672/
API_MODE=debug
LOGGER_LEVEL=debug
```

## Database Schema

```sql
CREATE TABLE IF NOT EXISTS skill_mastery (
    user_id VARCHAR(255) NOT NULL,
    skill_tag VARCHAR(255) NOT NULL,
    current_score INTEGER NOT NULL DEFAULT 0 CHECK (current_score >= 0 AND current_score <= 100),
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, skill_tag)
);

CREATE INDEX idx_user_skill ON skill_mastery(user_id, skill_tag);
```

## Logging

The service provides detailed structured logs:

**API Requests:**
```
2025-11-22T10:30:00.000+0700 INFO  learner.usecase.GetMastery: starting | user_id=user123 | skill_tag=math_algebra
2025-11-22T10:30:00.050+0700 INFO  learner.usecase.GetMastery: success | current_score=75
```

**Event Processing:**
```
2025-11-22T10:30:00.000+0700 INFO  consumer.Start: Listening for events on queue: learner.updates
2025-11-22T10:30:10.123+0700 INFO  consumer.handleMessage: Received submission event | user_id=user123 | skill=math_algebra | score=0
2025-11-22T10:30:10.180+0700 INFO  learner.usecase.UpdateMastery: calculated | old=60 | obtained=0 | new=30
2025-11-22T10:30:10.200+0700 INFO  consumer.handleMessage: Successfully updated mastery | user_id=user123 | new_score=30
```

## Troubleshooting

### Database Connection Error

**Error:** `pq: relation "skill_mastery" does not exist`

**Solution:**
```bash
psql -U postgres -d postgres -f migration/01_create_skill_mastery_table.sql
```

### RabbitMQ Connection Error

**Error:** `Failed to connect to RabbitMQ`

**Solution:**
1. Check RabbitMQ is running: `sudo systemctl status rabbitmq-server`
2. Verify `RABBITMQ_URL` format: `amqp://user:pass@host:port/`
3. Check RabbitMQ logs: `sudo tail -f /var/log/rabbitmq/rabbit@*.log`

### SSL Mode Error

**Error:** `pq: unsupported sslmode "prefer"`

**Solution:** Set `POSTGRES_SSLMODE=disable` in environment variables

## Dependencies

| Service | Port | Purpose | Required |
|---------|------|---------|----------|
| PostgreSQL | 5432 | Store mastery data | Yes |
| RabbitMQ | 5672 | Receive submission events | Yes |
| Scoring Service | 8082 | Publishes events | Yes (indirect) |

## Performance Considerations

- Database queries use composite primary key (user_id, skill_tag) for fast lookups
- Consumer processes messages asynchronously
- Database connection pooling enabled
- Graceful shutdown ensures all messages are processed

## Additional Resources

- [CHANGELOG.md](./CHANGELOG.md) - Version history
- [Architecture Documentation](../architecture.md) - System overview

## Support

For issues:
1. Check logs for error details
2. Verify database and RabbitMQ connections
3. Review environment variables
4. Consult CHANGELOG.md
