# Testing Guide

This document provides comprehensive instructions for running tests in the client application, including unit tests, component tests, and end-to-end (E2E) tests.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Unit and Component Tests](#unit-and-component-tests)
- [E2E Testing Setup](#e2e-testing-setup)
- [Running E2E Tests](#running-e2e-tests)
- [Service Management](#service-management)
- [Test Data Management](#test-data-management)
- [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Software

- **Node.js**: Version 18 or higher
- **npm**: Version 9 or higher
- **Docker Desktop**: Version 4.0 or higher (for E2E tests)
- **Docker Compose**: Included with Docker Desktop

### Required Services

For E2E tests, the following services must be running:
- PostgreSQL (port 5432)
- RabbitMQ (port 5672, management UI on 15672)
- Content Service (port 8081)
- Scoring Service (port 8082)
- Learner Model Service (port 8083)
- Adaptive Engine Service (port 8084)

## Unit and Component Tests

### Running Tests

```bash
# Run all tests
npm test

# Run tests in watch mode
npm run test:watch

# Run tests with coverage
npm run test:coverage
```

### Test Structure

- **Unit Tests**: `__tests__/lib/` - Tests for utility functions
- **Component Tests**: `__tests__/components/` - Tests for React components
- **Service Tests**: `__tests__/services/` - Tests for API service layer

### Writing Tests

Tests use Jest and React Testing Library. Example:

```typescript
import { render, screen } from '@testing-library/react'
import { MyComponent } from '@/components/MyComponent'

test('renders component', () => {
  render(<MyComponent />)
  expect(screen.getByText('Hello')).toBeInTheDocument()
})
```

## E2E Testing Setup

### Quick Start

1. **Start all backend services:**
   ```bash
   npm run e2e:start-services
   # Or manually:
   ./scripts/start-services.sh
   ```

2. **Setup test data:**
   ```bash
   npm run e2e:setup-data
   # Or manually:
   ./scripts/setup-test-data.sh
   ```

3. **Run E2E tests:**
   ```bash
   npm run e2e:run-local
   # Or manually:
   ./scripts/run-e2e-local.sh
   ```

### Service Startup Order

Services must be started in the following order:

1. **Infrastructure Services** (started automatically):
   - PostgreSQL
   - RabbitMQ

2. **Application Services** (started automatically):
   - Content Service (depends on PostgreSQL)
   - Scoring Service (depends on PostgreSQL, RabbitMQ)
   - Learner Model API (depends on PostgreSQL)
   - Learner Model Consumer (depends on PostgreSQL, RabbitMQ)
   - Adaptive Engine (depends on Content Service, Learner Model API)

### Service Dependencies

```
Client Application
    │
    ├─► Adaptive Engine Service (8084)
    │   ├─► Content Service (8081)
    │   │   └─► PostgreSQL (content_db)
    │   └─► Learner Model Service (8083)
    │       └─► PostgreSQL (learner_db)
    │
    ├─► Scoring Service (8082)
    │   ├─► PostgreSQL (scoring_db)
    │   └─► RabbitMQ
    │
    └─► Learner Model Service (8083)
        └─► PostgreSQL (learner_db)
```

## Running E2E Tests

### Using npm Scripts

```bash
# Complete workflow (recommended)
npm run e2e:run-local

# Individual steps
npm run e2e:start-services    # Start all backend services
npm run e2e:verify-services    # Verify services are healthy
npm run e2e:setup-data         # Setup test data
npm run test:e2e               # Run Playwright tests
npm run e2e:stop-services      # Stop all services
```

### Using Scripts Directly

```bash
# Start services
./scripts/start-services.sh

# Verify services
./scripts/verify-services.sh

# Setup test data
./scripts/setup-test-data.sh

# Run E2E tests
./scripts/run-e2e-local.sh

# Stop services
./scripts/stop-services.sh
```

### Playwright Commands

```bash
# Run all E2E tests
npm run test:e2e

# Run E2E tests with UI
npm run test:e2e:ui

# Run specific test file
npx playwright test e2e/dashboard.spec.ts

# Run tests in headed mode
npx playwright test --headed

# View test report
npx playwright show-report
```

## Service Management

### Starting Services

The `start-services.sh` script:
1. Starts infrastructure (PostgreSQL, RabbitMQ)
2. Waits for infrastructure to be ready
3. Starts application services
4. Verifies all services are healthy

```bash
./scripts/start-services.sh
```

### Stopping Services

```bash
./scripts/stop-services.sh
```

### Verifying Services

Check if all services are healthy:

```bash
./scripts/verify-services.sh
```

This checks:
- Content Service: `http://localhost:8081/health`
- Scoring Service: `http://localhost:8082/health`
- Learner Model Service: `http://localhost:8083/health`
- Adaptive Engine Service: `http://localhost:8084/health`
- PostgreSQL: Docker container health
- RabbitMQ: Management UI accessibility

## Test Data Management

### Expected Test Data State

#### Content Service (content_db)

- **Math Questions**: 
  - 5 remedial questions (difficulty 1)
  - 5 standard questions (difficulty 2)
- **Science Questions**:
  - 5 remedial questions (difficulty 1)
  - 5 standard questions (difficulty 2)

#### Learner Model Service (learner_db)

- **Test User**: `test-user-123`
  - Math skill: 50% mastery
  - Science skill: 60% mastery

### Setting Up Test Data

```bash
./scripts/setup-test-data.sh
```

This script:
1. Inserts test questions into Content Service
2. Creates test user mastery data in Learner Model Service
3. Verifies data was inserted correctly

### Cleaning Up Test Data

```bash
./scripts/cleanup-test-data.sh
```

This removes:
- All test questions (math, science) from Content Service
- Test user mastery data from Learner Model Service

## Troubleshooting

### Services Not Starting

**Problem**: Services fail to start or are not accessible.

**Solutions**:
1. Check Docker is running:
   ```bash
   docker info
   ```

2. Check if ports are already in use:
   ```bash
   lsof -i :8081  # Content Service
   lsof -i :8082  # Scoring Service
   lsof -i :8083  # Learner Model
   lsof -i :8084  # Adaptive Engine
   ```

3. Check service logs:
   ```bash
   docker-compose -f docker-compose.yml logs
   docker-compose -f docker-compose.infra.yml logs
   ```

4. Restart services:
   ```bash
   ./scripts/stop-services.sh
   ./scripts/start-services.sh
   ```

### Services Not Healthy

**Problem**: `verify-services.sh` reports services as unhealthy.

**Solutions**:
1. Wait a bit longer - services may still be starting (up to 60 seconds)
2. Check service logs for errors
3. Verify Docker containers are running:
   ```bash
   docker ps
   ```
4. Check service health endpoints manually:
   ```bash
   curl http://localhost:8081/health
   curl http://localhost:8082/health
   curl http://localhost:8083/health
   curl http://localhost:8084/health
   ```

### Test Data Issues

**Problem**: Tests fail because test data is missing or incorrect.

**Solutions**:
1. Re-run test data setup:
   ```bash
   ./scripts/setup-test-data.sh
   ```

2. Verify test data exists:
   ```bash
   # Check Content Service questions
   docker exec its-postgres psql -U postgres -d content_db -c "SELECT COUNT(*) FROM questions WHERE skill_tag IN ('math', 'science');"
   
   # Check Learner Model mastery
   docker exec its-postgres psql -U postgres -d learner_db -c "SELECT * FROM skill_mastery WHERE user_id = 'test-user-123';"
   ```

3. Clean and re-setup:
   ```bash
   ./scripts/cleanup-test-data.sh
   ./scripts/setup-test-data.sh
   ```

### E2E Tests Failing

**Problem**: E2E tests fail with connection errors or timeouts.

**Solutions**:
1. Verify all services are running:
   ```bash
   ./scripts/verify-services.sh
   ```

2. Check if Next.js dev server is running (Playwright config starts it automatically)

3. Check test logs:
   ```bash
   npx playwright show-report
   ```

4. Run tests in headed mode to see what's happening:
   ```bash
   npx playwright test --headed
   ```

5. Check network connectivity:
   ```bash
   curl http://localhost:3000  # Client app
   curl http://localhost:8084/api/adaptive/next-lesson?user_id=test-user-123&skill_tag=math
   ```

### Database Connection Issues

**Problem**: Cannot connect to PostgreSQL.

**Solutions**:
1. Verify PostgreSQL is running:
   ```bash
   docker ps | grep postgres
   ```

2. Check PostgreSQL logs:
   ```bash
   docker logs its-postgres
   ```

3. Test connection manually:
   ```bash
   docker exec its-postgres pg_isready -U postgres
   ```

### RabbitMQ Connection Issues

**Problem**: Services cannot connect to RabbitMQ.

**Solutions**:
1. Verify RabbitMQ is running:
   ```bash
   docker ps | grep rabbitmq
   ```

2. Check RabbitMQ management UI:
   ```bash
   open http://localhost:15672
   # Login: admintest / adminTest2025
   ```

3. Check RabbitMQ logs:
   ```bash
   docker logs its-rabbitmq
   ```

## Additional Resources

- [Playwright Documentation](https://playwright.dev/docs/intro)
- [Jest Documentation](https://jestjs.io/docs/getting-started)
- [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)

## Service Requirements Summary

### Infrastructure Services

| Service | Port | Health Check | Dependencies |
|---------|------|--------------|--------------|
| PostgreSQL | 5432 | `pg_isready` | None |
| RabbitMQ | 5672, 15672 | HTTP 15672 | None |

### Application Services

| Service | Port | Health Check | Dependencies |
|---------|------|--------------|--------------|
| Content Service | 8081 | `/health` | PostgreSQL (content_db) |
| Scoring Service | 8082 | `/health` | PostgreSQL (scoring_db), RabbitMQ |
| Learner Model API | 8083 | `/health` | PostgreSQL (learner_db) |
| Learner Model Consumer | - | N/A | PostgreSQL (learner_db), RabbitMQ |
| Adaptive Engine | 8084 | `/health` | Content Service, Learner Model API |

### Client Application

| Service | Port | Health Check | Dependencies |
|---------|------|--------------|--------------|
| Next.js Client | 3000 | HTTP 3000 | All backend services |

