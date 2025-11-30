#!/bin/bash

# Run Playwright E2E Tests with Automatic Service Management
# This script runs Playwright tests with automatic backend service startup and teardown
# It's a convenience wrapper that ensures services are ready before tests run

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

printf "${BLUE}Run Playwright E2E Tests with Service Management${NC}\n"
printf "==================================================\n\n"

cd "$CLIENT_DIR"

# Determine environment
ENV="${NEXT_PUBLIC_ENV:-local}"

printf "${YELLOW}Environment: $ENV${NC}\n\n"

# Check if we should skip service teardown
SKIP_TEARDOWN="${SKIP_SERVICE_TEARDOWN:-false}"
if [ "$SKIP_TEARDOWN" = "true" ]; then
    printf "${YELLOW}Note: Services will NOT be stopped after tests (SKIP_SERVICE_TEARDOWN=true)${NC}\n\n"
fi

# For local environment, Playwright global setup will handle service startup
# For test/staging, services should already be deployed
if [ "$ENV" = "local" ]; then
    printf "${BLUE}Local environment detected.${NC}\n"
    printf "Playwright global setup will automatically:\n"
    printf "  1. Start backend services\n"
    printf "  2. Verify services are healthy\n"
    printf "  3. Setup test data\n"
    printf "\n"
    printf "To keep services running after tests, set:\n"
    printf "  SKIP_SERVICE_TEARDOWN=true npm run test:e2e\n"
    printf "\n"
else
    printf "${BLUE}Test/Staging environment detected.${NC}\n"
    printf "Assuming services are already deployed.\n"
    printf "Make sure services are accessible before running tests.\n"
    printf "\n"
    
    # Verify test environment connectivity
    if [ -f "$SCRIPT_DIR/verify-test-env.sh" ]; then
        printf "Verifying test environment connectivity...\n"
        if ! "$SCRIPT_DIR/verify-test-env.sh" "$ENV"; then
            printf "${YELLOW}⚠${NC} Test environment verification failed, but continuing...\n"
            printf "Tests may fail if services are not accessible.\n"
        fi
        printf "\n"
    fi
fi

# Run Playwright tests
printf "${BLUE}Running Playwright E2E tests...${NC}\n"
printf "\n"

# Pass through all arguments to Playwright
if npm run test:e2e -- "$@"; then
    printf "\n${GREEN}✅ Playwright E2E tests completed successfully!${NC}\n\n"
    printf "Test results:\n"
    printf "  • HTML Report: playwright-report/index.html\n"
    printf "  • Test Artifacts: test-results/\n"
    printf "\n"
    if [ "$ENV" = "local" ] && [ "$SKIP_TEARDOWN" != "true" ]; then
        printf "Note: Services will be stopped by Playwright global teardown.\n"
        printf "To keep services running, set SKIP_SERVICE_TEARDOWN=true\n"
    fi
    exit 0
else
    printf "\n${RED}✗${NC} Playwright E2E tests failed.\n"
    printf "\nTo view test report:\n"
    printf "  npx playwright show-report\n"
    printf "\n"
    exit 1
fi

