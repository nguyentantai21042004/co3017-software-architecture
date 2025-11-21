#!/bin/bash

# View Logs Script
# This script shows all service logs in real-time

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Define log files
CONTENT_LOG="$SCRIPT_DIR/content/content.log"
SCORING_LOG="$SCRIPT_DIR/scoring/scoring.log"
LEARNER_LOG="$SCRIPT_DIR/learner-model/learner.log"
ADAPTIVE_LOG="$SCRIPT_DIR/adaptive-engine/adaptive.log"

printf "${BLUE}📋 Service Logs Viewer${NC}\n"
printf "======================\n\n"
printf "${GREEN}Viewing all service logs...${NC}\n"
printf "Press Ctrl+C to exit\n\n"

# View all logs
tail -f "$CONTENT_LOG" "$SCORING_LOG" "$LEARNER_LOG" "$ADAPTIVE_LOG"

