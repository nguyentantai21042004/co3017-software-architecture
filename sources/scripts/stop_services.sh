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

printf "${BLUE}Stopping All Microservices${NC}\n"
printf "==============================\n"
printf "\n"

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Function to stop a service
stop_service() {
    local service_name=$1
    local service_dir=$2
    local pid_file_name=${3:-.service.pid}
    local pid_file="$SCRIPT_DIR/$service_dir/$pid_file_name"
    
    if [ -f "$pid_file" ]; then
        local pid=$(cat "$pid_file")
        
        if ps -p $pid > /dev/null 2>&1; then
            printf "${YELLOW}Stopping $service_name (PID: $pid)...${NC}\n"
            kill $pid 2>/dev/null || true
            
            # Wait for process to stop
            local count=0
            while ps -p $pid > /dev/null 2>&1 && [ $count -lt 10 ]; do
                sleep 1
                count=$((count + 1))
            done
            
            # Force kill if still running
            if ps -p $pid > /dev/null 2>&1; then
                printf "${YELLOW}Force stopping $service_name...${NC}\n"
                kill -9 $pid 2>/dev/null || true
            fi
            
            printf "${GREEN}$service_name stopped${NC}\n"
        else
            printf "${YELLOW}$service_name is not running (PID $pid not found)${NC}\n"
        fi
        
        rm -f "$pid_file"
    else
        printf "${YELLOW}No PID file found for $service_name${NC}\n"
    fi
    printf "\n"
}

# Stop all services
stop_service "Content Service" "content"
stop_service "Scoring Service" "scoring"
stop_service "Learner Model Service" "learner-model"
stop_service "Learner Model Consumer" "learner-model" ".consumer.pid"
stop_service "Adaptive Engine" "adaptive-engine"

# Also try to kill by port (backup method)
printf "Checking for any remaining processes on service ports...\n"

for port in 8081 8082 8083 8084; do
    pid=$(lsof -ti:$port 2>/dev/null || true)
    if [ ! -z "$pid" ]; then
        printf "${YELLOW}Found process on port $port (PID: $pid), killing...${NC}\n"
        kill -9 $pid 2>/dev/null || true
    fi
done

printf "==============================\n"
printf "${GREEN}All services stopped!${NC}\n"
printf "\n"
