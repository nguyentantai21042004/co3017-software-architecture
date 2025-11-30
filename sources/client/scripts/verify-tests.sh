#!/bin/bash

# Script to verify test setup and run tests
# This script checks if all test infrastructure is properly configured

set -e

echo "🔍 Verifying Test Infrastructure..."
echo ""

# Check if dependencies are installed
echo "1. Checking dependencies..."
if [ ! -d "node_modules" ]; then
  echo "   ❌ node_modules not found. Run 'npm install' first."
  exit 1
fi
echo "   ✅ Dependencies installed"

# Check if Jest is available
echo "2. Checking Jest..."
if ! command -v npx jest &> /dev/null; then
  echo "   ❌ Jest not found"
  exit 1
fi
echo "   ✅ Jest available"

# Check if Playwright is available
echo "3. Checking Playwright..."
if ! command -v npx playwright &> /dev/null; then
  echo "   ❌ Playwright not found"
  exit 1
fi
echo "   ✅ Playwright available"

# Check if test files exist
echo "4. Checking test files..."
COMPONENT_TESTS=$(find __tests__ -name "*.test.ts" -o -name "*.test.tsx" | wc -l)
E2E_TESTS=$(find e2e -name "*.spec.ts" | wc -l)
echo "   ✅ Found $COMPONENT_TESTS component test files"
echo "   ✅ Found $E2E_TESTS E2E test files"

# Run component tests
echo ""
echo "5. Running component tests..."
npm test -- --passWithNoTests

echo ""
echo "✅ All test infrastructure verified!"
echo ""
echo "To run E2E tests, ensure:"
echo "  - Backend services are running (Content, Scoring, Learner Model, Adaptive Engine)"
echo "  - Next.js dev server can start (npm run dev)"
echo "  - Run: npm run test:e2e"

