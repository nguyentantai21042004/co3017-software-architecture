# Learner Model Service (Golang)

**Architecture:** Microservices (API + Consumer)
**Database:** learner_db
**Technology:** Go 1.25.4, Gin, PostgreSQL, RabbitMQ, SQLBoiler

## Overview

Learner Model Service tracks user skill mastery levels. It consists of TWO separate services:

1. **API Service** (Port 8083): REST API for querying mastery scores
2. **Consumer Service**: Event processor for real-time mastery updates

**Database ORM:** Uses [SQLBoiler](https://github.com/aarondl/sqlboiler) for type-safe, efficient database operations. See [SQLBOILER_MIGRATION.md](./SQLBOILER_MIGRATION.md) for details.

**Consumer Documentation:** See [CONSUMER_SERVICE.md](./CONSUMER_SERVICE.md) for detailed consumer architecture and deployment.

## Setup

```bash
# 1. Initialize database
psql -U postgres -h localhost -p 5432 -f ../init-scripts/03-init-learner-db.sql

# 2. Install dependencies
go mod tidy

# 3. Run service
go run cmd/api/main.go
```

Service starts on **http://localhost:8083**

## How It Works

### 1. Consumes RabbitMQ Events

Listens to queue `learner.updates` for `SubmissionCompleted` events:

```json
{
  "event": "SubmissionCompleted",
  "user_id": "user_01",
  "skill_tag": "math_algebra",
  "score_obtained": 0,
  "timestamp": "2025-11-21T..."
}
```

### 2. Updates Mastery Score

Formula: `NewScore = (OldScore + ScoreObtained) / 2`

Example:
- Old mastery: 10
- Score obtained: 0 (wrong answer)
- New mastery: (10 + 0) / 2 = 5

### 3. Provides API for Adaptive Engine

```bash
GET /internal/learner/{user_id}/mastery?skill={skill}
```

## API Endpoints

### GET /internal/learner/:user_id/mastery

Query user's mastery for a skill.

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
```

**Response:**
```json
{
  "user_id": "user_01",
  "skill_tag": "math_algebra",
  "mastery_score": 5,
  "last_updated": "2025-11-21T..."
}
```

### GET /health

Health check endpoint.

## Testing Flow

### Test 1: Verify Initial State

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
# Expected: mastery_score=10 (from seed data)
```

### Test 2: Submit Wrong Answer (via Scoring Service)

```bash
# This will trigger RabbitMQ event
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "question_id": 1, "answer": "C"}'
```

### Test 3: Verify Updated Mastery

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
# Expected: mastery_score=5 (updated from 10)
```

## Logs

The service logs all event processing:

```
🎧 Listening for events on queue: learner.updates
📥 Received message: {"event":"SubmissionCompleted",...}
🧮 Calculating new mastery for user: user_01, skill: math_algebra
📊 Mastery update: user_01 [math_algebra] - Old: 10, Obtained: 0, New: 5
✅ Updated mastery: user=user_01, skill=math_algebra, new_score=5
```

## Dependencies

1. PostgreSQL (learner_db) - Must be running
2. RabbitMQ (5672) - Must be running to receive events
3. Scoring Service (8082) - Publishes events this service consumes

## Next Steps

1. ✅ Content Service (Port 8081) - Done
2. ✅ Scoring Service (Port 8082) - Done
3. ✅ Learner Model Service (Port 8083) - **YOU ARE HERE**
4. ⏭️ Adaptive Engine (Port 8084) - Orchestrate learning flow
