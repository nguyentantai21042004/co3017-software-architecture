# Adaptive Engine Service (Golang)

**Port:** 8084
**Database:** Stateless (No Database)
**Technology:** Go 1.25.4, Gin
**Dependencies:** Content Service, Learner Model Service

## Overview

Adaptive Engine is the orchestration service that recommends personalized content based on student mastery. It communicates with the Learner Model to get mastery scores and Content Service to fetch appropriate questions.

## Quick Start (Docker)

```bash
# Start service via Docker Compose (from sources/ root)
make adaptive
```

## Local Setup (Development)

1. **Start Infrastructure and Dependencies**:
   ```bash
   # Ensure other services are running
   cd ../ && make dev
   # You also need Content and Learner services running
   ```

2. **Run Service**:
   ```bash
   go run cmd/api/main.go
   ```

## Configuration

Environment variables (automatically set in `docker-compose.yml`):

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Service Port | 8084 |
| `CONTENT_SERVICE_URL` | Content Service URL | `http://content-service:8081` |
| `LEARNER_SERVICE_URL` | Learner Model URL | `http://learner-model-api:8083` |

## API Endpoints

### POST /api/adaptive/next-lesson

Get personalized recommendation.

```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

**Response:**
```json
{
  "next_lesson_id": 2,
  "reason": "Your mastery is 30%. Let's review the basics.",
  "content_type": "remedial"
}
```

## Health Check

```bash
curl http://localhost:8084/health
```
