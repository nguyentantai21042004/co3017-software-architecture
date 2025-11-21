# Adaptive Engine (Golang)

**Port:** 8084
**Database:** None (Stateless)
**Technology:** Go 1.25.4, Gin

## Overview

Adaptive Engine is the "brain" of the ITS system. It orchestrates calls to Learner Model and Content services to recommend personalized learning paths.

## Setup

```bash
# Install dependencies
go mod tidy

# Run service
go run cmd/api/main.go
```

Service starts on **http://localhost:8084**

## How It Works

### Adaptive Logic Flow

```
1. Client requests next lesson
   ↓
2. Adaptive Engine calls Learner Model Service
   → Get user's mastery score
   ↓
3. Apply Adaptive Rule:
   - If mastery < 50 → Recommend "remedial" content
   - If mastery >= 50 → Recommend "standard" content
   ↓
4. Adaptive Engine calls Content Service
   → Get recommended question based on type
   ↓
5. Return recommendation to client
```

## API Endpoint

### POST /api/adaptive/next-lesson

Get personalized next lesson recommendation.

**Request:**
```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

**Response (Low Mastery - Remedial):**
```json
{
  "next_lesson_id": 2,
  "reason": "Your mastery is 5%. Let's review the basics.",
  "mastery_score": 5,
  "content_type": "remedial"
}
```

**Response (High Mastery - Standard):**
```json
{
  "next_lesson_id": 1,
  "reason": "Great! Your mastery is 80%. Continue with the next challenge.",
  "mastery_score": 80,
  "content_type": "standard"
}
```

## Complete Test Scenario

### Scenario: User Struggles → Gets Remedial Content

**Initial State:**
- user_01 has algebra mastery = 10 (low)

**Step 1: Submit Wrong Answer**
```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "question_id": 1,
    "answer": "C"
  }'
```
Result: Score = 0, RabbitMQ event → Mastery updated to 5

**Step 2: Request Next Lesson**
```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

**Result:** Adaptive Engine detects mastery=5 < 50, recommends Question ID 2 (remedial)

## Dependencies

The Adaptive Engine needs:
1. **Learner Model Service** (8083) - To query mastery scores
2. **Content Service** (8081) - To get recommended questions

Both services must be running.

## Configuration

Edit `.env`:

```env
APP_PORT=8084
LEARNER_SERVICE_URL=http://localhost:8083
CONTENT_SERVICE_URL=http://localhost:8081
```

## Logs

```
🧠 Adaptive Engine: user=user_01, skill=math_algebra
📊 Current mastery: 5
🔄 Recommending REMEDIAL (score=5 < 50)
📚 Recommended question ID: 2 (type: remedial)
```

## All Services Summary

| Service | Port | Purpose |
|---------|------|---------|
| Content | 8081 | Manage questions |
| Scoring | 8082 | Score answers, publish events |
| Learner Model | 8083 | Track mastery, consume events |
| **Adaptive Engine** | **8084** | **Orchestrate adaptive learning** |
