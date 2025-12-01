# Client E2E Testing Makefile
# Convenient commands for running Playwright E2E tests

.PHONY: help test test-ui test-dashboard test-all clean-cache services-start services-stop services-verify setup-data cleanup-data

# Default target
help:
	@echo "Client E2E Testing Commands"
	@echo "============================"
	@echo ""
	@echo "Testing:"
	@echo "  make test              - Run all E2E tests"
	@echo "  make test-ui           - Run E2E tests in UI mode"
	@echo "  make test-dashboard    - Run only dashboard tests"
	@echo "  make test-debug        - Run tests with debug output"
	@echo ""
	@echo "Services:"
	@echo "  make services-start    - Start all backend services"
	@echo "  make services-stop     - Stop all backend services"
	@echo "  make services-verify   - Verify services are healthy"
	@echo ""
	@echo "Data Management:"
	@echo "  make setup-data        - Setup test data in databases"
	@echo "  make cleanup-data      - Clean up test data"
	@echo ""
	@echo "Maintenance:"
	@echo "  make clean-cache       - Clear Next.js cache"
	@echo "  make clean-all         - Clean cache and stop services"
	@echo ""

# Run all E2E tests
test:
	@echo "Running all E2E tests..."
	npm run test:e2e

# Run E2E tests in UI mode
test-ui:
	@echo "Opening Playwright UI..."
	npm run test:e2e:ui

# Run only dashboard tests
test-dashboard:
	@echo "Running dashboard tests..."
	npx playwright test e2e/dashboard.spec.ts

# Run only learning flow tests
test-learning:
	@echo "Running learning flow tests..."
	npx playwright test e2e/learning-flow.spec.ts

# Run only API integration tests
test-api:
	@echo "Running API integration tests..."
	npx playwright test e2e/api-integration.spec.ts

# Run only UI component tests
test-components:
	@echo "Running UI component tests..."
	npx playwright test e2e/ui-components.spec.ts

# Run only error handling tests
test-error:
	@echo "Running error handling tests..."
	npx playwright test e2e/error-handling.spec.ts

# Run only test data verification tests
test-data:
	@echo "Running test data verification..."
	npx playwright test e2e/test-data.spec.ts

# Run Antigravity Browser features tests
test-antigravity:
	@echo "Running Antigravity Browser features tests..."
	npx playwright test e2e/antigravity-features.spec.ts

# Run tests with debug output
test-debug:
	@echo "Running tests with debug output..."
	DEBUG=pw:api npx playwright test

# Run tests with services (full integration)
test-with-services:
	@echo "Running E2E tests with automatic service management..."
	npm run test:e2e:with-services

# Start backend services
services-start:
	@echo "Starting backend services..."
	npm run e2e:start-services

# Stop backend services
services-stop:
	@echo "Stopping backend services..."
	npm run e2e:stop-services

# Verify services are healthy
services-verify:
	@echo "Verifying service health..."
	npm run e2e:verify-services

# Setup test data
setup-data:
	@echo "Setting up test data..."
	npm run e2e:setup-data

# Cleanup test data
cleanup-data:
	@echo "Cleaning up test data..."
	npm run e2e:cleanup-data

# Clear Next.js cache
clean-cache:
	@echo "Clearing Next.js cache..."
	rm -rf .next
	@echo "✓ Cache cleared"

# Clean everything
clean-all: clean-cache services-stop
	@echo "✓ All cleaned up"

# Quick test cycle (clean + test)
quick-test: clean-cache test-dashboard
	@echo "✓ Quick test completed"

# Full test cycle (clean + services + test)
full-test: clean-cache services-start setup-data test services-stop
	@echo "✓ Full test cycle completed"
