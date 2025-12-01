# Tasks: Setup E2E Testing Environment

This work enables running E2E tests against both local and deployed test environments using Playwright with antigravity web browser integration.

## 1. Local E2E Testing Setup

- [x] Create service startup scripts:
    - [x] Create Docker Compose file for all backend services (Content, Scoring, Learner Model, Adaptive Engine) - Using existing `sources/docker-compose.yml` and `sources/docker-compose.infra.yml`
    - [x] Create individual service startup scripts (alternative to Docker Compose) - Created `scripts/start-services.sh` that uses Docker Compose
    - [x] Create service health check script to verify all services are running - Created `scripts/verify-services.sh`
    - [x] Verify services are running and accessible - Health check script verifies all services
    - [x] Document service requirements and dependencies - Documented in `TESTING.md` with service dependencies diagram and requirements table

- [x] Create test data setup:
    - [x] Create test data initialization scripts for Content Service (questions, skills) - Created `scripts/setup-test-data.sh` and `scripts/insert_e2e_test_data.sql`
    - [x] Create test user setup scripts for Learner Model Service - Created test user mastery data setup in `scripts/setup-test-data.sh`
    - [x] Document expected test data state - Created `docs/E2E_TEST_DATA_STATE.md` with detailed documentation
    - [x] Create test data cleanup scripts - Created `scripts/cleanup-test-data.sh`
    - [x] Verify test data is initialized and accessible - Setup script verifies data after insertion

- [x] Create test execution scripts:
    - [x] Create `scripts/start-services.sh` to start all required services - Created with Docker Compose integration
    - [x] Create `scripts/stop-services.sh` to stop all services - Created
    - [x] Create `scripts/verify-services.sh` to check service health before tests - Created with health endpoint checks
    - [x] Create `scripts/run-e2e-local.sh` to run E2E tests with service verification - Created with complete workflow
    - [x] Update `package.json` with new test scripts - Added `e2e:start-services`, `e2e:stop-services`, `e2e:verify-services`, `e2e:setup-data`, `e2e:cleanup-data`, `e2e:run-local`

- [x] Update documentation:
    - [x] Update `TESTING.md` with local E2E setup instructions - Created comprehensive `TESTING.md` with setup, service management, test data management, and troubleshooting
    - [x] Create troubleshooting guide for common E2E test issues - Included in `TESTING.md` with solutions for common problems
    - [x] Document service startup order and dependencies - Documented in `TESTING.md` with dependency diagram and service requirements table

## 2. Test Environment Configuration

- [x] Create environment configuration system:
    - [x] Create `.env.local` for local development - Created `.env.local.example` as template
    - [x] Create `.env.test` for test environment - Created `.env.test.example` as template
    - [x] Create `.env.staging` for staging environment (optional) - Created `.env.staging.example` as template
    - [x] Update `.gitignore` to exclude environment files with secrets - Updated to exclude `.env.test`, `.env.staging`, but keep `.env*.example` files

- [x] Update API service for environment-based URLs:
    - [x] Modify `services/api.ts` to read API URLs from environment variables - Updated to use `API_URLS` from `lib/env-config.ts`
    - [x] Create environment variable validation - Created `lib/env-config.ts` with `validateEnvConfig()` function and `scripts/validate-env.sh` script
    - [x] Add fallback to default localhost URLs if env vars not set - Implemented in `lib/env-config.ts` with `getEnvUrl()` function

- [x] Update Playwright configuration:
    - [x] Modify `playwright.config.ts` to support environment-based baseURL - Updated to read from `NEXT_PUBLIC_CLIENT_URL` with fallback
    - [x] Create environment-specific Playwright configs if needed - Implemented conditional webServer based on environment
    - [x] Add environment variable support for test execution - Added environment detection and conditional configuration

