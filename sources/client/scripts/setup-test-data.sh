#!/bin/bash

# Setup Test Data for E2E Testing
# This script initializes test data for:
# 1. Content Service (questions in content_db)
# 2. Learner Model Service (test user mastery data in learner_db)

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the project root directory (sources/)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SOURCES_DIR="$SCRIPT_DIR/sources"
CLIENT_SCRIPTS_DIR="$SCRIPT_DIR/sources/client/scripts"

printf "${BLUE}Setting up E2E Test Data${NC}\n"
printf "============================\n\n"

# Database configuration
DB_HOST="${POSTGRES_HOST:-localhost}"
DB_PORT="${POSTGRES_PORT:-5432}"
DB_USER="${POSTGRES_USER:-postgres}"
DB_PASSWORD="${POSTGRES_PASSWORD:-postgres}"

# Test user ID (matches E2E test setup)
TEST_USER_ID="test-user-123"

# ============================================================================
# Step 1: Setup Content Service Test Data
# ============================================================================

printf "${YELLOW}Step 1: Setting up Content Service test data...${NC}\n"

CONTENT_DB="content_db"
SQL_FILE="$CLIENT_SCRIPTS_DIR/insert_e2e_test_data.sql"

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
    printf "${RED}✗${NC} SQL file not found: $SQL_FILE\n"
    exit 1
fi

# Check PostgreSQL connection
printf "Checking PostgreSQL connection to $CONTENT_DB...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -c "SELECT 1" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Cannot connect to PostgreSQL database '$CONTENT_DB'\n"
    printf "\nPlease ensure:\n"
    printf "  1. PostgreSQL is running\n"
    printf "  2. Database '$CONTENT_DB' exists\n"
    printf "  3. Services are started (run ./scripts/start-services.sh)\n"
    exit 1
fi
printf "${GREEN}✓${NC} PostgreSQL connection successful\n"

# Check if questions table exists
printf "Checking if questions table exists...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -c "\d questions" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Table 'questions' does not exist\n"
    printf "\nPlease start the Content Service first:\n"
    printf "  ./scripts/start-services.sh\n"
    exit 1
fi
printf "${GREEN}✓${NC} Table 'questions' exists\n"

# Insert test data
printf "Inserting test questions...\n"
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -f "$SQL_FILE" > /dev/null 2>&1; then
    printf "${GREEN}✓${NC} Test questions inserted successfully\n"
else
    printf "${YELLOW}⚠${NC} Some questions may already exist (this is OK)\n"
fi

# Verify inserted data
printf "Verifying test questions...\n"
QUESTION_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -t -c "SELECT COUNT(*) FROM questions WHERE skill_tag IN ('math', 'science');" | xargs)
printf "${GREEN}✓${NC} Found $QUESTION_COUNT test questions\n\n"

# ============================================================================
# Step 2: Setup Learner Model Service Test Data
# ============================================================================

printf "${YELLOW}Step 2: Setting up Learner Model Service test data...${NC}\n"

LEARNER_DB="learner_db"

# Check PostgreSQL connection
printf "Checking PostgreSQL connection to $LEARNER_DB...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB -c "SELECT 1" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Cannot connect to PostgreSQL database '$LEARNER_DB'\n"
    printf "\nPlease ensure:\n"
    printf "  1. PostgreSQL is running\n"
    printf "  2. Database '$LEARNER_DB' exists\n"
    printf "  3. Services are started (run ./scripts/start-services.sh)\n"
    exit 1
fi
printf "${GREEN}✓${NC} PostgreSQL connection successful\n"

# Check if skill_mastery table exists
printf "Checking if skill_mastery table exists...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB -c "\d skill_mastery" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Table 'skill_mastery' does not exist\n"
    printf "\nPlease start the Learner Model Service first:\n"
    printf "  ./scripts/start-services.sh\n"
    exit 1
fi
printf "${GREEN}✓${NC} Table 'skill_mastery' exists\n"

# Insert test user mastery data
printf "Inserting test user mastery data...\n"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB <<EOF
INSERT INTO skill_mastery (user_id, skill_tag, current_score, last_updated)
VALUES
    ('$TEST_USER_ID', 'math', 50, NOW()),
    ('$TEST_USER_ID', 'science', 60, NOW())
ON CONFLICT (user_id, skill_tag) 
DO UPDATE SET 
    current_score = EXCLUDED.current_score,
    last_updated = NOW();
EOF

if [ $? -eq 0 ]; then
    printf "${GREEN}✓${NC} Test user mastery data inserted successfully\n"
else
    printf "${RED}✗${NC} Failed to insert test user mastery data\n"
    exit 1
fi

# Verify inserted data
printf "Verifying test user mastery data...\n"
MASTERY_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $LEARNER_DB -t -c "SELECT COUNT(*) FROM skill_mastery WHERE user_id = '$TEST_USER_ID';" | xargs)
printf "${GREEN}✓${NC} Found $MASTERY_COUNT skill mastery records for test user\n\n"

# ============================================================================
# Summary
# ============================================================================

printf "${GREEN}✅ E2E test data setup complete!${NC}\n\n"
printf "Summary:\n"
printf "  • Content Service:\n"
printf "    - Math questions: $(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -t -c "SELECT COUNT(*) FROM questions WHERE skill_tag = 'math';" | xargs)\n"
printf "    - Science questions: $(PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $CONTENT_DB -t -c "SELECT COUNT(*) FROM questions WHERE skill_tag = 'science';" | xargs)\n"
printf "  • Learner Model Service:\n"
printf "    - Test user ID: $TEST_USER_ID\n"
printf "    - Skills: math (50%%), science (60%%)\n"
printf "\n"
printf "You can now run E2E tests:\n"
printf "  npm run test:e2e\n"
printf "\n"

