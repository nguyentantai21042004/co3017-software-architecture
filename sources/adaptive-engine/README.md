# Adaptive Engine Service

**Port:** 8084
**Database:** None (Stateless)
**Technology:** Go 1.25.4, Gin
**Architecture:** Module-First + Clean Architecture

## Overview

Adaptive Engine is the "brain" of the Intelligent Tutoring System (ITS). It orchestrates personalized learning by analyzing user mastery and recommending appropriate content.

### Key Responsibilities
- Query user mastery levels from Learner Model Service
- Apply adaptive logic to determine content difficulty
- Fetch appropriate questions from Content Service
- Provide personalized learning recommendations

### Adaptive Logic
```
IF mastery_score < 50:
    → Recommend REMEDIAL content (easier review material)
ELSE:
    → Recommend STANDARD content (normal progression)
```

## Architecture

### Module-First + Clean Architecture

```
adaptive-engine/
├── cmd/
│   └── api/
│       └── main.go                    # Application entry point
├── internal/
│   └── adaptive/                      # Adaptive module
│       ├── delivery/
│       │   └── http/
│       │       ├── handler.go         # HTTP request handlers
│       │       └── routes.go          # Route definitions
│       ├── usecase/
│       │   └── adaptive.go            # Business logic & orchestration
│       ├── error.go                   # Module-specific errors
│       ├── type.go                    # Input/Output types
│       └── interface.go               # UseCase interface
├── pkg/
│   ├── curl/                          # HTTP client utilities
│   └── log/                           # Logging utilities
└── config/                            # Configuration management
```

### Layer Responsibilities

- **Delivery Layer** (`delivery/http/`)
  - HTTP request/response handling
  - Request validation
  - Route mapping
  - Depends on: UseCase interface

- **UseCase Layer** (`usecase/`)
  - Orchestrates calls to external services
  - Implements adaptive recommendation logic
  - Decision making based on mastery scores
  - Pure business logic, no HTTP concerns

- **Types & Interfaces**
  - `type.go`: RecommendInput, RecommendOutput
  - `interface.go`: UseCase interface for dependency inversion
  - `error.go`: Domain-specific error definitions

## Setup

### Prerequisites
- Go 1.25.4+
- Learner Model Service running on port 8083
- Content Service running on port 8081

### Installation

```bash
# Clone and navigate to service
cd sources/adaptive-engine

# Install dependencies
go mod tidy

# Set environment variables (or use .env)
export APP_PORT=8084
export LEARNER_SERVICE_URL=http://localhost:8083
export CONTENT_SERVICE_URL=http://localhost:8081
export API_MODE=debug
export LOGGER_LEVEL=debug

# Run service
go run cmd/api/main.go
```

Service starts on **http://localhost:8084**

### Docker Setup

```bash
# Build image
docker build -t adaptive-engine:latest .

# Run container
docker run -p 8084:8084 \
  -e LEARNER_SERVICE_URL=http://learner-model:8083 \
  -e CONTENT_SERVICE_URL=http://content:8081 \
  adaptive-engine:latest
```

## API Reference

### POST /api/adaptive/next-lesson

Get personalized next lesson recommendation based on user's current mastery.

**Request:**
```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

**Request Body:**
```json
{
  "user_id": "string",      // Required: User identifier
  "current_skill": "string"  // Required: Skill tag (e.g., "math_algebra")
}
```

**Response (Low Mastery - Remedial):**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "next_lesson_id": 2,
    "reason": "Your mastery is 30%. Let's review the basics with easier content.",
    "mastery_score": 30,
    "content_type": "remedial"
  }
}
```

**Response (High Mastery - Standard):**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "next_lesson_id": 1,
    "reason": "Great! Your mastery is 75%. Continue with the next challenge.",
    "mastery_score": 75,
    "content_type": "standard"
  }
}
```

**Error Responses:**
```json
{
  "error_code": 400,
  "message": "Invalid request: user_id and current_skill are required",
  "data": null,
  "errors": null
}
```

### GET /health

Health check endpoint.

```bash
curl http://localhost:8084/health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "adaptive-engine",
  "timestamp": "2025-11-22T12:00:00Z"
}
```

## Workflow Example

### Complete User Journey

**Scenario:** User struggles with algebra, system adapts by providing remedial content

**1. Initial State**
```
User: user_01
Skill: math_algebra
Current Mastery: 60
```

**2. User Submits Wrong Answer** (via Scoring Service)
```bash
curl -X POST http://localhost:8082/api/scoring/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "question_id": 1,
    "answer": "C"
  }'
```

Result:
- Score: 0 (incorrect)
- RabbitMQ event published
- Learner Model updates mastery: 60 → 30

**3. Request Next Lesson** (Adaptive Engine)
```bash
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_01",
    "current_skill": "math_algebra"
  }'
```

Internal Flow:
```
Adaptive Engine
    ↓
1. Call Learner Model: GET /internal/learner/user_01/mastery?skill=math_algebra
   Response: mastery_score = 30
    ↓
