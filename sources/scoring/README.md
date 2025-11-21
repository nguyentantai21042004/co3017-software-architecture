# Scoring Service (Golang)

**Port:** 8082
**Database:** scoring_db
**Technology:** Go 1.25.4, Gin, PostgreSQL, RabbitMQ, SQLBoiler

**Database ORM:** Uses [SQLBoiler](https://github.com/aarondl/sqlboiler) for type-safe, efficient database operations. See [SQLBOILER_MIGRATION.md](./SQLBOILER_MIGRATION.md) for details.

## Overview

Scoring Service handles answer submissions and publishes events to RabbitMQ.

## Setup

```bash
# 1. Initialize database
psql -U postgres -h localhost -p 5432 -f ../init-scripts/02-init-scoring-db.sql

# 2. Install dependencies
go mod tidy

# 3. Run service
go run cmd/api/main.go
```

Service starts on **http://localhost:8082**

## API

### POST /api/scoring/submit

```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "question_id": 1,
    "answer": "A"
  }'
```

**Response:**
```json
{
  "correct": true,
  "score": 100,
  "feedback": "Chính xác!"
}
```

## Testing

```bash
# Test wrong answer (triggers remedial)
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "question_id": 1, "answer": "C"}'
```

## Dependencies

1. Content Service (8081) - Must be running
2. PostgreSQL (scoring_db)
3. RabbitMQ (5672)
