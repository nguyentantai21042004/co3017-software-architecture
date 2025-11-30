#!/bin/bash

# Validate Environment Configuration
# This script validates that environment variables are properly set for the current environment

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

printf "${BLUE}Validating Environment Configuration${NC}\n"
printf "==========================================\n\n"

cd "$CLIENT_DIR"

# Determine environment
ENV="${NEXT_PUBLIC_ENV:-local}"
if [ -f ".env.local" ]; then
  ENV="local"
elif [ -f ".env.test" ]; then
  ENV="test"
elif [ -f ".env.staging" ]; then
  ENV="staging"
fi

printf "${YELLOW}Environment: $ENV${NC}\n\n"

# Required environment variables
REQUIRED_VARS=(
  "NEXT_PUBLIC_CONTENT_API_URL"
  "NEXT_PUBLIC_SCORING_API_URL"
  "NEXT_PUBLIC_LEARNER_API_URL"
  "NEXT_PUBLIC_ADAPTIVE_API_URL"
  "NEXT_PUBLIC_CLIENT_URL"
)

# Check if .env file exists for the environment
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
esac

if [ -n "$ENV_FILE" ] && [ ! -f "$ENV_FILE" ]; then
  printf "${YELLOW}⚠${NC} Environment file '$ENV_FILE' not found.\n"
  printf "Using default localhost URLs.\n\n"
fi

# Load environment variables if file exists
if [ -f "$ENV_FILE" ]; then
  printf "Loading environment from $ENV_FILE...\n"
  set -a
  source "$ENV_FILE"
  set +a
fi

# Validate each required variable
ALL_VALID=true
MISSING_VARS=()
INVALID_URLS=()

for var in "${REQUIRED_VARS[@]}"; do
  value="${!var}"
  
  if [ -z "$value" ]; then
    printf "${RED}✗${NC} $var is not set\n"
    MISSING_VARS+=("$var")
    ALL_VALID=false
  else
    # Basic URL validation
    if [[ ! "$value" =~ ^https?:// ]]; then
      printf "${RED}✗${NC} $var has invalid URL format: $value\n"
      INVALID_URLS+=("$var")
      ALL_VALID=false
    else
      printf "${GREEN}✓${NC} $var = $value\n"
    fi
  fi
done

printf "\n"

# Summary
if [ "$ALL_VALID" = true ]; then
  printf "${GREEN}✅ Environment configuration is valid!${NC}\n\n"
  
  # Show current configuration
  printf "Current API URLs:\n"
  printf "  Content:   ${NEXT_PUBLIC_CONTENT_API_URL:-http://localhost:8081}\n"
  printf "  Scoring:    ${NEXT_PUBLIC_SCORING_API_URL:-http://localhost:8082}\n"
  printf "  Learner:    ${NEXT_PUBLIC_LEARNER_API_URL:-http://localhost:8083}\n"
  printf "  Adaptive:   ${NEXT_PUBLIC_ADAPTIVE_API_URL:-http://localhost:8084/api/adaptive}\n"
  printf "  Client:     ${NEXT_PUBLIC_CLIENT_URL:-http://localhost:3000}\n"
  printf "\n"
  
  exit 0
else
  printf "${RED}✗${NC} Environment configuration has issues:\n"
  
  if [ ${#MISSING_VARS[@]} -gt 0 ]; then
    printf "\nMissing variables:\n"
    for var in "${MISSING_VARS[@]}"; do
      printf "  • $var\n"
    done
  fi
  
  if [ ${#INVALID_URLS[@]} -gt 0 ]; then
    printf "\nInvalid URLs:\n"
    for var in "${INVALID_URLS[@]}"; do
      printf "  • $var\n"
    done
  fi
  
  printf "\n"
  printf "To fix:\n"
  printf "  1. Copy the appropriate .env.example file:\n"
  printf "     cp .env.local.example .env.local\n"
  printf "  2. Update the values in .env.local\n"
  printf "  3. Run this script again to validate\n"
  printf "\n"
  
  exit 1
fi

