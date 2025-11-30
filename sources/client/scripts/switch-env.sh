#!/bin/bash

# Switch Environment Configuration
# This script helps switch between different environment configurations
# by copying the appropriate .env.example file to .env.local

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

printf "${BLUE}Switch Environment Configuration${NC}\n"
printf "====================================\n\n"

cd "$CLIENT_DIR"

# Show usage if no argument provided
if [ $# -eq 0 ]; then
    printf "Usage: $0 <environment>\n"
    printf "\n"
    printf "Available environments:\n"
    printf "  local    - Local development (default)\n"
    printf "  test     - Test environment\n"
    printf "  staging  - Staging environment\n"
    printf "\n"
    printf "Examples:\n"
    printf "  $0 local\n"
    printf "  $0 test\n"
    printf "  $0 staging\n"
    printf "\n"
    exit 1
fi

ENV=$1

# Validate environment
case $ENV in
    local|test|staging)
        ;;
    *)
        printf "${RED}✗${NC} Invalid environment: $ENV\n"
        printf "Valid environments: local, test, staging\n"
        exit 1
        ;;
esac

# Determine source and target files
SOURCE_FILE=".env.$ENV.example"
TARGET_FILE=".env.local"

# Check if source file exists
if [ ! -f "$SOURCE_FILE" ]; then
    printf "${RED}✗${NC} Environment template not found: $SOURCE_FILE\n"
    printf "\n"
    printf "Available templates:\n"
    ls -1 .env.*.example 2>/dev/null || printf "  (none found)\n"
    exit 1
fi

# Confirm if .env.local already exists
if [ -f "$TARGET_FILE" ]; then
    printf "${YELLOW}⚠${NC} $TARGET_FILE already exists.\n"
    printf "This will overwrite the existing file.\n"
    read -p "Continue? [y/N] " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        printf "${YELLOW}Cancelled.${NC}\n"
        exit 0
    fi
fi

# Copy template to .env.local
printf "Copying $SOURCE_FILE to $TARGET_FILE...\n"
cp "$SOURCE_FILE" "$TARGET_FILE"

if [ $? -eq 0 ]; then
    printf "${GREEN}✓${NC} Environment configuration switched to: $ENV\n"
    printf "\n"
    printf "Configuration file: $TARGET_FILE\n"
    printf "\n"
    printf "${YELLOW}Note:${NC} Please review and update the values in $TARGET_FILE if needed.\n"
    printf "\n"
    printf "To validate the configuration:\n"
    printf "  npm run e2e:validate-env\n"
    printf "\n"
else
    printf "${RED}✗${NC} Failed to copy environment configuration\n"
    exit 1
fi

