#!/bin/bash

# Stop All Microservices Script
# This script stops all running microservices

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🛑 Stopping All Microservices${NC}"
echo "=============================="
echo ""

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Function to stop a service
stop_service() {
    local service_name=$1
    local service_dir=$2
    local pid_file="$SCRIPT_DIR/$service_dir/.service.pid"
    
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        
        if ps -p $pid > /dev/null 2>&1; then
            echo -e "${YELLOW}Stopping $service_name (PID: $pid)...${NC}"
            kill $pid 2>/dev/null || true
            
            # Wait for process to stop
            local count=0
            while ps -p $pid > /dev/null 2>&1 && [ $count -lt 10 ]; do
                sleep 1
                count=$((count + 1))
            done
            
            # Force kill if still running
            if ps -p $pid > /dev/null 2>&1; then
                echo -e "${YELLOW}Force stopping $service_name...${NC}"
                kill -9 $pid 2>/dev/null || true
            fi
            
            echo -e "${GREEN}✓${NC} $service_name stopped"
        else
            echo -e "${YELLOW}⚠${NC}  $service_name is not running (PID $pid not found)"
        fi
        
        rm -f "$pid_file"
    else
        echo -e "${YELLOW}⚠${NC}  No PID file found for $service_name"
    fi
    echo ""
}

# Stop all services
stop_service "Content Service" "content"
stop_service "Scoring Service" "scoring"
stop_service "Learner Model Service" "learner-model"
stop_service "Adaptive Engine" "adaptive-engine"

# Also try to kill by port (backup method)
echo "🔍 Checking for any remaining processes on service ports..."

for port in 8081 8082 8083 8084; do
    pid=$(lsof -ti:$port 2>/dev/null || true)
    if [ ! -z "$pid" ]; then
        echo -e "${YELLOW}Found process on port $port (PID: $pid), killing...${NC}"
        kill -9 $pid 2>/dev/null || true
    fi
done

echo "=============================="
echo -e "${GREEN}✅ All services stopped!${NC}"
echo ""
