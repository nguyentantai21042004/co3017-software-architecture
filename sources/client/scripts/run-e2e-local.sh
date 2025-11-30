#!/bin/bash

# Run E2E Tests Locally
# This script verifies services are healthy, sets up test data, and runs E2E tests

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

printf "${BLUE}Running E2E Tests Locally${NC}\n"
printf "============================\n\n"

# Step 1: Validate environment configuration
printf "${YELLOW}Step 1: Validating environment configuration...${NC}\n"
if ! "$SCRIPT_DIR/validate-env.sh"; then
    printf "\n${YELLOW}⚠${NC} Environment validation failed, but continuing...\n"
    printf "Please ensure environment variables are set correctly.\n"
fi
printf "\n"

# Step 2: Verify services are running
printf "${YELLOW}Step 2: Verifying backend services...${NC}\n"
if ! "$SCRIPT_DIR/verify-services.sh"; then
    printf "\n${RED}✗${NC} Services are not healthy. Please start services first:\n"
    printf "  ./scripts/start-services.sh\n"
    exit 1
fi
printf "\n"

# Step 3: Setup test data (optional - skip if already set up)
printf "${YELLOW}Step 3: Setting up test data...${NC}\n"
read -p "Do you want to setup/refresh test data? [y/N] " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    if ! "$SCRIPT_DIR/setup-test-data.sh"; then
        printf "${RED}✗${NC} Failed to setup test data\n"
        exit 1
    fi
else
    printf "${YELLOW}⚠${NC} Skipping test data setup (assuming data already exists)\n"
fi
printf "\n"

# Step 4: Run E2E tests
printf "${YELLOW}Step 4: Running E2E tests...${NC}\n"
cd "$CLIENT_DIR"

# Check if node_modules exists
if [ ! -d "node_modules" ]; then
    printf "${YELLOW}Installing dependencies...${NC}\n"
    npm install
fi

# Run Playwright tests
printf "Executing Playwright E2E tests...\n"
if npm run test:e2e; then
    printf "\n${GREEN}✅ E2E tests completed successfully!${NC}\n\n"
    printf "Test results are available in:\n"
    printf "  • HTML Report: playwright-report/index.html\n"
    printf "  • Test Artifacts: test-results/\n"
    exit 0
else
    printf "\n${RED}✗${NC} E2E tests failed. Check the output above for details.\n"
    printf "\nTo view test report:\n"
    printf "  npx playwright show-report\n"
    exit 1
fi

