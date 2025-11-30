#!/bin/bash

# Load Environment Variables
# This script loads environment variables from the appropriate .env file
# Usage: source ./scripts/load-env.sh [environment]
# Or: eval $(./scripts/load-env.sh [environment])

set -e

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CLIENT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

cd "$CLIENT_DIR"

# Determine environment
ENV="${1:-local}"
if [ -z "$1" ]; then
    # Auto-detect from existing files
    if [ -f ".env.local" ]; then
        ENV="local"
    elif [ -f ".env.test" ]; then
        ENV="test"
    elif [ -f ".env.staging" ]; then
        ENV="staging"
    fi
fi

# Determine .env file
ENV_FILE=""
case $ENV in
    local)
        ENV_FILE=".env.local"
        ;;
    test)
        ENV_FILE=".env.test"
        ;;
    staging)
        ENV_FILE=".env.staging"
        ;;
    *)
        echo "Error: Invalid environment '$ENV'. Use: local, test, staging" >&2
        exit 1
        ;;
esac

# Check if file exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Warning: Environment file '$ENV_FILE' not found. Using defaults." >&2
    exit 0
fi

# Export variables (for sourcing)
if [ "${BASH_SOURCE[0]}" != "${0}" ]; then
    # Script is being sourced
    set -a
    source "$ENV_FILE"
    set +a
    echo "Loaded environment from $ENV_FILE"
else
    # Script is being executed
    # Output export statements for eval
    while IFS= read -r line || [ -n "$line" ]; do
        # Skip comments and empty lines
        if [[ "$line" =~ ^[[:space:]]*# ]] || [[ -z "$line" ]]; then
            continue
        fi
        # Export the variable
        echo "export $line"
    done < "$ENV_FILE"
fi

