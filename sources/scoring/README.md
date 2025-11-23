# Scoring Service (Golang)

**Port:** 8082
**Database:** scoring_db (PostgreSQL)
**Technology:** Go 1.23, Gin, PostgreSQL, RabbitMQ, SQLBoiler

## Overview

Scoring Service handles answer submissions and publishes events to RabbitMQ. It is fully containerized and integrated into the ITS microservices ecosystem.

## Quick Start (Docker)

```bash
# Start service via Docker Compose (from sources/ root)
make scoring
```

## Local Setup (Development)

1. **Start Infrastructure**:
   ```bash
   # Ensure Postgres and RabbitMQ are running
   cd ../ && make dev
   ```

2. **Run Service**:
   ```bash
   go mod tidy
   go run cmd/api/main.go
   ```

## Configuration

Environment variables (automatically set in `docker-compose.yml`):

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Service Port | 8082 |
| `POSTGRES_HOST` | Database Host | `postgres-scoring` (Docker) or `localhost` |
| `POSTGRES_PORT` | Database Port | `5432` (Docker) or `5434` (Local) |
| `RABBITMQ_URL` | RabbitMQ URL | `amqp://admintest:adminTest2025@rabbitmq:5672/` |

## API Endpoints

### POST /api/scoring/submit

Submit an answer for evaluation.

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

## Health Check

```bash
curl http://localhost:8082/health
```
