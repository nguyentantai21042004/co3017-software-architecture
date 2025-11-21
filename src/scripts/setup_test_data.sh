#!/bin/bash

# Setup Test Data for Integration Tests
# This script inserts test questions into the Content Service database

set -e

printf "🔧 Setting up test data for integration tests\n"
printf "==============================================\n\n"

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# Database configuration
DB_HOST="${POSTGRES_HOST:-localhost}"
DB_PORT="${POSTGRES_PORT:-5432}"
DB_USER="${POSTGRES_USER:-postgres}"
DB_PASSWORD="${POSTGRES_PASSWORD:-postgres}"
DB_NAME="content_db"

# Get script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SQL_FILE="$SCRIPT_DIR/insert_test_data.sql"

# Check if SQL file exists
if [ ! -f "$SQL_FILE" ]; then
    printf "${RED}✗${NC} SQL file not found: $SQL_FILE\n"
    exit 1
fi

printf "${YELLOW}Database Configuration:${NC}\n"
printf "  Host: $DB_HOST\n"
printf "  Port: $DB_PORT\n"
printf "  Database: $DB_NAME\n"
printf "  User: $DB_USER\n\n"

# Check if PostgreSQL is accessible
printf "Checking PostgreSQL connection...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "SELECT 1" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Cannot connect to PostgreSQL\n"
    printf "\nPlease ensure:\n"
    printf "  1. PostgreSQL is running\n"
    printf "  2. Database '$DB_NAME' exists\n"
    printf "  3. Credentials are correct\n"
    exit 1
fi
printf "${GREEN}✓${NC} PostgreSQL connection successful\n\n"

# Check if questions table exists
printf "Checking if questions table exists...\n"
if ! PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "\d questions" > /dev/null 2>&1; then
    printf "${RED}✗${NC} Table 'questions' does not exist\n"
    printf "\nPlease run the Content Service first to create tables:\n"
    printf "  cd src && ./scripts/start_services.sh\n"
    exit 1
fi
printf "${GREEN}✓${NC} Table 'questions' exists\n\n"

# Clean up existing test data (optional)
printf "Cleaning up existing test data...\n"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "DELETE FROM questions WHERE skill_tag IN ('math', 'science');" > /dev/null 2>&1
printf "${GREEN}✓${NC} Cleaned up old test data\n\n"

# Insert test data
printf "Inserting test data from $SQL_FILE...\n"
if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$SQL_FILE" > /dev/null 2>&1; then
    printf "${GREEN}✓${NC} Test data inserted successfully\n\n"
else
    printf "${RED}✗${NC} Failed to insert test data\n"
    exit 1
fi

# Verify inserted data
printf "Verifying inserted data...\n"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
SELECT 
    skill_tag, 
    is_remedial, 
    COUNT(*) as question_count 
FROM questions 
WHERE skill_tag IN ('math', 'science')
GROUP BY skill_tag, is_remedial 
ORDER BY skill_tag, is_remedial;
"

printf "\n${GREEN}✅ Test data setup complete!${NC}\n\n"
printf "Summary:\n"
printf "  • Math remedial questions: 5\n"
printf "  • Math standard questions: 5\n"
printf "  • Science remedial questions: 5\n"
printf "  • Science standard questions: 5\n"
printf "  • Total: 20 questions\n\n"

printf "You can now run integration tests:\n"
printf "  cd src/tests/integration && ./run_tests.sh\n\n"