2. Apply Logic: 30 < 50 → Recommend REMEDIAL
    ↓
3. Call Content Service: GET /api/content/recommend?skill=math_algebra&type=remedial
   Response: question_id = 2 (easier review content)
    ↓
4. Return Recommendation
```

**4. Response**
```json
{
  "error_code": 0,
  "message": "Success",
  "data": {
    "next_lesson_id": 2,
    "reason": "Your mastery is 30%. Let's review the basics with easier content.",
    "mastery_score": 30,
    "content_type": "remedial"
  }
}
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_PORT` | Service port | `8084` | No |
| `LEARNER_SERVICE_URL` | Learner Model Service URL | - | Yes |
| `CONTENT_SERVICE_URL` | Content Service URL | - | Yes |
| `API_MODE` | Server mode (debug/release) | `debug` | No |
| `LOGGER_LEVEL` | Log level (debug/info/warn/error) | `debug` | No |
| `LOGGER_MODE` | Logger mode (debug/production) | `debug` | No |
| `LOGGER_ENCODING` | Log encoding (console/json) | `console` | No |

### Sample .env File

```env
APP_PORT=8084
LEARNER_SERVICE_URL=http://localhost:8083
CONTENT_SERVICE_URL=http://localhost:8081
API_MODE=debug
LOGGER_LEVEL=debug
LOGGER_MODE=debug
LOGGER_ENCODING=console
```

## Testing

### Manual Testing

```bash
# Test 1: User with low mastery (should get remedial)
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_low", "current_skill": "math_algebra"}'

# Test 2: User with high mastery (should get standard)
curl -X POST http://localhost:8084/api/adaptive/next-lesson \
  -H "Content-Type: application/json" \
  -d '{"user_id": "user_high", "current_skill": "math_algebra"}'

# Test 3: Health check
curl http://localhost:8084/health
```

### Integration Testing

Ensure all dependent services are running:
```bash
# Terminal 1: Content Service
cd sources/content && mvn spring-boot:run

# Terminal 2: Scoring Service
cd sources/scoring && go run cmd/api/main.go

# Terminal 3: Learner Model Service
cd sources/learner-model && go run cmd/api/main.go

# Terminal 4: Adaptive Engine
cd sources/adaptive-engine && go run cmd/api/main.go
```

## Logging

The service provides detailed logs for debugging:

```
2025-11-22T12:00:00.000+0700 INFO  adaptive.usecase.RecommendNextLesson: starting | user_id=user_01 | skill=math_algebra
2025-11-22T12:00:00.100+0700 INFO  adaptive.usecase.RecommendNextLesson: fetched mastery | mastery_score=30
2025-11-22T12:00:00.150+0700 INFO  adaptive.usecase.RecommendNextLesson: applying logic | threshold=50 | recommendation=remedial
2025-11-22T12:00:00.250+0700 INFO  adaptive.usecase.RecommendNextLesson: fetched content | lesson_id=2 | type=remedial
2025-11-22T12:00:00.300+0700 INFO  adaptive.usecase.RecommendNextLesson: success | lesson_id=2
```

## Troubleshooting

### Service Connection Errors

**Error:** `failed to fetch mastery from learner service`

**Solution:**
1. Check Learner Model Service is running: `curl http://localhost:8083/health`
2. Verify `LEARNER_SERVICE_URL` environment variable
3. Check network connectivity

**Error:** `failed to fetch content from content service`

**Solution:**
1. Check Content Service is running: `curl http://localhost:8081/health`
2. Verify `CONTENT_SERVICE_URL` environment variable
3. Ensure Content Service has questions for the requested skill

### Invalid Requests

**Error:** `invalid request: user_id and current_skill are required`

**Solution:** Ensure both `user_id` and `current_skill` are provided in request body

## Dependencies

| Service | Port | Purpose | Required |
|---------|------|---------|----------|
| Content Service | 8081 | Provide questions | Yes |
| Learner Model Service | 8083 | Provide mastery scores | Yes |
| RabbitMQ | 5672 | (Not used directly) | No |
| PostgreSQL | 5432 | (Not used) | No |

## Performance Considerations

- Service is stateless - can be horizontally scaled
- Each request makes 2 HTTP calls (Learner Model + Content Service)
- Consider adding caching for mastery scores if needed
- Timeout configuration for external service calls

## Future Enhancements

- [ ] Multiple adaptive strategies (not just threshold-based)
- [ ] Machine learning integration for personalized recommendations
- [ ] Caching layer for frequently accessed mastery scores
- [ ] A/B testing framework for adaptive algorithms
- [ ] Detailed analytics and recommendation tracking

## Additional Resources

- [CHANGELOG.md](./CHANGELOG.md) - Version history and changes
- [Swagger Documentation](http://localhost:8084/swagger/index.html) - Interactive API docs
- [Architecture Documentation](../architecture.md) - Overall system architecture

## Support

For issues or questions:
1. Check logs for detailed error messages
2. Verify all dependent services are running
3. Review configuration and environment variables
4. Consult CHANGELOG.md for recent changes