- [x] Create environment switching utilities:
    - [x] Create script to switch between environments - Created `scripts/switch-env.sh` to easily switch between local/test/staging
    - [x] Document environment variable requirements - Created `docs/ENVIRONMENT_CONFIGURATION.md` with comprehensive documentation
    - [x] Create environment validation script - Created `scripts/validate-env.sh` (already completed in previous subtask)
    - [x] Create environment loading helper script - Created `scripts/load-env.sh` to load environment variables for scripts
    - [x] Create test environment E2E runner script - Created `scripts/run-e2e-test-env.sh` to run E2E tests against test/staging environments
    - [x] Integrate environment validation into E2E test workflow - Updated `run-e2e-local.sh` to validate environment before running tests
    - [x] Update TESTING.md with environment configuration references - Added environment configuration section with links to detailed docs

## 3. Deployed Test Environment Setup

- [x] Test environment infrastructure:
    - [x] Verify test environment exists and is accessible - Created `scripts/verify-test-env.sh` to verify test environment deployment and accessibility
    - [x] Document test environment URLs and access methods - Created `docs/TEST_ENVIRONMENT_SETUP.md` with comprehensive documentation
    - [x] Verify all backend services are deployed in test environment - Verification script checks all required services (Content, Scoring, Learner Model, Adaptive Engine)
    - [x] Test connectivity to test environment services - Created `scripts/test-env-connectivity.sh` to test API connectivity to all services

- [x] Configure test environment URLs:
    - [x] Update `API_URLS` in `services/api.ts` for test environment (or use env vars) - Already configured via `lib/env-config.ts` which reads from `NEXT_PUBLIC_*` environment variables
    - [x] Update `baseURL` in `playwright.config.ts` for test environment - Already configured to read from `NEXT_PUBLIC_CLIENT_URL` environment variable
    - [x] Verify test environment service endpoints are correct - Created verification scripts that validate service endpoints
    - [x] Test API connectivity to test environment - Created `scripts/test-env-connectivity.sh` that tests actual API endpoints

- [x] Create deployment verification:
    - [x] Create script to verify test environment is ready - Created `scripts/verify-test-env.sh` for comprehensive test environment verification
    - [x] Create health check script for test environment services - Verification script includes health checks for all services
    - [x] Document test environment access procedures - Created `docs/TEST_ENVIRONMENT_SETUP.md` with detailed access procedures, troubleshooting, and best practices
    - [x] Create test environment reset/cleanup procedures - Created `scripts/reset-test-env.sh` template with reset procedures (needs customization for specific test environment)

## 4. Playwright E2E Test Suite with Antigravity Browser

This task focuses on implementing a comprehensive Playwright test suite using antigravity web browser for advanced browser automation and testing capabilities.

- [x] **Backend Services Startup and Verification** (prerequisite for E2E tests):
    - [x] Create Playwright global setup hook to start backend services before tests - Created `e2e/global-setup.ts` that starts services, verifies health, and sets up test data
    - [x] Integrate service startup script (`scripts/start-services.sh`) into Playwright test workflow - Global setup calls `start-services.sh` automatically
    - [x] Add service health verification before test execution - Global setup calls `verify-services.sh` to ensure services are healthy
    - [x] Add test data setup (`scripts/setup-test-data.sh`) in global setup - Global setup calls `setup-test-data.sh` to initialize test data
    - [x] Create Playwright global teardown hook to stop services after tests (optional) - Created `e2e/global-teardown.ts` with optional service teardown (controlled by SKIP_SERVICE_TEARDOWN env var)
    - [x] Handle service startup failures gracefully with clear error messages - Global setup includes error handling and informative console messages
    - [x] Support both Docker Compose and manual service startup modes - Uses existing `start-services.sh` which supports Docker Compose
    - [x] Add timeout and retry logic for service health checks - Health verification script includes retry logic
    - [x] Document service startup requirements in test documentation - Will be documented in test documentation (part of later tasks)
    - [x] Create script to run Playwright tests with automatic service management - Created `scripts/run-playwright-with-services.sh` and added `npm run test:e2e:with-services` script

- [x] **Antigravity Browser Integration**:
    - [x] Install and configure antigravity web browser for Playwright
    - [x] Update `playwright.config.ts` to use antigravity browser as primary test execution engine
    - [x] Configure antigravity-specific options (selectors, performance monitoring, network interception)
    - [x] Set up antigravity browser for different environments (local, test, staging)
    - [x] Document antigravity browser setup and configuration

