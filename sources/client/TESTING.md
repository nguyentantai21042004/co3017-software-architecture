# Testing Guide

This document provides comprehensive instructions for running tests in the client application, including unit tests, component tests, and end-to-end (E2E) tests.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Unit and Component Tests](#unit-and-component-tests)
- [E2E Testing Setup](#e2e-testing-setup)
- [Running E2E Tests](#running-e2e-tests)
- [Environment Configuration](#environment-configuration)
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

### E2E Test Structure

The E2E test suite is organized in `sources/client/e2e/`:

- **`specs/`**: Test files (e.g., `dashboard.spec.ts`, `learning-flow.spec.ts`)
  - `dashboard.spec.ts`: Dashboard functionality (loading, navigation, logout)
  - `learning-flow.spec.ts`: Full learning session flow (questions, feedback, mastery)
  - `api-integration.spec.ts`: Backend service integration verification
  - `ui-components.spec.ts`: UI component isolation tests
  - `error-handling.spec.ts`: Edge cases and error resilience
  - `test-data.spec.ts`: Data verification and mock mode
  - `antigravity-features.spec.ts`: Advanced browser features
- **`fixtures/`**: Test fixtures and setup
  - `antigravity-fixture.ts`: Custom Playwright fixture for Antigravity Browser
  - `mock-data.ts`: Centralized mock data objects
- **`utils/`**: Helper utilities
  - `artifacts.ts`: Helper for persistent screenshot paths

### Adding New E2E Tests

1. **Create a new spec file** in `sources/client/e2e/`.
2. **Import the custom fixture**:
   ```typescript
   import { test, expect } from './fixtures/antigravity-fixture';
   import { screenshotPath } from './utils/artifacts';
   ```
3. **Write your test**:
   ```typescript
   test('should perform action', async ({ agPage }) => {
     await agPage.goto('/page');
     await expect(agPage.locator('selector')).toBeVisible();
     await agPage.screenshot({ path: screenshotPath('action-result.png'), fullPage: true });
   });
   ```
4. **Use `agPage`** instead of `page` to leverage Antigravity features.
5. **Run the test** to verify:
   ```bash
   npx playwright test e2e/your-test.spec.ts
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

## Environment Configuration

The application supports multiple environments (local, test, staging) through environment variables. See [Environment Configuration Guide](docs/ENVIRONMENT_CONFIGURATION.md) for detailed documentation.

### Quick Reference

```bash
# Switch environment
npm run e2e:switch-env local    # Switch to local
npm run e2e:switch-env test     # Switch to test
npm run e2e:switch-env staging  # Switch to staging

# Validate environment
npm run e2e:validate-env

# Load environment variables (for scripts)
source ./scripts/load-env.sh [environment]
```

### Environment Variables

Required environment variables (all prefixed with `NEXT_PUBLIC_`):
- `NEXT_PUBLIC_CONTENT_API_URL` - Content Service URL
- `NEXT_PUBLIC_SCORING_API_URL` - Scoring Service URL
- `NEXT_PUBLIC_LEARNER_API_URL` - Learner Model Service URL
- `NEXT_PUBLIC_ADAPTIVE_API_URL` - Adaptive Engine Service URL
- `NEXT_PUBLIC_CLIENT_URL` - Client Application URL (for E2E tests)

See [Environment Configuration Guide](docs/ENVIRONMENT_CONFIGURATION.md) for complete documentation.

## Service Management

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

### E2E Artifacts (Reports, Screenshots, Videos)

Playwright E2E runs now store artifacts in **per-run folders** so that reports and screenshots from different runs do **not overwrite** each other:

- **Run identifier**
  - Each run is assigned a `PW_RUN_ID` (timestamp by default, or you can set it manually).
  - All artifact paths are scoped under this run ID.

- **Directories**
  - Test artifacts (traces, raw screenshots, videos, etc.):
    - `test-results/<PW_RUN_ID>/...`
  - HTML reports:
    - `playwright-report/<PW_RUN_ID>/index.html`
  - Named screenshots captured in tests (for documentation/debugging):
    - `test-results/<PW_RUN_ID>/screenshots/*.png`

- **Customising run id**

```bash
# Use a custom run id (e.g. CI build number or short git SHA)
PW_RUN_ID=build-123 npm run test:e2e

# Reports and screenshots will be stored under:
#   test-results/build-123/
#   playwright-report/build-123/
```

- **Opening older reports**
  - Open a specific run’s report:

```bash
npx playwright show-report playwright-report/<PW_RUN_ID>
```

### Interpreting Test Results

1. **Console Summary**:
   - When running via `run-e2e-test-env.sh`, a summary is printed at the end:
     ```
     =============================================
        E2E TEST SUMMARY
     =============================================
     Total Tests: 54
     ✅ Passed:   54
     ❌ Failed:   0
     ...
     ```

2. **HTML Report**:
   - Detailed view of all tests, steps, and artifacts.
   - **Traces**: Click on a failed test to see the full execution trace.
   - **Screenshots**: View captured screenshots for visual verification.
   - **Console Logs**: See browser console output (including Antigravity logs).
   - **Video**: Watch a video recording of the test execution.

3. **JSON Report**:
   - Raw data available in `test-results/<run-id>/results.json`.

## Environment Configuration

The application supports multiple environments (local, test, staging) through environment variables. See [Environment Configuration Guide](docs/ENVIRONMENT_CONFIGURATION.md) for detailed documentation.

### Quick Reference

```bash
# Switch environment
npm run e2e:switch-env local    # Switch to local
npm run e2e:switch-env test     # Switch to test
npm run e2e:switch-env staging  # Switch to staging

# Validate environment
npm run e2e:validate-env

# Load environment variables (for scripts)
source ./scripts/load-env.sh [environment]
```

### Environment Variables

Required environment variables (all prefixed with `NEXT_PUBLIC_`):
- `NEXT_PUBLIC_CONTENT_API_URL` - Content Service URL
- `NEXT_PUBLIC_SCORING_API_URL` - Scoring Service URL
- `NEXT_PUBLIC_LEARNER_API_URL` - Learner Model Service URL
- `NEXT_PUBLIC_ADAPTIVE_API_URL` - Adaptive Engine Service URL
- `NEXT_PUBLIC_CLIENT_URL` - Client Application URL (for E2E tests)

See [Environment Configuration Guide](docs/ENVIRONMENT_CONFIGURATION.md) for complete documentation.

### Test Environment Setup

For running E2E tests against a deployed test environment, see [Test Environment Setup Guide](docs/TEST_ENVIRONMENT_SETUP.md).

**Quick commands:**
```bash
# Verify test environment deployment
npm run e2e:verify-test-env [test|staging]

# Test API connectivity
npm run e2e:test-connectivity [test|staging]

# Reset test environment (template)
npm run e2e:reset-test-env [test|staging]
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
```bash
   curl http://localhost:8084/api/adaptive/next-lesson?user_id=test-user-123&skill_tag=math
   ```

### Debugging E2E Tests

1. **Use Antigravity Debugging Tools**:
   - **Console Logs**: Tests capture browser console logs. Check the test report for captured logs.
   - **Network Interception**: Use `agPage.on('request')` to log API calls.

2. **Run in Headed Mode**:
   ```bash
   npx playwright test --headed
   ```

3. **Use Playwright Inspector**:
   ```bash
   npx playwright test --debug
   ```

4. **Check Screenshots**:
   - Look at `test-results/<run-id>/screenshots/` for visual state at failure points.

### Handling Flaky Tests

1. **Wait for Network Idle**:
   - Use `await agPage.waitForLoadState('networkidle')` before assertions if the page loads data.

2. **Use Robust Selectors**:
   - Prefer `getByRole`, `getByText`, or `:has-text` over CSS classes.
   - Example: `agPage.getByRole('button', { name: 'Submit' })`

3. **Increase Timeouts (Sparingly)**:
   - If a specific action is slow, pass a timeout option: `await expect(locator).toBeVisible({ timeout: 10000 })`.

4. **Isolate State**:
   - Ensure `beforeEach` cleans up or sets up unique state (e.g., unique user IDs if possible, or reset data).

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

### Antigravity Browser Issues

**Problem**: `agPage` methods are not available or behave unexpectedly.

**Solutions**:
1. **Check Fixture Usage**: Ensure you import `test` from `./fixtures/antigravity-fixture` and destructure `agPage`.
2. **Check Initialization**: The fixture injects an init script. Verify `window.__ANTIGRAVITY_ENABLED__` in the browser console if debugging.
3. **Update Playwright**: Ensure `@playwright/test` is up to date.

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

