#!/bin/bash

# Stop All Backend Services for E2E Testing
# This script stops all backend services and infrastructure

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

printf "${BLUE}Stopping Backend Services${NC}\n"
printf "============================\n\n"

# Use docker compose (v2) if available, otherwise docker-compose (v1)
if docker compose version > /dev/null 2>&1; then
    DOCKER_COMPOSE="docker compose"
else
    DOCKER_COMPOSE="docker-compose"
fi

cd "$SOURCES_DIR"

# Stop application services
printf "${YELLOW}Stopping application services...${NC}\n"
$DOCKER_COMPOSE -f docker-compose.yml down

# Stop infrastructure services
printf "${YELLOW}Stopping infrastructure services...${NC}\n"
$DOCKER_COMPOSE -f docker-compose.infra.yml down

printf "\n${GREEN}✅ All services stopped!${NC}\n\n"

