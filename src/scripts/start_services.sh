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

echo -e "${BLUE}🚀 Starting All Microservices${NC}"
echo "=============================="
echo ""

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Define log files
CONTENT_LOG="$SCRIPT_DIR/content/content.log"
SCORING_LOG="$SCRIPT_DIR/scoring/scoring.log"
LEARNER_LOG="$SCRIPT_DIR/learner-model/learner.log"
ADAPTIVE_LOG="$SCRIPT_DIR/adaptive-engine/adaptive.log"

# Clean up old log files
echo "🧹 Cleaning up old log files..."
rm -f "$CONTENT_LOG" "$SCORING_LOG" "$LEARNER_LOG" "$ADAPTIVE_LOG"
echo -e "${GREEN}✓${NC} Log files cleaned"
echo ""

# Function to start a service
start_service() {
    local service_name=$1
    local service_dir=$2
    local log_file=$3
    local start_command=$4
    
    echo -e "${YELLOW}Starting $service_name...${NC}"
    
    cd "$SCRIPT_DIR/$service_dir"
    
    # Start the service in background and redirect output to log file
    nohup $start_command > "$log_file" 2>&1 &
    local pid=$!
    
    echo -e "${GREEN}✓${NC} $service_name started (PID: $pid)"
    echo "  Log: $log_file"
    echo ""
    
    # Save PID to file for later stopping
    echo $pid > "$SCRIPT_DIR/$service_dir/.service.pid"
}

# Start Content Service (Java/Spring Boot)
start_service "Content Service" "content" "$CONTENT_LOG" "mvn spring-boot:run"

# Wait a bit for Content Service to initialize
echo "⏳ Waiting 10 seconds for Content Service to initialize..."
sleep 10

# Start Scoring Service (Go)
start_service "Scoring Service" "scoring" "$SCORING_LOG" "go run cmd/api/main.go"

# Start Learner Model Service (Go)
start_service "Learner Model Service" "learner-model" "$LEARNER_LOG" "go run cmd/api/main.go"

# Start Adaptive Engine (Go)
start_service "Adaptive Engine" "adaptive-engine" "$ADAPTIVE_LOG" "go run cmd/api/main.go"

echo "=============================="
echo -e "${GREEN}✅ All services started!${NC}"
echo ""
echo "Services:"
echo "  • Content Service:      http://localhost:8081"
echo "  • Scoring Service:      http://localhost:8082"
echo "  • Learner Model:        http://localhost:8083"
echo "  • Adaptive Engine:      http://localhost:8084"
echo ""
echo "Logs:"
echo "  • Content:      $CONTENT_LOG"
echo "  • Scoring:      $SCORING_LOG"
echo "  • Learner:      $LEARNER_LOG"
echo "  • Adaptive:     $ADAPTIVE_LOG"
echo ""
echo "To view logs in real-time:"
echo "  tail -f $CONTENT_LOG"
echo "  tail -f $SCORING_LOG"
echo "  tail -f $LEARNER_LOG"
echo "  tail -f $ADAPTIVE_LOG"
echo ""
echo "To stop all services:"
echo "  ./stop_services.sh"
echo ""
