#!/bin/bash

# Start All Microservices Script
# This script starts all 4 microservices in the background with separate log files

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

printf "${BLUE}🚀 Starting All Microservices${NC}\n"
printf "==============================\n"
printf "\n"

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Define log files
CONTENT_LOG="$SCRIPT_DIR/content/content.log"
SCORING_LOG="$SCRIPT_DIR/scoring/scoring.log"
LEARNER_LOG="$SCRIPT_DIR/learner-model/learner.log"
ADAPTIVE_LOG="$SCRIPT_DIR/adaptive-engine/adaptive.log"

# Clean up old log files
printf "🧹 Cleaning up old log files...\n"
rm -f "$CONTENT_LOG" "$SCORING_LOG" "$LEARNER_LOG" "$ADAPTIVE_LOG"
printf "${GREEN}✓${NC} Log files cleaned\n"
printf "\n"

# Function to start a service
start_service() {
    local service_name=$1
    local service_dir=$2
    local log_file=$3
    local start_command=$4
    
    printf "${YELLOW}Starting $service_name...${NC}\n"
    
    cd "$SCRIPT_DIR/$service_dir"
    
    # Start the service in background and redirect output to log file
    nohup $start_command > "$log_file" 2>&1 &
    local pid=$!
    
    printf "${GREEN}✓${NC} $service_name started (PID: $pid)\n"
    printf "  Log: $log_file\n"
    printf "\n"
    
    # Save PID to file for later stopping
    echo $pid > "$SCRIPT_DIR/$service_dir/.service.pid"
}

# Start Content Service (Java/Spring Boot)
start_service "Content Service" "content" "$CONTENT_LOG" "mvn spring-boot:run"

# Wait a bit for Content Service to initialize
printf "⏳ Waiting 10 seconds for Content Service to initialize...\n"
sleep 10

# Start Scoring Service (Go)
start_service "Scoring Service" "scoring" "$SCORING_LOG" "go run cmd/api/main.go"

# Start Learner Model Service (Go)
start_service "Learner Model Service" "learner-model" "$LEARNER_LOG" "go run cmd/api/main.go"

# Start Adaptive Engine (Go)
start_service "Adaptive Engine" "adaptive-engine" "$ADAPTIVE_LOG" "go run cmd/api/main.go"

printf "==============================\n"
printf "${GREEN}✅ All services started!${NC}\n"
printf "\n"
printf "Services:\n"
printf "  • Content Service:      http://localhost:8081\n"
printf "  • Scoring Service:      http://localhost:8082\n"
printf "  • Learner Model:        http://localhost:8083\n"
printf "  • Adaptive Engine:      http://localhost:8084\n"
printf "\n"
printf "Logs:\n"
printf "  • Content:      $CONTENT_LOG\n"
printf "  • Scoring:      $SCORING_LOG\n"
printf "  • Learner:      $LEARNER_LOG\n"
printf "  • Adaptive:     $ADAPTIVE_LOG\n"
printf "\n"
printf "To view logs in real-time:\n"
printf "  ./scripts/view_logs.sh\n"
printf "\n"
printf "To stop all services:\n"
printf "  ./scripts/stop_services.sh\n"
printf "\n"
