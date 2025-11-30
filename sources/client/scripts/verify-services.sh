#!/bin/bash

# Verify All Backend Services Health
# This script checks if all required backend services are running and healthy
# Used before running E2E tests to ensure services are ready

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

printf "${BLUE}Verifying Backend Services Health${NC}\n"
printf "=====================================\n\n"

# Service health endpoints
SERVICES=(
    "Content Service:http://localhost:8081/health"
    "Scoring Service:http://localhost:8082/health"
    "Learner Model Service:http://localhost:8083/health"
    "Adaptive Engine Service:http://localhost:8084/health"
)

ALL_HEALTHY=true
FAILED_SERVICES=()

for service_info in "${SERVICES[@]}"; do
    IFS=':' read -r service_name service_url <<< "$service_info"
    printf "Checking $service_name... "
    
    if curl -f -s "$service_url" > /dev/null 2>&1; then
        printf "${GREEN}✓${NC} $service_name is healthy\n"
    else
        printf "${RED}✗${NC} $service_name is not responding\n"
        ALL_HEALTHY=false
        FAILED_SERVICES+=("$service_name")
    fi
done

# Check infrastructure services
printf "\nChecking infrastructure services...\n"

# Check PostgreSQL
printf "Checking PostgreSQL... "
if docker exec its-postgres pg_isready -U postgres > /dev/null 2>&1; then
    printf "${GREEN}✓${NC} PostgreSQL is ready\n"
else
    printf "${RED}✗${NC} PostgreSQL is not responding\n"
    ALL_HEALTHY=false
    FAILED_SERVICES+=("PostgreSQL")
fi

# Check RabbitMQ
printf "Checking RabbitMQ... "
if curl -f -s http://localhost:15672 > /dev/null 2>&1; then
    printf "${GREEN}✓${NC} RabbitMQ is ready\n"
else
    printf "${RED}✗${NC} RabbitMQ is not responding\n"
    ALL_HEALTHY=false
    FAILED_SERVICES+=("RabbitMQ")
fi

printf "\n"

if [ "$ALL_HEALTHY" = true ]; then
    printf "${GREEN}✅ All services are healthy and ready for E2E testing!${NC}\n\n"
    exit 0
else
    printf "${RED}✗${NC} The following services are not healthy:\n"
    for service in "${FAILED_SERVICES[@]}"; do
        printf "  • $service\n"
    done
    printf "\n"
    printf "Please start services using:\n"
    printf "  ./scripts/start-services.sh\n"
    printf "\n"
    exit 1
fi

