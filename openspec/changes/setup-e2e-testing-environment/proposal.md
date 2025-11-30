# Change: Setup E2E Testing Environment

## Why

The client application has comprehensive E2E tests written with Playwright that cover the full user workflow (login, question presentation, answer submission, mastery updates, adaptive question selection). However, these tests require:
1. All backend services to be running and accessible
2. A properly configured test environment with all dependencies
3. Infrastructure setup for running tests against a deployed system

Currently, E2E tests are configured but cannot be executed because they require manual setup of backend services and test infrastructure. This proposal addresses the need to create a systematic approach for running E2E tests in both local and deployed environments.

## What Changes

### 1. Local E2E Testing Setup
- Create scripts to start all required backend services (Content, Scoring, Learner Model, Adaptive Engine)
- Create test data setup scripts for consistent test execution
- Document the complete setup process for local E2E testing
- Create a test runner script that verifies all services are running before executing tests

### 2. Deployed Test Environment Setup
- Configure test environment URLs in test configuration files
- Create environment-specific configuration files (local, test, staging)
- Create test environment deployment documentation

### 3. Playwright E2E Test Suite Implementation
- Implement comprehensive Playwright test suite using antigravity web browser for automated browser testing
- Create test cases for all critical user workflows (dashboard, learning session, API integration, UI components)
- Set up test data fixtures and mocking strategies
- Document test structure and maintenance procedures

### 4. Test Execution and Reporting
- Create scripts to run E2E tests with proper service verification
- Set up test result reporting and artifact collection
- Create test failure analysis and debugging guides
- Document test maintenance procedures

## Impact

- **Affected Specs:** 
  - `specs/system-integration-testing/spec.md` - MODIFIED (add E2E testing infrastructure requirements)
  - `specs/client-app-stabilization/spec.md` - MODIFIED (add Playwright E2E test suite requirements)
- **Affected Code:**
  - `sources/client/services/api.ts` - May need environment-based URL configuration
  - `sources/client/playwright.config.ts` - May need environment-specific configuration and antigravity browser integration
  - New scripts in `sources/client/scripts/` for service management
  - New configuration files for different environments
  - E2E test files in `sources/client/e2e/` - Enhanced with antigravity browser support
- **Breaking Changes:** None
- **Dependencies:** Requires all backend services to be deployable and accessible

## Implementation Approach

### Phase 1: Local E2E Testing Setup
1. Create service startup scripts (Docker Compose or individual service scripts)
2. Create test data initialization scripts
3. Create pre-test verification script (check all services are healthy)
4. Update `TESTING.md` with local E2E setup instructions
5. Test the complete local setup workflow

### Phase 2: Test Environment Configuration
1. Create environment configuration system (`.env.test`, `.env.staging`)
2. Update API service to use environment-based URLs
3. Update Playwright config to support multiple environments
4. Create environment switching scripts

### Phase 3: Playwright E2E Test Suite with Antigravity Browser
1. Configure Playwright to use antigravity web browser for automated testing
2. Implement comprehensive test suite covering dashboard, learning flows, API integration, and UI components
3. Set up test data fixtures and mocking
4. Create test maintenance documentation
5. Integrate antigravity browser capabilities for advanced browser automation

### Phase 4: Deployed Environment Testing
1. Set up test environment infrastructure (if not exists)
2. Configure test environment URLs
3. Create deployment verification scripts
4. Document test environment access and usage

## Success Criteria

1. ✅ All E2E tests can be run locally with a single command after services are started
2. ✅ Test environment is properly configured and accessible
3. ✅ E2E tests can be executed against deployed test environment
4. ✅ Test results are properly documented and reported
5. ✅ Test failures are properly logged and debuggable
6. ✅ Comprehensive Playwright test suite covers all critical user workflows using antigravity web browser
7. ✅ Antigravity browser integration enables advanced browser automation and testing capabilities

## Risks and Mitigations

- **Risk:** Backend services may not be available or may have different configurations
  - **Mitigation:** Create health check scripts and clear documentation for service requirements
- **Risk:** Test environment may have different data than expected
  - **Mitigation:** Create test data setup scripts and document expected state
- **Risk:** E2E tests may be flaky due to timing or network issues
  - **Mitigation:** Add proper waits, retries, and timeout configurations. Use antigravity browser's advanced capabilities for more reliable test execution.
- **Risk:** Test execution may be slow
  - **Mitigation:** Optimize test execution, use parallel execution where possible, leverage antigravity browser's performance optimizations
- **Risk:** Antigravity browser integration may require additional configuration
  - **Mitigation:** Document antigravity browser setup and configuration requirements clearly

## Open Questions

1. What is the preferred method for starting backend services? (Docker Compose, individual scripts, or manual)
2. Should test data be seeded automatically or manually?
3. What is the target test environment infrastructure? (Cloud provider, on-premise, etc.)
4. What specific antigravity browser features should be leveraged? (Advanced selectors, performance monitoring, network interception, etc.)

## Related Work

- This builds on the E2E tests created in `refactor-and-stabilize-microservices` change
- Requires backend services from Tasks 1-4 to be stable and deployable
- Extends `system-integration-testing` and `client-app-stabilization` specs with E2E testing infrastructure

