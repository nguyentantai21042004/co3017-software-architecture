# Learner Model Service (Golang)

**Port:** 8083 (API)
**Consumer:** Background Worker
**Database:** learner_db (PostgreSQL)
**Storage:** MinIO
**Technology:** Go 1.23, Gin, PostgreSQL, RabbitMQ, MinIO

## Overview

Learner Model Service tracks student mastery levels. It consists of:
1. **API Service**: REST endpoints for querying mastery (Port 8083).
2. **Consumer Service**: Processes scoring events from RabbitMQ to update mastery.

## Quick Start (Docker)

```bash
# Start both API and Consumer via Docker Compose (from sources/ root)
make learner
```

## Local Setup (Development)

1. **Start Infrastructure**:
   ```bash
   # Ensure Postgres, RabbitMQ, and MinIO are running
   cd ../ && make dev
   ```

2. **Run API**:
   ```bash
   go run cmd/api/main.go
   ```

3. **Run Consumer**:
   ```bash
   go run cmd/consumer/main.go
   ```

## Configuration

Environment variables (automatically set in `docker-compose.yml`):

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | API Port | 8083 |
| `POSTGRES_HOST` | Database Host | `postgres-learner` (Docker) or `localhost` |
| `POSTGRES_PORT` | Database Port | `5432` (Docker) or `5435` (Local) |
| `RABBITMQ_URL` | RabbitMQ URL | `amqp://admintest:adminTest2025@rabbitmq:5672/` |
| `MINIO_ENDPOINT` | MinIO URL | `minio:9000` (Docker) or `localhost:9000` |

## API Endpoints

### GET /internal/learner/:user_id/mastery

Get mastery score for a specific skill.

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
```

## Health Check

```bash
curl http://localhost:8083/health
```
