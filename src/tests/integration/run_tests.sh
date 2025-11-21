#!/bin/bash

# Integration Test Runner Script
# This script helps run integration tests with proper setup verification

set -e

printf "🧪 Integration Test Runner\n"
printf "==========================\n"
printf "\n"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to check if a service is running
check_service() {
    local url=$1
    local name=$2
    
    if curl -s -f -o /dev/null "$url/health" 2>/dev/null; then
        printf "${GREEN}✓${NC} $name is running\n"
        return 0
    else
        printf "${RED}✗${NC} $name is NOT running at $url\n"
        return 1
    fi
}

# Check all services
printf "Checking services...\n"
all_running=true

check_service "http://localhost:8081" "Content Service" || all_running=false
check_service "http://localhost:8082" "Scoring Service" || all_running=false
check_service "http://localhost:8083" "Learner Model Service" || all_running=false
check_service "http://localhost:8084" "Adaptive Engine" || all_running=false

printf "\n"

if [ "$all_running" = false ]; then
    printf "${RED}ERROR: Not all services are running!${NC}\n"
    printf "\n"
    printf "Please start all services before running integration tests:\n"
    printf "  cd src && ./scripts/start_services.sh\n"
    printf "\n"
    exit 1
fi

# Check PostgreSQL
printf "Checking PostgreSQL...\n"
if pg_isready -h localhost -p 5432 >/dev/null 2>&1; then
    printf "${GREEN}✓${NC} PostgreSQL is running\n"
else
    printf "${YELLOW}⚠${NC}  Cannot verify PostgreSQL (pg_isready not found or DB not running)\n"
fi

# Check RabbitMQ
printf "Checking RabbitMQ...\n"
if curl -s -f -o /dev/null "http://localhost:15672" 2>/dev/null; then
    printf "${GREEN}✓${NC} RabbitMQ is running\n"
else
    printf "${YELLOW}⚠${NC}  Cannot verify RabbitMQ (management UI not accessible)\n"
fi

printf "\n"
printf "==========================\n"
printf "Running Integration Tests\n"
printf "==========================\n"
printf "\n"

# Run tests
cd "$(dirname "$0")"

if [ "$#" -eq 0 ]; then
    # Run all tests
    go test -v -timeout 60s
else
    # Run specific test
    go test -v -timeout 60s -run "$1"
fi

exit_code=$?

printf "\n"
if [ $exit_code -eq 0 ]; then
    printf "${GREEN}✅ All tests passed!${NC}\n"
else
    printf "${RED}❌ Some tests failed!${NC}\n"
fi

exit $exit_code

