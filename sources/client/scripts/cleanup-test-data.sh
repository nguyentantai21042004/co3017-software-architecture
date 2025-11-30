#!/bin/bash

# Cleanup E2E Test Data
# This script removes test data created for E2E testing
# Use this to reset the test environment to a clean state

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the project root directory (sources/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

printf "${BLUE}Cleaning up E2E Test Data${NC}\n"
printf "=============================\n\n"

# Database configuration
DB_HOST="${POSTGRES_HOST:-localhost}"
DB_PORT="${POSTGRES_PORT:-5432}"
DB_USER="${POSTGRES_USER:-postgres}"
DB_PASSWORD="${POSTGRES_PASSWORD:-postgres}"

# Test user ID
TEST_USER_ID="test-user-123"

# Confirmation prompt
printf "${YELLOW}This will delete all E2E test data:${NC}\n"
printf "  • Test questions (math, science) from Content Service\n"
printf "  • Test user mastery data ($TEST_USER_ID) from Learner Model Service\n"
printf "\n"
read -p "Are you sure you want to continue? [y/N] " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    printf "${YELLOW}Cancelled.${NC}\n"
    exit 0
fi

# ============================================================================
# Step 1: Cleanup Content Service Test Data
# ============================================================================

printf "\n${YELLOW}Step 1: Cleaning up Content Service test data...${NC}\n"

CONTENT_DB="content_db"

# Check PostgreSQL connection
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -c "SELECT 1" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Cannot connect to PostgreSQL database '$CONTENT_DB'\n"
    exit 1
fi

# Delete test questions
printf "Deleting test questions...\n"
DELETED_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -t -c "DELETE FROM questions WHERE skill_tag IN ('math', 'science') RETURNING id;" | wc -l | xargs)
printf "${GREEN}✓${NC} Deleted $DELETED_COUNT test questions\n\n"

# ============================================================================
# Step 2: Cleanup Learner Model Service Test Data
# ============================================================================

printf "${YELLOW}Step 2: Cleaning up Learner Model Service test data...${NC}\n"

LEARNER_DB="learner_db"

# Check PostgreSQL connection
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB -c "SELECT 1" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Cannot connect to PostgreSQL database '$LEARNER_DB'\n"
    exit 1
fi

# Delete test user mastery data
printf "Deleting test user mastery data...\n"
DELETED_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB -t -c "DELETE FROM skill_mastery WHERE user_id = '$TEST_USER_ID' RETURNING user_id;" | wc -l | xargs)
printf "${GREEN}✓${NC} Deleted $DELETED_COUNT skill mastery records\n\n"

# ============================================================================
# Summary
# ============================================================================

printf "${GREEN}✅ E2E test data cleanup complete!${NC}\n\n"
printf "To set up test data again, run:\n"
printf "  ./scripts/setup-test-data.sh\n"
printf "\n"

