#!/bin/bash

# Run E2E Tests Against Test Environment
# This script runs E2E tests against a deployed test environment
# It validates environment configuration and runs Playwright tests

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

printf "${BLUE}Running E2E Tests Against Test Environment${NC}\n"
printf "===============================================\n\n"

cd "$CLIENT_DIR"

# Determine environment (default: test)
ENV="${1:-test}"

if [ "$ENV" != "test" ] && [ "$ENV" != "staging" ]; then
    printf "${RED}✗${NC} Invalid environment: $ENV\n"
    printf "Valid environments: test, staging\n"
    exit 1
fi

printf "${YELLOW}Environment: $ENV${NC}\n\n"

# Step 1: Load environment configuration
printf "${YELLOW}Step 1: Loading environment configuration...${NC}\n"
if [ -f ".env.$ENV" ]; then
    set -a
    source ".env.$ENV"
    set +a
    printf "${GREEN}✓${NC} Loaded configuration from .env.$ENV\n"
elif [ -f ".env.local" ]; then
    printf "${YELLOW}⚠${NC} .env.$ENV not found, using .env.local\n"
    set -a
    source ".env.local"
    set +a
else
    printf "${YELLOW}⚠${NC} No .env file found. Using environment variables or defaults.\n"
fi

# Set environment variable
export NEXT_PUBLIC_ENV="$ENV"
printf "\n"

# Step 2: Validate environment configuration
printf "${YELLOW}Step 2: Validating environment configuration...${NC}\n"
if ! "$SCRIPT_DIR/validate-env.sh"; then
    printf "\n${RED}✗${NC} Environment validation failed.\n"
    printf "Please fix the configuration issues before running tests.\n"
    exit 1
fi
printf "\n"

# Step 3: Verify test environment connectivity
printf "${YELLOW}Step 3: Verifying test environment connectivity...${NC}\n"
if ! "$SCRIPT_DIR/test-env-connectivity.sh" "$ENV"; then
    printf "\n${RED}✗${NC} Connectivity check failed. Aborting tests.\n"
    exit 1
fi
printf "\n"

# Step 4: Run E2E tests
printf "${YELLOW}Step 4: Running E2E tests against $ENV environment...${NC}\n"

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    printf "${YELLOW}Installing dependencies...${NC}\n"
    npm install
fi

# Generate Run ID
export PW_RUN_ID="test-env-$(date +%Y%m%d-%H%M%S)"
printf "Run ID: ${BLUE}$PW_RUN_ID${NC}\n"

# Run Playwright tests with environment variables
printf "Executing Playwright E2E tests...\n"
NEXT_PUBLIC_ENV="$ENV" npm run test:e2e || true # Allow script to continue to aggregation

# Step 5: Aggregate Results
printf "${YELLOW}Step 5: Aggregating results...${NC}\n"
node "$SCRIPT_DIR/aggregate-results.js"
TEST_EXIT_CODE=$?

# Step 6: Post-test Cleanup (Optional)
printf "${YELLOW}Step 6: Post-test cleanup...${NC}\n"
# Add any cleanup logic here (e.g., resetting test data via API)
printf "Cleanup complete.\n"

if [ $TEST_EXIT_CODE -eq 0 ]; then
    printf "\n${GREEN}✅ E2E tests completed successfully!${NC}\n\n"
    printf "Test results are available in:\n"
    printf "  • HTML Report: playwright-report/$PW_RUN_ID/index.html\n"
    printf "  • Test Artifacts: test-results/$PW_RUN_ID/\n"
    exit 0
else
    printf "\n${RED}✗${NC} E2E tests failed.\n"
    printf "\nTo view test report:\n"
    printf "  npx playwright show-report playwright-report/$PW_RUN_ID\n"
    exit 1
fi