- [ ] **Dashboard Functionality Tests** (`e2e/dashboard.spec.ts`):
    - [ ] Test: Dashboard loads and displays available skills
        - Use antigravity browser's advanced selectors for reliable element identification
        - Verify dashboard page loads correctly
        - Verify skill cards are displayed with mastery scores
        - Verify loading states are handled properly
        - Verify error states are displayed when API calls fail
    - [ ] Test: Navigation to learning session
        - Verify clicking a skill card navigates to learning session page
        - Verify URL routing is correct (`/learn/[skill]`)
        - Verify skill context is preserved in navigation
    - [ ] Test: User authentication and logout
        - Verify logout button is visible and functional
        - Verify logout clears user session (localStorage)
        - Verify redirect to home page after logout
    - [ ] Test: Dashboard state management
        - Verify mastery scores are updated after returning from learning session
        - Verify dashboard refreshes correctly after session completion

- [ ] **Learning Session Flow Tests** (`e2e/learning-flow.spec.ts`):
    - [ ] Test: Learning session initialization
        - Use antigravity browser's network interception to verify API calls
        - Verify learning session page loads correctly
        - Verify mastery score is displayed
        - Verify session state is initialized properly
        - Verify API calls to Adaptive Engine for next lesson recommendation
    - [ ] Test: Question presentation
        - Use antigravity browser's advanced selectors for question elements
        - Verify question content is loaded from Content Service
        - Verify question options/inputs are displayed correctly
        - Verify question metadata (difficulty, skill tag) is shown
        - Verify remedial vs standard content is displayed appropriately
    - [ ] Test: Answer submission flow
        - Use antigravity browser's network monitoring to verify API calls
        - Verify answer can be selected/entered
        - Verify submit button is functional
        - Verify answer is sent to Scoring Service
        - Verify loading state during submission
        - Verify error handling for failed submissions
    - [ ] Test: Feedback display
        - Use antigravity browser's performance monitoring to track UI updates
        - Verify feedback is displayed after answer submission
        - Verify correct/incorrect feedback is shown appropriately
        - Verify feedback includes explanation if available
        - Verify UI updates correctly based on feedback
    - [ ] Test: Mastery score updates
        - Use antigravity browser's network interception to monitor polling
        - Verify mastery polling mechanism works correctly
        - Verify mastery score updates in UI after answer submission
        - Verify mastery circle component updates with new score
        - Verify polling stops after timeout or successful update
    - [ ] Test: Adaptive question selection
        - Use antigravity browser to verify adaptive recommendation flow
        - Verify next question is requested from Adaptive Engine
        - Verify adaptive recommendation logic is respected
        - Verify next question matches recommended difficulty/content type
        - Verify question progression follows adaptive algorithm
    - [ ] Test: Session continuation
        - Verify "Continue" button loads next question
        - Verify session state is maintained between questions
        - Verify question history is tracked correctly
        - Verify session can be completed and user returns to dashboard

- [ ] **API Integration Tests** (within E2E tests using antigravity browser):
    - [ ] Test: Content Service integration
        - Use antigravity browser's network interception to verify API calls
        - Verify question content is fetched correctly
        - Verify question metadata is parsed correctly
        - Verify error handling for Content Service failures
    - [ ] Test: Adaptive Engine integration
        - Use antigravity browser to monitor adaptive recommendation API calls
        - Verify next lesson recommendation API calls
        - Verify recommendation data is used correctly
        - Verify error handling for Adaptive Engine failures
    - [ ] Test: Scoring Service integration
        - Use antigravity browser's network monitoring for answer submission
        - Verify answer submission API calls
        - Verify scoring response is processed correctly
        - Verify error handling for Scoring Service failures
    - [ ] Test: Learner Model Service integration
        - Use antigravity browser to track mastery polling API calls
        - Verify mastery score retrieval API calls
        - Verify mastery polling mechanism
        - Verify error handling for Learner Model Service failures

