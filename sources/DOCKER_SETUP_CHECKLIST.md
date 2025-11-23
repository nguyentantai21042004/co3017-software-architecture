# Docker Setup Checklist

Quick reference for setting up your ITS Microservices with Docker.

---

## ✅ Pre-Flight Checklist

Before running any commands:

- [ ] Docker and Docker Compose installed
- [ ] All schema files reviewed (see `SCHEMA_SYNC_REPORT.md`)
- [ ] No services running on ports: 3000, 8081, 8082, 8083, 8084, 5432, 5672, 9000, 9001, 15672
- [ ] At least 4GB RAM available for Docker

---

## 🚀 Quick Start (New Installation)

```bash
# 1. Build all Docker images
make build

# 2. Start everything (infrastructure + services + DB init + health check)
make setup

# 3. Verify all services are healthy
make health
```

**Expected Output:**
```
✓ Client (3000)
✓ Content Service (8081)
✓ Scoring Service (8082)
✓ Learner Model API (8083)
✓ Adaptive Engine (8084)
✓ RabbitMQ (15672)
✓ MinIO (9001)
```

**Access Points:**
- Frontend: http://localhost:3000
- Content API Docs: http://localhost:8081/swagger-ui.html
- RabbitMQ Management: http://localhost:15672 (guest/guest)
- MinIO Console: http://localhost:9001 (minioadmin/minioadmin)

---

## 🔄 Daily Development Workflow

### Starting Services

```bash
# Option 1: Start everything
make up

# Option 2: Infrastructure only (for local dev)
make dev
# Then run services locally:
# cd content && mvn spring-boot:run
# cd scoring && go run cmd/api/main.go
# etc.
```

### Stopping Services

```bash
# Stop all services (keeps data)
make down

# Stop and DELETE ALL DATA
make down-volumes  # ⚠️ WARNING: Destructive!
```

### Viewing Logs

```bash
# All services
make logs

# Infrastructure only
make logs-infra

# Application services only
make logs-app
```

### Checking Status

```bash
# Service status
make status

# Health checks
make health
```

---

## 🗄️ Database Operations

### Initialize Databases

```bash
# Initialize all 3 databases with schema + seed data
make db-init
```

This runs:
1. `scripts/01-init-content-db.sql` → content_db
2. `scripts/02-init-scoring-db.sql` → scoring_db
3. `scripts/03-init-learner-db.sql` → learner_db

### Connect to Databases

```bash
# Connect to content_db
make db-content

# Connect to scoring_db
make db-scoring

# Connect to learner_db
make db-learner
```

### Backup Databases

```bash
# Backup all databases to backups/ directory
make db-backup
```

Creates timestamped files:
- `backups/content_db_20231123_143022.sql.gz`
- `backups/scoring_db_20231123_143022.sql.gz`
- `backups/learner_db_20231123_143022.sql.gz`

---

## 🧪 Testing

### Quick E2E Test

```bash
make test
```

Runs a sample workflow:
1. Check initial mastery
2. Submit wrong answer
3. Request next lesson (should be remedial)

### Integration Tests

```bash
make test-integration
```

---

## 🐛 Troubleshooting

### Services Not Starting

```bash
# Check logs for errors
make logs

# Restart specific service
docker-compose -f docker-compose.yml restart content-service

# Restart all
make restart
```

### Database Connection Issues

```bash
# Check if postgres is running
make status | grep postgres

# Restart infrastructure
make restart-infra

# Recreate databases
make down-volumes
make infra
sleep 5
make db-init
```

### Port Already in Use

```bash
# Find process using port (e.g., 8081)
lsof -ti:8081

# Kill process
lsof -ti:8081 | xargs kill

# Or change port in docker-compose.yml
```

### Build Failures

```bash
# Clean build cache
make build-no-cache

# Full cleanup and rebuild
make clean
make build
```

---

## 🔧 Advanced Operations

### Scale Consumer Service

```bash
# Scale learner-model-consumer to 3 instances
make scale-consumer-3
```

### Open Shell in Container

```bash
# Open shell in content-service container
make shell-content-service

# Open shell in postgres container
make shell-postgres
```

### Start Individual Services

```bash
# Content Service + dependencies
make content

# Scoring Service + dependencies
make scoring

# Learner Model + dependencies
make learner

# Adaptive Engine
make adaptive

# Client/Frontend
make client
```

---

## 📊 Schema Management

### Current Schema Status

✅ All schemas are synchronized with application models (see `SCHEMA_SYNC_REPORT.md`)

### Schema Changes Made

