#!/bin/bash

# Intelligent Tutoring System - Demo Startup Script
# This script starts all services and prepares the system for demo

set -e  # Exit on error

echo "🚀 Starting Intelligent Tutoring System Demo..."
echo ""

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Step 1: Check prerequisites
echo -e "${BLUE}📋 Step 1/5: Checking prerequisites...${NC}"
if ! command -v docker &> /dev/null; then
    echo "❌ Docker not found. Please install Docker Desktop."
    exit 1
fi

if ! command -v node &> /dev/null; then
    echo "❌ Node.js not found. Please install Node.js 18+."
    exit 1
fi

if ! docker info &> /dev/null; then
    echo "❌ Docker daemon not running. Please start Docker Desktop."
    exit 1
fi

echo -e "${GREEN}✅ Prerequisites OK${NC}"
echo ""

# Step 2: Start backend services
echo -e "${BLUE}📦 Step 2/5: Starting backend services...${NC}"
cd sources/services

# Check if services are already running
if docker compose ps | grep -q "Up"; then
    echo -e "${YELLOW}⚠️  Services already running. Restarting...${NC}"
    docker compose down
fi

docker compose up -d

echo "⏳ Waiting for services to be healthy..."
sleep 10

# Wait for services to be ready
MAX_WAIT=60
ELAPSED=0
while [ $ELAPSED -lt $MAX_WAIT ]; do
    if docker compose ps | grep -q "healthy"; then
        echo -e "${GREEN}✅ Backend services ready${NC}"
        break
    fi
    sleep 5
    ELAPSED=$((ELAPSED + 5))
    echo "   Still waiting... (${ELAPSED}s/${MAX_WAIT}s)"
done

if [ $ELAPSED -ge $MAX_WAIT ]; then
    echo "❌ Services failed to start within ${MAX_WAIT}s"
    echo "Run 'docker compose logs' to see errors"
    exit 1
fi

# Show running services
echo ""
echo "📊 Running services:"
docker compose ps
echo ""

# Step 3: Setup test data
echo -e "${BLUE}📝 Step 3/5: Setting up test data...${NC}"
cd ../client

if [ ! -f "scripts/setup-test-data.sh" ]; then
    echo "❌ Test data script not found"
    exit 1
fi

chmod +x scripts/setup-test-data.sh
./scripts/setup-test-data.sh

echo -e "${GREEN}✅ Test data ready${NC}"
echo ""

# Step 4: Install frontend dependencies (if needed)
echo -e "${BLUE}📦 Step 4/5: Checking frontend dependencies...${NC}"
if [ ! -d "node_modules" ]; then
    echo "Installing dependencies (this may take a few minutes)..."
    npm install
else
    echo "Dependencies already installed"
fi
echo -e "${GREEN}✅ Frontend dependencies ready${NC}"
echo ""

# Step 5: Start frontend
echo -e "${BLUE}🌐 Step 5/5: Starting frontend...${NC}"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo -e "${GREEN}✅ All services started successfully!${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "📊 System Status:"
echo "   ✅ Backend Services: Running (Docker Compose)"
echo "   ✅ Test Data: Loaded (test-user-123)"
echo "   ✅ Frontend: Starting on http://localhost:3001"
echo ""
echo "🎯 Demo Ready!"
echo ""
echo "📖 Quick Start:"
echo "   1. Open browser to: http://localhost:3001"
echo "   2. Dashboard will auto-login as test-user-123"
echo "   3. Click 'Continue Learning' on any skill"
echo "   4. Answer questions and watch mastery update"
echo ""
echo "🧪 Run E2E Tests:"
echo "   npx playwright test --project=chromium"
echo ""
echo "📚 Full Demo Guide:"
echo "   See DEMO_GUIDE.md for detailed demo script"
echo ""
echo "🛑 To stop all services:"
echo "   cd sources/services && docker compose down"
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Starting Next.js development server..."
echo ""

# Start frontend (this will keep running)
npm run dev
