#!/bin/sh

# =============================================================================
# Create databases if they don't exist
# This script can be run after container is started
# =============================================================================

set -e

POSTGRES_USER="${POSTGRES_USER:-postgres}"
POSTGRES_PASSWORD="${POSTGRES_PASSWORD:-postgres}"
POSTGRES_HOST="${POSTGRES_HOST:-localhost}"
POSTGRES_PORT="${POSTGRES_PORT:-5432}"

DATABASES="content_db scoring_db learner_db"

echo "Creating databases if they don't exist..."

for db in $DATABASES; do
    echo "Checking database: $db"
    
    # Check if database exists
    DB_EXISTS=$(PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -tAc "SELECT 1 FROM pg_database WHERE datname='$db'" postgres 2>/dev/null || echo "0")
    
    if [ "$DB_EXISTS" = "1" ]; then
        echo "  Database '$db' already exists, skipping..."
    else
        echo "  Creating database '$db'..."
        PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -c "CREATE DATABASE $db;" postgres
        PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" -U "$POSTGRES_USER" -c "GRANT ALL PRIVILEGES ON DATABASE $db TO $POSTGRES_USER;" postgres
        echo "  Database '$db' created successfully!"
    fi
done

echo "Done!"

