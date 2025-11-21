# Content Service (Java/Spring Boot)

**Port:** 8081
**Database:** content_db
**Technology:** Java 17, Spring Boot 3.5.6, PostgreSQL

## Overview

Content Service manages learning questions and provides content recommendation APIs for the Adaptive Engine.

## Database Setup

Run the initialization script first:

```bash
psql -U postgres -h localhost -p 5432 -f ../init-scripts/01-init-content-db.sql
```

This creates:
- Database: `content_db`
- Table: `questions` (5 sample questions)

## Running the Service

### Option 1: Maven

```bash
# From this directory
mvn clean install
mvn spring-boot:run
```

### Option 2: IDE

Run `ContentServiceApplication.java`

### Verify

Service starts on **http://localhost:8081**

## API Endpoints

### 1. Get Question by ID

```bash
GET http://localhost:8081/api/content/{id}

# Example
curl http://localhost:8081/api/content/1
```

**Response:**
```json
{
  "success": true,
  "message": "Success",
  "data": {
    "id": 1,
    "content": "Bài toán khó: Giải phương trình bậc hai x² + 5x + 6 = 0",
    "options": ["A. x = -2 và x = -3", "B. x = 2 và x = 3", ...],
    "correct_answer": "A",
    "skill_tag": "math_algebra",
    "difficulty_level": 3,
    "is_remedial": false
  }
}
```

### 2. Recommend Question ⭐ (Core API for Adaptive Engine)

```bash
GET http://localhost:8081/api/content/recommend?skill={skill}&type={type}

# Example 1: Get remedial question for algebra
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=remedial"

# Example 2: Get standard question
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=standard"
```

**Parameters:**
- `skill` (required): "math_algebra", "math_geometry", etc.
- `type` (optional, default="standard"): "remedial" or "standard"
- `random` (optional, default=false): Return random question

**Response (Remedial):**
```json
{
  "success": true,
  "message": "Recommended remedial question for skill: math_algebra",
  "data": {
    "id": 2,
    "content": "Bài ôn tập: Nhắc lại quy tắc chuyển vế...",
    "is_remedial": true,
    "difficulty_level": 1
  }
}
```

### 3. Get All Questions for a Skill

```bash
curl http://localhost:8081/api/content/skill/math_algebra
```

### 4. Get Questions by Skill and Difficulty

```bash
curl http://localhost:8081/api/content/skill/math_algebra/difficulty/3
```

## Sample Data (5 Questions)

| ID | Skill | Type | Difficulty | Purpose |
|----|-------|------|------------|---------|
| 1 | math_algebra | Standard | 3 (Hard) | Main test question |
| **2** | **math_algebra** | **Remedial** | **1 (Easy)** | **Adaptive recommendation** |
| 3 | math_algebra | Standard | 2 (Medium) | Normal progression |
| 4 | math_geometry | Remedial | 1 (Easy) | Geometry review |
| 5 | math_geometry | Standard | 3 (Hard) | Challenge |

## Adaptive Engine Integration

```go
// Adaptive Engine calls this endpoint
if masteryScore < 50 {
    // Student struggling -> recommend remedial
    url = "http://localhost:8081/api/content/recommend?skill=math_algebra&type=remedial"
} else {
    // Student doing well -> standard content
    url = "http://localhost:8081/api/content/recommend?skill=math_algebra&type=standard"
}
```

## Testing

```bash
# Test 1: Get remedial question
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=remedial"
# Expected: Returns Question ID 2 (is_remedial=true)

# Test 2: Get standard question
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=standard"
# Expected: Returns Question ID 1 or 3 (is_remedial=false)

# Test 3: Get question by ID
curl http://localhost:8081/api/content/1
# Expected: Returns hard algebra question
```

## Troubleshooting

### Database Connection Error

```bash
# Check if database exists
psql -U postgres -h localhost -p 5432 -c "\l" | grep content_db

# If not, run init script
psql -U postgres -h localhost -p 5432 -f ../init-scripts/01-init-content-db.sql
```

### No Questions Found

```bash
# Check question count
psql -U postgres -h localhost -p 5432 -d content_db -c "SELECT COUNT(*) FROM questions;"
# Expected: 5

# If 0, re-run init script
```

## Configuration

Edit `src/main/resources/application.yml`:

```yaml
server:
  port: 8081

spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/content_db
    username: postgres
    password: postgres
```

## Next Steps

1. ✅ Content Service (Port 8081) - **YOU ARE HERE**
2. ⏭️ Scoring Service (Port 8082) - Submit answers
3. ⏭️ Learner Model Service (Port 8083) - Track mastery
4. ⏭️ Adaptive Engine (Port 8084) - Orchestrate learning
