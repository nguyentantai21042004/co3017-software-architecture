# Content Service (Java/Spring Boot)

**Port:** 8081
**Database:** content_db (PostgreSQL)
**Technology:** Java 17, Spring Boot 3.5.6, Maven, PostgreSQL

## Overview

Content Service manages questions and learning materials. It provides APIs for retrieving questions by ID, skill, and difficulty, and includes logic for content recommendation.

## Quick Start (Docker)

```bash
# Start service via Docker Compose (from sources/ root)
make content
```

## Local Setup (Development)

1. **Start Infrastructure**:
   ```bash
   # Ensure Postgres is running
   cd ../ && make dev
   ```

2. **Run Service**:
   ```bash
   mvn spring-boot:run
   ```

## Configuration

Environment variables (automatically set in `docker-compose.yml`):

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Service Port | 8081 |
| `SPRING_DATASOURCE_URL` | Database URL | `jdbc:postgresql://postgres-content:5432/content_db` |

## API Endpoints

### GET /api/content/{id}

Get question details.

```bash
curl http://localhost:8081/api/content/1
```

### GET /api/content/recommend

Recommend content based on skill and type.

```bash
curl "http://localhost:8081/api/content/recommend?skill=math_algebra&type=remedial"
```

## Health Check

```bash
curl http://localhost:8081/health
```