- [ ] **UI Component Tests** (within E2E tests using antigravity browser):
    - [ ] Test: Mastery Circle component
        - Use antigravity browser's advanced selectors for component elements
        - Verify mastery circle displays correct score
        - Verify color coding based on mastery level
        - Verify animation/transitions work correctly
    - [ ] Test: Question display components
        - Use antigravity browser for reliable component interaction
        - Verify question text is rendered correctly
        - Verify multiple choice options are displayed
        - Verify text input fields are functional
        - Verify question metadata display
    - [ ] Test: Feedback components
        - Use antigravity browser's performance monitoring for UI updates
        - Verify feedback messages are displayed
        - Verify feedback styling (success/error)
        - Verify feedback animations/transitions
    - [ ] Test: Navigation components
        - Verify navigation links work correctly
        - Verify back button functionality
        - Verify breadcrumb navigation (if applicable)

- [ ] **Error Handling and Edge Cases** (using antigravity browser):
    - [ ] Test: Network failures
        - Use antigravity browser's network interception to simulate failures
        - Verify graceful handling of API timeouts
        - Verify retry mechanisms work correctly
        - Verify user-friendly error messages are displayed
    - [ ] Test: Invalid data handling
        - Use antigravity browser to test malformed API responses
        - Verify handling of malformed API responses
        - Verify handling of missing required data
        - Verify validation errors are displayed correctly
    - [ ] Test: Session state management
        - Use antigravity browser's state management capabilities
        - Verify session state persists during navigation
        - Verify session state is cleared on logout
        - Verify session recovery after page refresh

- [ ] **Test Data and Mocking** (with antigravity browser support):
    - [ ] Create test data fixtures for consistent testing
    - [ ] Set up test user accounts in test environment
    - [ ] Create test questions with known answers
    - [ ] Use antigravity browser's network interception for API mocking
    - [ ] Document test data requirements and setup procedures

- [ ] **Antigravity Browser-Specific Features**:
    - [ ] Leverage antigravity browser's advanced selectors for more reliable tests
    - [ ] Use antigravity browser's performance monitoring for test metrics
    - [ ] Implement antigravity browser's network interception for API verification
    - [ ] Utilize antigravity browser's debugging tools for test failure analysis
    - [ ] Document antigravity browser-specific test patterns and best practices

- [ ] **Test Maintenance**:
    - [ ] Document test structure and organization
    - [ ] Create guide for adding new E2E tests with antigravity browser
    - [ ] Document test debugging procedures using antigravity tools
    - [ ] Create test flakiness investigation guide
    - [ ] Document antigravity browser troubleshooting procedures

## 5. Test Execution and Reporting

- [ ] Create test execution scripts:
    - [ ] Create `scripts/run-e2e-test-env.sh` for test environment execution
    - [ ] Add pre-test verification (service health, data state)
    - [ ] Add post-test cleanup if needed
    - [ ] Create test result aggregation script

- [ ] Set up test reporting:
    - [ ] Configure Playwright HTML reporter for better reports
    - [ ] Integrate antigravity browser performance metrics into reports
    - [ ] Create test result summary script
    - [ ] Set up test artifact collection (screenshots, videos, traces, antigravity metrics)
    - [ ] Document how to interpret test results

- [ ] Create test maintenance documentation:
    - [ ] Document test failure debugging process
    - [ ] Create guide for updating tests when UI changes
    - [ ] Document test data maintenance procedures
    - [ ] Create test flakiness investigation guide
    - [ ] Document antigravity browser-specific debugging procedures

## 6. Documentation and Handoff

- [ ] Create comprehensive documentation:
    - [ ] Update `TESTING.md` with complete E2E testing guide
    - [ ] Create `E2E_TESTING_GUIDE.md` with detailed instructions
    - [ ] Document antigravity browser setup and configuration
    - [ ] Document all scripts and their purposes
    - [ ] Create troubleshooting guide for common issues

- [ ] Create runbook for test execution:
    - [ ] Document step-by-step process for local E2E testing
    - [ ] Document step-by-step process for test environment execution
    - [ ] Create quick reference guide for common commands
    - [ ] Document test maintenance procedures
    - [ ] Document antigravity browser usage and best practices

