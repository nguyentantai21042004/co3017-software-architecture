#!/bin/bash

# Reset Test Environment
# This script provides procedures for resetting test environment data
# Note: This is a template - actual implementation depends on test environment setup

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

printf "${BLUE}Reset Test Environment${NC}\n"
printf "=======================\n\n"

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
    printf "${YELLOW}⚠${NC} No .env file found.\n"
fi

printf "\n"
printf "${YELLOW}⚠${NC} WARNING: This script is a template for test environment reset procedures.\n"
printf "Actual implementation depends on your test environment setup.\n\n"

printf "Test environment reset procedures should include:\n"
printf "  1. Connect to test environment databases\n"
printf "  2. Clean up test user data\n"
printf "  3. Reset test questions if needed\n"
printf "  4. Reset mastery scores for test users\n"
printf "  5. Verify data cleanup was successful\n"
printf "\n"

# Example procedures (commented out - adapt for your test environment)
printf "${BLUE}Example Reset Procedures:${NC}\n"
printf "\n"
printf "# Connect to Content Service database and clean test questions:\n"
printf "# psql -h test-content-db-host -U postgres -d content_db -c \"DELETE FROM questions WHERE skill_tag IN ('math', 'science');\"\n"
printf "\n"
printf "# Connect to Learner Model Service database and reset test user mastery:\n"
printf "# psql -h test-learner-db-host -U postgres -d learner_db -c \"DELETE FROM skill_mastery WHERE user_id LIKE 'test-%';\"\n"
printf "\n"
printf "# Connect to Scoring Service database and clean test submissions:\n"
printf "# psql -h test-scoring-db-host -U postgres -d scoring_db -c \"DELETE FROM submissions WHERE user_id LIKE 'test-%';\"\n"
printf "\n"

# Confirmation prompt
read -p "${YELLOW}Do you want to proceed with test environment reset? [y/N] ${NC}" -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    printf "${BLUE}Reset cancelled.${NC}\n"
    exit 0
fi

printf "\n"
printf "${YELLOW}⚠${NC} This script needs to be customized for your test environment.\n"
printf "Please implement the actual reset procedures based on your test environment setup.\n"
printf "\n"
printf "For local environment, you can use:\n"
printf "  ./scripts/cleanup-test-data.sh\n"
printf "\n"

exit 0

