# Intelligent Tutoring System (ITS) - Microservices Architecture

Complete implementation of an adaptive learning system using microservices architecture, fully containerized with Docker.

## System Overview

The ITS provides personalized learning paths by:
1. **Content Service** - Manages questions and learning materials
2. **Scoring Service** - Evaluates answers and publishes events
3. **Learner Model Service** - Tracks student mastery levels
4. **Adaptive Engine** - Recommends personalized content
5. **Client** - Web-based frontend for students

## Architecture

```
┌──────────────────────────────────────────────────────────────────┐
│                         Docker Network                           │
│                         (its-network)                            │
│                                                                  │
│  ┌─────────────┐                                                 │
│  │   Client    │ :3000                                           │
│  │  (Next.js)  │                                                 │
│  └──────┬──────┘                                                 │
│         │                                                        │
│         ▼                                                        │
│  ┌─────────────────┐      ┌────────────────┐                     │
│  │ Adaptive Engine │─────▶│ Learner Model  │                     │
│  │    :8084        │      │  API :8083     │                     │
│  └────────┬────────┘      └────────┬───────┘                     │
│           │                        │                             │
│           ▼                        │                             │
│  ┌─────────────────┐               │                             │
│  │ Content Service │               │                             │
│  │    :8081        │               │                             │
│  └─────────────────┘               │                             │
│                             ┌──────┴───────┐                     │
│                             │   RabbitMQ   │                     │
│                             │    :5672     │                     │
│                             └──────▲───────┘                     │
│                                    │                             │
│                             ┌──────┴────────┐                    │
│                             │    Scoring    │                    │
│                             │    :8082      │                    │
│                             └───────────────┘                    │
│                                                                  │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐            │
│  │ PostgreSQL   │  │ PostgreSQL   │  │ PostgreSQL   │            │
│  │ Content DB   │  │ Scoring DB   │  │ Learner DB   │            │
│  │   :5433      │  │   :5434      │  │   :5435      │            │
│  └──────────────┘  └──────────────┘  └──────────────┘            │
│                                                                  │
│  ┌──────────────┐                                                │
│  │    MinIO     │                                                │
│  │   :9000      │                                                │
│  └──────────────┘                                                │
└──────────────────────────────────────────────────────────────────┘
```

## Services & Databases

| Service | Port | Database | Technology | Purpose |
|---------|------|----------|------------|---------|
| **Content** | 8081 | content_db (5433) | Java 17/Spring Boot | Question management |
| **Scoring** | 8082 | scoring_db (5434) | Go 1.23/Gin | Answer evaluation |
| **Learner Model** | 8083 | learner_db (5435) | Go 1.23/Gin | Mastery tracking |
| **Adaptive Engine** | 8084 | None | Go 1.23/Gin | Learning orchestration |
| **Client** | 3000 | None | Next.js 15/Node 20 | User Interface |

## Quick Start

### Prerequisites

- Docker Desktop 4.0+
- Make (optional, but recommended)

### 1. Start All Services

We provide a `Makefile` to simplify Docker operations.

```bash
# Setup everything (Build, Start, Init DB, Check Health)
make setup
```

Or using Docker Compose directly:

```bash
# Start services
docker-compose up -d

# Initialize databases (run once)
make db-init
```

### 2. Access Applications

- **Frontend**: [http://localhost:3000](http://localhost:3000)
- **RabbitMQ**: [http://localhost:15672](http://localhost:15672) (user: `admintest`, pass: `adminTest2025`)
- **MinIO**: [http://localhost:9001](http://localhost:9001) (user: `minioadmin`, pass: `minioadmin123`)

### 3. Verify Services

```bash
make health
```

Or manually:
```bash
curl http://localhost:8081/health  # Content
curl http://localhost:8082/health  # Scoring
curl http://localhost:8083/health  # Learner Model
curl http://localhost:8084/health  # Adaptive Engine
```

## Development

### Useful Commands

```bash
make help           # Show all available commands
make logs           # View logs from all services
make status         # Check service status
make test           # Run end-to-end test scenario
make clean          # Stop and remove containers
make rebuild        # Rebuild and restart services
```

### Local Development (Without Docker)

If you prefer to run services locally while keeping infrastructure in Docker:

```bash
# Start only infrastructure (Databases, RabbitMQ, MinIO)
make dev

# Then run each service locally
cd content && mvn spring-boot:run
cd scoring && go run cmd/api/main.go
# ... etc
```

## Project Structure

```
sources/
├── docker-compose.yml      # Main Docker composition
├── Makefile                # Management commands
├── scripts/                # Database initialization & test scripts
├── content/                # Java/Spring Boot Service
├── scoring/                # Golang Service
├── learner-model/          # Golang Service
├── adaptive-engine/        # Golang Service
├── client/                 # Next.js Frontend
└── README.md               # Project documentation
```

## Testing

### End-to-End Test

Run a complete flow test (Check Mastery -> Submit Answer -> Get Recommendation):

```bash
make test
```

### Common Issues

1. **Ports in use**: Ensure ports 3000, 8081-8084, 5433-5435, 5672, 9000-9001 are free.
2. **Database Connection**: Ensure containers are healthy (`make status`).
3. **Build Failures**: Try `make build-no-cache`.
