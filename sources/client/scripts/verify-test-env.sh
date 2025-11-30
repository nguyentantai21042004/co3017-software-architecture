#!/bin/bash

# Verify Test Environment Deployment
# This script verifies that a deployed test environment is accessible and all services are healthy
# It reads environment URLs from .env.test or environment variables

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLIENT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

printf "${BLUE}Verifying Test Environment Deployment${NC}\n"
printf "==========================================\n\n"

cd "$CLIENT_DIR"

# Determine environment (default: test)
ENV="${1:-test}"

if [ "$ENV" != "test" ] && [ "$ENV" != "staging" ]; then
    printf "${RED}✗${NC} Invalid environment: $ENV\n"
    printf "Valid environments: test, staging\n"
    exit 1
fi

printf "${YELLOW}Environment: $ENV${NC}\n\n"

# Load environment configuration
ENV_FILE=".env.$ENV"
if [ ! -f "$ENV_FILE" ] && [ -f ".env.local" ]; then
    printf "${YELLOW}⚠${NC} $ENV_FILE not found, using .env.local\n"
    ENV_FILE=".env.local"
fi

if [ -f "$ENV_FILE" ]; then
    printf "Loading configuration from $ENV_FILE...\n"
    set -a
    source "$ENV_FILE"
    set +a
else
    printf "${YELLOW}⚠${NC} No .env file found. Using environment variables or defaults.\n"
fi

# Set environment variable
export NEXT_PUBLIC_ENV="$ENV"

# Get service URLs from environment or use defaults
CONTENT_URL="${NEXT_PUBLIC_CONTENT_API_URL:-http://localhost:8081}"
SCORING_URL="${NEXT_PUBLIC_SCORING_API_URL:-http://localhost:8082}"
LEARNER_URL="${NEXT_PUBLIC_LEARNER_API_URL:-http://localhost:8083}"
ADAPTIVE_URL="${NEXT_PUBLIC_ADAPTIVE_API_URL:-http://localhost:8084/api/adaptive}"
CLIENT_URL="${NEXT_PUBLIC_CLIENT_URL:-http://localhost:3000}"

printf "\n${BLUE}Service URLs:${NC}\n"
printf "  Content Service:   $CONTENT_URL\n"
printf "  Scoring Service:    $SCORING_URL\n"
printf "  Learner Model:      $LEARNER_URL\n"
printf "  Adaptive Engine:    $ADAPTIVE_URL\n"
printf "  Client Application: $CLIENT_URL\n"
printf "\n"

# Verify connectivity
ALL_ACCESSIBLE=true
FAILED_SERVICES=()

# Function to check service health
check_service() {
    local service_name=$1
    local service_url=$2
    local health_endpoint=$3
    
    printf "Checking $service_name... "
    
    # Try health endpoint first, then base URL
    if [ -n "$health_endpoint" ]; then
        if curl -f -s --max-time 10 "$health_endpoint" > /dev/null 2>&1; then
            printf "${GREEN}✓${NC} $service_name is accessible and healthy\n"
            return 0
        fi
    fi
    
    # Fallback to base URL check
    if curl -f -s --max-time 10 "$service_url" > /dev/null 2>&1; then
        printf "${GREEN}✓${NC} $service_name is accessible\n"
        return 0
    else
        printf "${RED}✗${NC} $service_name is not accessible\n"
        return 1
    fi
}

# Check Content Service
if ! check_service "Content Service" "$CONTENT_URL" "$CONTENT_URL/health"; then
    ALL_ACCESSIBLE=false
    FAILED_SERVICES+=("Content Service")
fi

# Check Scoring Service
if ! check_service "Scoring Service" "$SCORING_URL" "$SCORING_URL/health"; then
    ALL_ACCESSIBLE=false
    FAILED_SERVICES+=("Scoring Service")
fi

# Check Learner Model Service
if ! check_service "Learner Model Service" "$LEARNER_URL" "$LEARNER_URL/health"; then
    ALL_ACCESSIBLE=false
    FAILED_SERVICES+=("Learner Model Service")
fi

# Check Adaptive Engine Service
# Extract base URL for health check (remove /api/adaptive)
ADAPTIVE_BASE="${ADAPTIVE_URL%/api/adaptive}"
if ! check_service "Adaptive Engine Service" "$ADAPTIVE_URL" "$ADAPTIVE_BASE/health"; then
    ALL_ACCESSIBLE=false
    FAILED_SERVICES+=("Adaptive Engine Service")
fi

# Check Client Application
printf "Checking Client Application... "
if curl -f -s --max-time 10 "$CLIENT_URL" > /dev/null 2>&1; then
    printf "${GREEN}✓${NC} Client Application is accessible\n"
else
    printf "${YELLOW}⚠${NC} Client Application is not accessible (may be expected if not deployed)\n"
fi

printf "\n"

# Summary
if [ "$ALL_ACCESSIBLE" = true ]; then
    printf "${GREEN}✅ All backend services are accessible and healthy!${NC}\n\n"
    printf "Test environment is ready for E2E testing.\n"
    exit 0
else
    printf "${RED}✗${NC} The following services are not accessible:\n"
    for service in "${FAILED_SERVICES[@]}"; do
        printf "  • $service\n"
    done
    printf "\n"
    printf "Please verify:\n"
    printf "  1. Test environment is deployed and running\n"
    printf "  2. Service URLs in $ENV_FILE are correct\n"
    printf "  3. Network connectivity to test environment\n"
    printf "  4. Firewall rules allow access to test environment\n"
    printf "\n"
    exit 1
fi

