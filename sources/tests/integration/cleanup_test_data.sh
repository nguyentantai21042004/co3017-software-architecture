#!/bin/bash

# Cleanup Test Data Script
# This script cleans up test data from previous test runs

set -e

printf "Cleaning up test data...\n"

# Database connection details
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="postgres"

# Clean learner_db (mastery data)
printf "  Cleaning learner_db...\n"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d learner_db -c "
DELETE FROM skill_mastery 
WHERE user_id LIKE 'test-user-%' 
   OR user_id LIKE 'integration-test-user-%';
" 2>/dev/null || printf "    Could not clean learner_db (may not be critical)\n"

# Clean scoring_db (submission data)  
printf "  Cleaning scoring_db...\n"
PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d scoring_db -c "
DELETE FROM submissions 
WHERE user_id LIKE 'test-user-%' 
   OR user_id LIKE 'integration-test-user-%';
" 2>/dev/null || printf "    Could not clean scoring_db (may not be critical)\n"

printf "Test data cleanup complete\n\n"
