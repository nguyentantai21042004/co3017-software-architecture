# ITS Microservices - Complete Testing Guide

This guide walks you through testing the entire Intelligent Tutoring System end-to-end.

## 🚀 Quick Start

### 1. Initialize Databases

```bash
cd /Users/tantai/Workspaces/hcmut/co3017-software-architecture

# Run all 3 init scripts
psql -U postgres -h localhost -p 5432 -f src/init-scripts/01-init-content-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/02-init-scoring-db.sql
psql -U postgres -h localhost -p 5432 -f src/init-scripts/03-init-learner-db.sql
```

### 2. Start All Services

Open 4 terminal windows:

**Terminal 1 - Content Service (Java)**
```bash
cd src/content
mvn clean spring-boot:run
# Starts on: http://localhost:8081
```

**Terminal 2 - Scoring Service (Golang)**
```bash
cd src/scoring
go run cmd/api/main.go
# Starts on: http://localhost:8082
```

**Terminal 3 - Learner Model Service (Golang)**
```bash
cd src/learner-model
go run cmd/api/main.go
# Starts on: http://localhost:8083
```

**Terminal 4 - Adaptive Engine (Golang)**
```bash
cd src/adaptive-engine
go run cmd/api/main.go
# Starts on: http://localhost:8084
```

---

## 🧪 End-to-End Test Scenario

### Scenario: User Struggles → Gets Remedial Content

**Initial State:**
- user_01 has algebra mastery = 10 (from seed data)
- Question 1 (ID=1) is a hard algebra question
- Question 2 (ID=2) is an easy remedial question

---

### Step 1: Check Initial Mastery

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
```

**Expected Response:**
```json
{
  "user_id": "user_01",
  "skill_tag": "math_algebra",
  "mastery_score": 10,
  "last_updated": "2025-11-21T..."
}
```

✅ **Verification:** Mastery is 10 (low)

---

### Step 2: Get Hard Question from Content Service

```bash
curl http://localhost:8081/api/content/1
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "content": "Bài toán khó: Giải phương trình bậc hai x² + 5x + 6 = 0",
    "correct_answer": "A",
    "is_remedial": false,
    "difficulty_level": 3
  }
}
```

✅ **Verification:** This is a hard question (difficulty=3)

---

### Step 3: Submit WRONG Answer

```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "question_id": 1,
    "answer": "C"
  }'
```

**Expected Response:**
```json
{
  "correct": false,
  "score": 0,
  "feedback": "Sai rồi, hãy thử lại!"
}
```

**What Happens Behind the Scenes:**
1. ✅ Scoring Service scores answer: score = 0
2. ✅ Saves to scoring_db
3. ✅ Publishes event to RabbitMQ
4. ✅ Learner Model Service consumes event
5. ✅ Updates mastery: (10 + 0) / 2 = 5

---

### Step 4: Verify Mastery Updated

```bash
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"
```

**Expected Response:**
```json
{
  "user_id": "user_01",
  "skill_tag": "math_algebra",
  "mastery_score": 5,
  "last_updated": "2025-11-21T..."
}
```

✅ **Verification:** Mastery decreased from 10 → 5

---

### Step 5: Request Next Lesson (Adaptive Recommendation)

```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

**Expected Response:**
```json
{
  "next_lesson_id": 2,
  "reason": "Your mastery is 5%. Let's review the basics.",
  "mastery_score": 5,
  "content_type": "remedial"
}
```

✅ **Verification:** 
- Mastery 5 < 50 → Recommends remedial content
- Returns Question ID 2 (easy remedial question)

---

### Step 6: Verify Remedial Question

```bash
curl http://localhost:8081/api/content/2
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "id": 2,
    "content": "Bài ôn tập: Nhắc lại quy tắc chuyển vế...",
    "correct_answer": "A",
    "is_remedial": true,
    "difficulty_level": 1
  }
}
```

✅ **Verification:** This is an easy remedial question!

---

## 🎉 Success!

The complete adaptive flow works:
1. User answers hard question incorrectly
2. Mastery score drops
3. Adaptive Engine detects low mastery
4. System recommends easier remedial content

---

## 📋 Quick Test Commands

### Test All Health Endpoints

```bash
curl http://localhost:8081/actuator/health  # Content
curl http://localhost:8082/health           # Scoring
curl http://localhost:8083/health           # Learner Model
curl http://localhost:8084/health           # Adaptive Engine
```

### Test Content Service

```bash
# Get question by ID
curl http://localhost:8081/api/content/1

# Get remedial question
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=remedial"

# Get standard question
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=standard"
```

### Test Scoring Service

```bash
# Submit correct answer
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_02", "question_id": 3, "answer": "B"}'

# Submit wrong answer
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "question_id": 1, "answer": "C"}'
```

### Test Learner Model Service

```bash
# Get mastery for user_01
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"

# Get mastery for user_02 (high mastery)
curl "http://localhost:8083/internal/learner/user_02/mastery?skill=math_algebra"
```

### Test Adaptive Engine

```bash
# For low mastery user (should get remedial)
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "current_skill": "math_algebra"}'

# For high mastery user (should get standard)
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_02", "current_skill": "math_algebra"}'
```

---

## 🐛 Troubleshooting

### RabbitMQ Not Connected

Check RabbitMQ is running:
```bash
curl http://localhost:15672/api/overview -u admintest:adminTest2025
```

### Database Not Found

Re-run init scripts:
```bash
psql -U postgres -h localhost -p 5432 -f src/init-scripts/01-init-content-db.sql
```

### Service Not Starting

Check logs and ports:
```bash
# Check if port is in use
lsof -i :8081
lsof -i :8082
lsof -i :8083
lsof -i :8084
```

---

## 📦 Service Ports Summary

| Service | Port | Database | Purpose |
|---------|------|----------|---------|
| Content | 8081 | content_db | Questions & content |
| Scoring | 8082 | scoring_db | Score answers, publish events |
| Learner Model | 8083 | learner_db | Track mastery, consume events |
| Adaptive Engine | 8084 | None (stateless) | Orchestrate adaptive learning |

---

## 🎯 Expected Flow Diagram

```
User answers wrong
       ↓
Scoring Service (8082)
   ├─→ Save to scoring_db
   └─→ Publish to RabbitMQ
       ↓
Learner Model Service (8083)
   ├─→ Consume event
   └─→ Update mastery in learner_db
       ↓
User requests next lesson
       ↓
Adaptive Engine (8084)
   ├─→ Query mastery from Learner Model (8083)
   ├─→ Apply adaptive logic (< 50 = remedial)
   └─→ Get content from Content Service (8081)
       ↓
Return remedial question to user
```
