#!/bin/bash

# System Test Runner
# Runs comprehensive system tests across all 4 microservices

set -e

printf "🧪 System Test Runner\n"
printf "=====================\n\n"

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Check services
printf "Checking services...\n"
all_running=true

check_service() {
    local url=$1
    local name=$2
    
    if curl -s -f -o /dev/null "$url/health" 2>/dev/null; then
        printf "${GREEN}✓${NC} $name is running\n"
        return 0
    else
        printf "${RED}✗${NC} $name is NOT running\n"
        return 1
    fi
}

check_service "http://localhost:8081" "Content Service" || all_running=false
check_service "http://localhost:8082" "Scoring Service" || all_running=false
check_service "http://localhost:8080" "Learner Model" || all_running=false
check_service "http://localhost:8084/api/adaptive" "Adaptive Engine" || all_running=false

printf "\n"

if [ "$all_running" = false ]; then
    printf "${RED}ERROR: Not all services are running!${NC}\n"
    printf "\nPlease start all services:\n"
    printf "  cd src && ./scripts/start_services.sh\n\n"
    exit 1
fi

# Run system tests
printf "=====================\n"
printf "Running System Tests\n"
printf "=====================\n\n"

cd "$(dirname "$0")"

# Run all tests
go test -v -timeout 60s

exit_code=$?

printf "\n"
if [ $exit_code -eq 0 ]; then
    printf "${GREEN}✅ All system tests passed!${NC}\n"
else
    printf "${RED}❌ Some system tests failed!${NC}\n"
fi

exit $exit_code
