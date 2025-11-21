# Intelligent Tutoring System (ITS) - Microservices Architecture

Complete implementation of an adaptive learning system using microservices architecture.

## System Overview

The ITS provides personalized learning paths by:
1. **Content Service** - Manages questions and learning materials
2. **Scoring Service** - Evaluates answers and publishes events
3. **Learner Model Service** - Tracks student mastery levels
4. **Adaptive Engine** - Recommends personalized content

## Architecture

```
┌─────────────┐
│   Client    │
└──────┬──────┘
       │
       ▼
┌─────────────────┐      ┌────────────────┐
│ Adaptive Engine │─────▶│ Learner Model  │
│   (Port 8084)   │      │  (Port 8083)   │
└────────┬────────┘      └────────────────┘
         │                      ▲
         │                      │
         ▼                ┌─────┴─────┐
┌─────────────────┐       │ RabbitMQ  │
│ Content Service │       └─────▲─────┘
│   (Port 8081)   │             │
└─────────────────┘       ┌─────┴────────┐
                          │    Scoring   │
                          │  (Port 8082) │
                          └──────────────┘
```

## Services & Databases

| Service | Port | Database | Technology | Purpose |
|---------|------|----------|------------|---------|
| Content | 8081 | content_db | Java/Spring Boot | Question management |
| Scoring | 8082 | scoring_db | Golang/Gin | Answer evaluation |
| Learner Model | 8083 | learner_db | Golang/Gin | Mastery tracking |
| Adaptive Engine | 8084 | None | Golang/Gin | Learning orchestration |

## Quick Start

### Prerequisites

- Java 17+
- Go 1.25.4+
- PostgreSQL 15+
- RabbitMQ 3+
- Maven 3.8+

### 1. Setup Databases

```bash
psql -U postgres -h localhost -p 5432 -f init-scripts/01-init-content-db.sql
psql -U postgres -h localhost -p 5432 -f init-scripts/02-init-scoring-db.sql
psql -U postgres -h localhost -p 5432 -f init-scripts/03-init-learner-db.sql
```

### 2. Start Services

```bash
# Terminal 1 - Content Service
cd content && mvn spring-boot:run

# Terminal 2 - Scoring Service
cd scoring && go run cmd/api/main.go

# Terminal 3 - Learner Model Service
cd learner-model && go run cmd/api/main.go

# Terminal 4 - Adaptive Engine
cd adaptive-engine && go run cmd/api/main.go
```

### 3. Verify All Services

```bash
curl http://localhost:8081/actuator/health  # Content
curl http://localhost:8082/health           # Scoring
curl http://localhost:8083/health           # Learner Model
curl http://localhost:8084/health           # Adaptive Engine
```

## Testing

### Quick Test (End-to-End Flow)

```bash
# 1. Check initial mastery (should be 10)
curl "http://localhost:8083/internal/learner/user_01/mastery?skill=math_algebra"

# 2. Submit wrong answer (mastery will drop to 5)
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "question_id": 1, "answer": "C"}'

# 3. Request next lesson (should recommend remedial)
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_01", "current_skill": "math_algebra"}'

# Expected: Recommends Question ID 2 (remedial content)
```

### Postman Collection

Import `ITS_Microservices.postman_collection.json` for complete API testing.

### Complete Testing Guide

See [TESTING_GUIDE.md](./TESTING_GUIDE.md) for detailed testing instructions.

## Project Structure

```
src/
├── init-scripts/           # Database initialization
│   ├── 01-init-content-db.sql
│   ├── 02-init-scoring-db.sql
│   └── 03-init-learner-db.sql
├── content/                # Java/Spring Boot Service
│   ├── src/main/java/...
│   ├── pom.xml
│   └── README.md
├── scoring/                # Golang Service
│   ├── cmd/api/main.go
│   ├── internal/...
│   └── README.md
├── learner-model/          # Golang Service
│   ├── cmd/api/main.go
│   ├── internal/...
│   └── README.md
├── adaptive-engine/        # Golang Service
│   ├── cmd/api/main.go
│   ├── internal/...
│   └── README.md
├── TESTING_GUIDE.md        # Complete testing guide
└── ITS_Microservices.postman_collection.json
```

## Key Features

### 1. Adaptive Learning
- Tracks student mastery per skill
- Recommends remedial content when mastery < 50%
- Progresses to harder content when mastery >= 50%

### 2. Event-Driven Architecture
- Scoring Service publishes submission events
- Learner Model Service consumes events asynchronously
- Decoupled services via RabbitMQ

### 3. Microservices Benefits
- Independent deployment
- Technology diversity (Java + Go)
- Scalable architecture
- Service isolation

## Configuration

Each service has its own `.env` file:

**Content Service** (`content/src/main/resources/application.yml`)
```yaml
server:
  port: 8081
spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/content_db
```

**Golang Services** (`.env` in each service)
```env
APP_PORT=8082  # or 8083, 8084
POSTGRES_DB=scoring_db  # or learner_db
RABBITMQ_URL=amqp://admintest:adminTest2025@localhost:5672/
```

## Documentation

- [Content Service README](./content/README.md)
- [Scoring Service README](./scoring/README.md)
- [Learner Model Service README](./learner-model/README.md)
- [Adaptive Engine README](./adaptive-engine/README.md)
- [Complete Testing Guide](./TESTING_GUIDE.md)

## Troubleshooting

### Services Won't Start

```bash
# Check if ports are in use
lsof -i :8081
lsof -i :8082
lsof -i :8083
lsof -i :8084
```

### Database Connection Failed

```bash
# Verify databases exist
psql -U postgres -h localhost -p 5432 -c "\l" | grep "_db"
```

### RabbitMQ Not Connected

```bash
# Check RabbitMQ Management UI
open http://localhost:15672
# Login: admintest / adminTest2025
```

## Demo Scenario

**Adaptive Learning in Action:**

1. **Initial State**: User has low algebra mastery (10%)
2. **User Action**: Attempts hard question, answers incorrectly
3. **System Response**: 
   - Mastery drops to 5%
   - System detects struggle
   - Recommends easier remedial content
4. **Result**: User gets personalized learning path

## Architecture Decisions

- **Microservices**: For independent scalability
- **Event-Driven**: Async processing via RabbitMQ
- **Database-per-Service**: Data ownership and isolation
- **RESTful APIs**: Simple HTTP communication
- **Stateless Adaptive Engine**: No database, pure orchestration

## Team

CO3017 - Software Architecture Project
HCMUT - 2025

## License

Educational Project - HCMUT
