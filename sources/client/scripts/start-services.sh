#!/bin/bash

# Start All Backend Services for E2E Testing
# This script starts all required backend services using Docker Compose
# Services: Content, Scoring, Learner Model, Adaptive Engine, and infrastructure (PostgreSQL, RabbitMQ)

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the project root directory (sources/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SOURCES_DIR="$SCRIPT_DIR"

printf "${BLUE}Starting Backend Services for E2E Testing${NC}\n"
printf "===============================================\n\n"

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    printf "${RED}✗${NC} Docker is not running. Please start Docker Desktop.\n"
    exit 1
fi
printf "${GREEN}✓${NC} Docker is running\n\n"

# Check if docker-compose is available
if ! command -v docker-compose > /dev/null 2>&1 && ! docker compose version > /dev/null 2>&1; then
    printf "${RED}✗${NC} docker-compose is not available. Please install docker-compose.\n"
    exit 1
fi

# Use docker compose (v2) if available, otherwise docker-compose (v1)
if docker compose version > /dev/null 2>&1; then
    DOCKER_COMPOSE="docker compose"
else
    DOCKER_COMPOSE="docker-compose"
fi

cd "$SOURCES_DIR"

# Step 1: Start infrastructure (PostgreSQL, RabbitMQ)
printf "${YELLOW}Step 1: Starting infrastructure services...${NC}\n"
$DOCKER_COMPOSE -f docker-compose.infra.yml up -d

# Wait for infrastructure to be ready
printf "Waiting for infrastructure services to be ready...\n"
sleep 5

# Check PostgreSQL health
printf "Checking PostgreSQL health...\n"
for i in {1..30}; do
    if docker exec its-postgres pg_isready -U postgres > /dev/null 2>&1; then
        printf "${GREEN}✓${NC} PostgreSQL is ready\n"
        break
    fi
    if [ $i -eq 30 ]; then
        printf "${RED}✗${NC} PostgreSQL failed to start within 30 seconds\n"
        exit 1
    fi
    sleep 1
done

# Check RabbitMQ health
printf "Checking RabbitMQ health...\n"
for i in {1..30}; do
    if curl -f http://localhost:15672 > /dev/null 2>&1; then
        printf "${GREEN}✓${NC} RabbitMQ is ready\n"
        break
    fi
    if [ $i -eq 30 ]; then
        printf "${RED}✗${NC} RabbitMQ failed to start within 30 seconds\n"
        exit 1
    fi
    sleep 1
done

printf "\n"

# Step 2: Start application services
printf "${YELLOW}Step 2: Starting application services...${NC}\n"
$DOCKER_COMPOSE -f docker-compose.yml up -d

# Wait for services to start
printf "Waiting for services to start...\n"
sleep 10

# Step 3: Verify all services are healthy
printf "\n${YELLOW}Step 3: Verifying service health...${NC}\n"

SERVICES=(
    "Content Service:http://localhost:8081/health"
    "Scoring Service:http://localhost:8082/health"
    "Learner Model Service:http://localhost:8083/health"
    "Adaptive Engine Service:http://localhost:8084/api/adaptive/health"
)

ALL_HEALTHY=true

for service_info in "${SERVICES[@]}"; do
    IFS=':' read -r service_name service_url <<< "$service_info"
    printf "Checking $service_name... "
    
    for i in {1..30}; do
        if curl -f "$service_url" > /dev/null 2>&1; then
            printf "${GREEN}✓${NC} $service_name is healthy\n"
            break
        fi
        if [ $i -eq 30 ]; then
            printf "${RED}✗${NC} $service_name is not responding\n"
            ALL_HEALTHY=false
        fi
        sleep 1
    done
done

printf "\n"

if [ "$ALL_HEALTHY" = true ]; then
    printf "${GREEN}✅ All services are running and healthy!${NC}\n\n"
    printf "Services:\n"
    printf "  • Content Service:      http://localhost:8081\n"
    printf "  • Scoring Service:      http://localhost:8082\n"
    printf "  • Learner Model:        http://localhost:8083\n"
    printf "  • Adaptive Engine:      http://localhost:8084\n"
    printf "  • PostgreSQL:           localhost:5432\n"
    printf "  • RabbitMQ:             http://localhost:15672\n"
    printf "\n"
    printf "To view logs:\n"
    printf "  docker-compose -f docker-compose.yml logs -f\n"
    printf "\n"
    printf "To stop services:\n"
    printf "  ./scripts/stop-services.sh\n"
    printf "\n"
    exit 0
else
    printf "${RED}✗${NC} Some services are not healthy. Please check the logs:\n"
    printf "  docker-compose -f docker-compose.yml logs\n"
    exit 1
fi

