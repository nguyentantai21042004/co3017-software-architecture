#!/bin/bash

# View Logs Script
# This script provides easy access to view service logs

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Get the src directory (parent of scripts directory)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Define log files
CONTENT_LOG="$SCRIPT_DIR/content/content.log"
SCORING_LOG="$SCRIPT_DIR/scoring/scoring.log"
LEARNER_LOG="$SCRIPT_DIR/learner-model/learner.log"
ADAPTIVE_LOG="$SCRIPT_DIR/adaptive-engine/adaptive.log"

echo -e "${BLUE}📋 Service Logs Viewer${NC}"
echo "======================"
echo ""
echo "Select a service to view logs:"
echo "  1) Content Service"
echo "  2) Scoring Service"
echo "  3) Learner Model Service"
echo "  4) Adaptive Engine"
echo "  5) All services (split screen)"
echo "  q) Quit"
echo ""
read -p "Enter choice: " choice

case $choice in
    1)
        echo -e "${GREEN}Viewing Content Service logs...${NC}"
        echo "Press Ctrl+C to exit"
        tail -f "$CONTENT_LOG"
        ;;
    2)
        echo -e "${GREEN}Viewing Scoring Service logs...${NC}"
        echo "Press Ctrl+C to exit"
        tail -f "$SCORING_LOG"
        ;;
    3)
        echo -e "${GREEN}Viewing Learner Model Service logs...${NC}"
        echo "Press Ctrl+C to exit"
        tail -f "$LEARNER_LOG"
        ;;
    4)
        echo -e "${GREEN}Viewing Adaptive Engine logs...${NC}"
        echo "Press Ctrl+C to exit"
        tail -f "$ADAPTIVE_LOG"
        ;;
    5)
        echo -e "${GREEN}Viewing all service logs...${NC}"
        echo "Press Ctrl+C to exit"
        tail -f "$CONTENT_LOG" "$SCORING_LOG" "$LEARNER_LOG" "$ADAPTIVE_LOG"
        ;;
    q|Q)
        echo "Exiting..."
        exit 0
        ;;
    *)
        echo -e "${YELLOW}Invalid choice${NC}"
        exit 1
        ;;
esac
