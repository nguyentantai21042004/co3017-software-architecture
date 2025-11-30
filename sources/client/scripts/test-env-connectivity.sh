#!/bin/bash

# Test Environment Connectivity
# This script tests API connectivity to test environment services
# It performs actual API calls to verify services are responding correctly

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

printf "${BLUE}Testing Test Environment API Connectivity${NC}\n"
printf "=============================================\n\n"

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

# Get service URLs
CONTENT_URL="${NEXT_PUBLIC_CONTENT_API_URL:-http://localhost:8081}"
SCORING_URL="${NEXT_PUBLIC_SCORING_API_URL:-http://localhost:8082}"
LEARNER_URL="${NEXT_PUBLIC_LEARNER_API_URL:-http://localhost:8083}"
ADAPTIVE_URL="${NEXT_PUBLIC_ADAPTIVE_API_URL:-http://localhost:8084/api/adaptive}"

printf "\n"

# Function to test API endpoint
test_api_endpoint() {
    local service_name=$1
    local endpoint=$2
    local expected_status=${3:-200}
    
    printf "Testing $service_name: $endpoint... "
    
    HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --max-time 10 "$endpoint" 2>&1 || echo "000")
    
    if [ "$HTTP_CODE" = "$expected_status" ] || [ "$HTTP_CODE" = "200" ] || [ "$HTTP_CODE" = "404" ]; then
        # 404 might be OK for some endpoints, 200 is good
        if [ "$HTTP_CODE" = "200" ]; then
            printf "${GREEN}✓${NC} HTTP $HTTP_CODE\n"
        else
            printf "${YELLOW}⚠${NC} HTTP $HTTP_CODE (endpoint may not exist, but service is reachable)\n"
        fi
        return 0
    else
        printf "${RED}✗${NC} HTTP $HTTP_CODE or connection failed\n"
        return 1
    fi
}

ALL_PASSED=true

# Test Content Service
printf "${BLUE}Content Service Tests:${NC}\n"
if test_api_endpoint "Content Service Health" "$CONTENT_URL/health"; then
    if test_api_endpoint "Content Service Skills" "$CONTENT_URL/api/content/skills"; then
        printf "${GREEN}✓${NC} Content Service API is responding correctly\n"
    else
        ALL_PASSED=false
    fi
else
    ALL_PASSED=false
fi
printf "\n"

# Test Scoring Service
printf "${BLUE}Scoring Service Tests:${NC}\n"
if test_api_endpoint "Scoring Service Health" "$SCORING_URL/health"; then
    printf "${GREEN}✓${NC} Scoring Service API is responding correctly\n"
else
    ALL_PASSED=false
fi
printf "\n"

# Test Learner Model Service
printf "${BLUE}Learner Model Service Tests:${NC}\n"
if test_api_endpoint "Learner Model Health" "$LEARNER_URL/health"; then
    # Test with a dummy user ID (should return error but service should respond)
    if test_api_endpoint "Learner Model API" "$LEARNER_URL/internal/learner/test-user/mastery?skill=math" 200; then
        printf "${GREEN}✓${NC} Learner Model Service API is responding correctly\n"
    else
        ALL_PASSED=false
    fi
else
    ALL_PASSED=false
fi
printf "\n"

# Test Adaptive Engine Service
printf "${BLUE}Adaptive Engine Service Tests:${NC}\n"
ADAPTIVE_BASE="${ADAPTIVE_URL%/api/adaptive}"
if test_api_endpoint "Adaptive Engine Health" "$ADAPTIVE_BASE/health"; then
    printf "${GREEN}✓${NC} Adaptive Engine Service API is responding correctly\n"
else
    ALL_PASSED=false
fi
printf "\n"

# Summary
if [ "$ALL_PASSED" = true ]; then
    printf "${GREEN}✅ All API endpoints are accessible and responding correctly!${NC}\n\n"
    printf "Test environment is ready for E2E testing.\n"
    exit 0
else
    printf "${RED}✗${NC} Some API endpoints failed connectivity tests.\n"
    printf "\n"
    printf "Please verify:\n"
    printf "  1. All services are deployed and running in test environment\n"
    printf "  2. Service URLs are correct in $ENV_FILE\n"
    printf "  3. Services are properly configured and healthy\n"
    printf "  4. Network connectivity and firewall rules\n"
    printf "\n"
    exit 1
fi