| Table | Column | Changed From | Changed To |
|-------|--------|--------------|------------|
| questions | id | SERIAL | BIGSERIAL |
| questions | correct_answer | VARCHAR(255) | TEXT |
| questions | skill_tag | VARCHAR(50) | VARCHAR(100) |
| submissions | id | SERIAL | BIGSERIAL |
| submissions | question_id | INT | BIGINT |
| skill_mastery | skill_tag | VARCHAR(50) | VARCHAR(100) |

### Adding New Content

After `make db-init`, you can add your custom content:

```bash
# Method 1: Via psql
make db-content
\i /path/to/insert_devops_quiz.sql
\i /path/to/insert_kubernetes_quiz.sql
\i /path/to/consolidate_kubernetes_skill.sql

# Method 2: Via docker exec
docker-compose -f docker-compose.infra.yml exec -T postgres \
    psql -U postgres -d content_db < scripts/insert_devops_quiz.sql
```

---

## 🧹 Cleanup

### Remove Stopped Containers

```bash
make clean
```

### Nuclear Option (Remove Everything)

```bash
make clean-all  # ⚠️ Removes ALL Docker resources
```

---

## 📁 Important Files

```
sources/
├── Makefile                          # All commands defined here
├── docker-compose.infra.yml          # Infrastructure services
├── docker-compose.yml                # Application services
├── scripts/
│   ├── 01-init-content-db.sql       # Content DB schema ✅ Fixed
│   ├── 02-init-scoring-db.sql       # Scoring DB schema ✅ Fixed
│   ├── 03-init-learner-db.sql       # Learner DB schema ✅ Fixed
│   ├── insert_devops_quiz.sql       # DevOps questions
│   ├── insert_devops_remedial.sql   # DevOps remedial questions
│   ├── insert_kubernetes_quiz.sql   # Kubernetes questions
│   └── consolidate_kubernetes_skill.sql  # Merge K8s tags
├── SCHEMA_SYNC_REPORT.md            # Schema analysis report
└── DOCKER_SETUP_CHECKLIST.md        # This file
```

---

## 🎯 Common Scenarios

### Scenario 1: Fresh Start (First Time Setup)

```bash
make build        # Build images
make setup        # Start + init + health check
make health       # Verify
```

### Scenario 2: Daily Development

```bash
make up           # Start services
# ... do work ...
make down         # Stop services (keeps data)
```

### Scenario 3: Clean Slate (Reset Everything)

```bash
make down-volumes # Delete all data
make setup        # Rebuild from scratch
```

### Scenario 4: Database Schema Update

```bash
# 1. Backup current data
make db-backup

# 2. Update schema files (01-init-*.sql)

# 3. Recreate databases
make down-volumes
make infra
sleep 5
make db-init

# 4. Restore custom data if needed
make db-content
\i backups/custom_data.sql
```

### Scenario 5: Update One Service

```bash
# 1. Rebuild service
docker-compose -f docker-compose.yml build content-service

# 2. Restart it
docker-compose -f docker-compose.yml up -d content-service

# 3. Check logs
docker-compose -f docker-compose.yml logs -f content-service
```

---

## ⚡ Quick Commands

| Command | Action |
|---------|--------|
| `make` | Show help |
| `make setup` | Complete setup |
| `make up` | Start all |
| `make down` | Stop all |
| `make restart` | Restart all |
| `make status` | Show status |
| `make logs` | Show logs |
| `make health` | Health check |
| `make db-init` | Init databases |
| `make db-backup` | Backup all DBs |
| `make clean` | Cleanup |

---

## 📝 Notes

- **Data Persistence**: Volumes are used to persist data. Use `make down` to keep data, `make down-volumes` to delete.
- **Ports**: All ports are exposed on localhost. Change in docker-compose files if needed.
- **Logs**: Stored in `/var/log` inside containers. Access via `docker logs` or `make logs`.
- **Health Checks**: Each service has a `/health` endpoint on its port.
- **Environment Variables**: Defined in docker-compose files. Override with `.env` file if needed.

---

## ✅ Post-Setup Verification

After `make setup`, verify:

```bash
# 1. All services running
make status
# Should show all services "Up"

# 2. All health checks passing
make health
# Should show all ✓

# 3. Databases initialized
make db-content
\dt
# Should show "questions" table

# 4. Sample query
SELECT COUNT(*) FROM questions;
# Should return 5 (seed data)

# 5. Frontend accessible
curl http://localhost:3000
# Should return HTML
```

---

**Last Updated:** 2025-11-23
**Status:** ✅ Ready for Production
