#!/bin/bash

# Integration Test Runner Script
# This script helps run integration tests with proper setup verification

set -e

echo "🧪 Integration Test Runner"
echo "=========================="
echo ""

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
        echo -e "${GREEN}✓${NC} $name is running"
        return 0
    else
        echo -e "${RED}✗${NC} $name is NOT running at $url"
        return 1
    fi
}

# Check all services
echo "Checking services..."
all_running=true

check_service "http://localhost:8081" "Content Service" || all_running=false
check_service "http://localhost:8082" "Scoring Service" || all_running=false
check_service "http://localhost:8083" "Learner Model Service" || all_running=false
check_service "http://localhost:8084" "Adaptive Engine" || all_running=false

echo ""

if [ "$all_running" = false ]; then
    echo -e "${RED}ERROR: Not all services are running!${NC}"
    echo ""
    echo "Please start all services before running integration tests:"
    echo "  1. Content Service:      cd src/content && mvn spring-boot:run"
    echo "  2. Scoring Service:      cd src/scoring && go run cmd/api/main.go"
    echo "  3. Learner Model:        cd src/learner-model && go run cmd/api/main.go"
    echo "  4. Adaptive Engine:      cd src/adaptive-engine && go run cmd/api/main.go"
    echo ""
    exit 1
fi

# Check PostgreSQL
echo "Checking PostgreSQL..."
if pg_isready -h localhost -p 5432 >/dev/null 2>&1; then
    echo -e "${GREEN}✓${NC} PostgreSQL is running"
else
    echo -e "${YELLOW}⚠${NC}  Cannot verify PostgreSQL (pg_isready not found or DB not running)"
fi

# Check RabbitMQ
echo "Checking RabbitMQ..."
if curl -s -f -o /dev/null "http://localhost:15672" 2>/dev/null; then
    echo -e "${GREEN}✓${NC} RabbitMQ is running"
else
    echo -e "${YELLOW}⚠${NC}  Cannot verify RabbitMQ (management UI not accessible)"
fi

echo ""
echo "=========================="
echo "Running Integration Tests"
echo "=========================="
echo ""

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

echo ""
if [ $exit_code -eq 0 ]; then
    echo -e "${GREEN}✅ All tests passed!${NC}"
else
    echo -e "${RED}❌ Some tests failed!${NC}"
fi

exit $exit_code
