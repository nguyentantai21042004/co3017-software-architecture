# Tasks: Setup E2E Testing Environment

This work enables running E2E tests against both local and deployed test environments using Playwright with antigravity web browser integration.

## 1. Local E2E Testing Setup

- [ ] Create service startup scripts:
    - [ ] Create Docker Compose file for all backend services (Content, Scoring, Learner Model, Adaptive Engine)
    - [ ] Create individual service startup scripts (alternative to Docker Compose)
    - [ ] Create service health check script to verify all services are running
    - [ ] Document service requirements and dependencies

- [ ] Create test data setup:
    - [ ] Create test data initialization scripts for Content Service (questions, skills)
    - [ ] Create test user setup scripts for Learner Model Service
    - [ ] Document expected test data state
    - [ ] Create test data cleanup scripts

- [ ] Create test execution scripts (if up to date, skip):
    - [ ] Create `scripts/start-services.sh` to start all required services
    - [ ] Create `scripts/stop-services.sh` to stop all services
    - [ ] Create `scripts/verify-services.sh` to check service health before tests
    - [ ] Create `scripts/run-e2e-local.sh` to run E2E tests with service verification
    - [ ] Update `package.json` with new test scripts

- [ ] Update documentation:
    - [ ] Update `TESTING.md` with local E2E setup instructions
    - [ ] Create troubleshooting guide for common E2E test issues
    - [ ] Document service startup order and dependencies

## 2. Test Environment Configuration

- [ ] Create environment configuration system:
    - [ ] Create `.env.local` for local development
    - [ ] Create `.env.test` for test environment
    - [ ] Create `.env.staging` for staging environment (optional)
    - [ ] Update `.gitignore` to exclude environment files with secrets

- [ ] Update API service for environment-based URLs:
    - [ ] Modify `services/api.ts` to read API URLs from environment variables
    - [ ] Create environment variable validation
    - [ ] Add fallback to default localhost URLs if env vars not set

- [ ] Update Playwright configuration:
    - [ ] Modify `playwright.config.ts` to support environment-based baseURL
    - [ ] Create environment-specific Playwright configs if needed
    - [ ] Add environment variable support for test execution

- [ ] Create environment switching utilities:
    - [ ] Create script to switch between environments
    - [ ] Document environment variable requirements
    - [ ] Create environment validation script

## 3. Deployed Test Environment Setup

- [ ] Test environment infrastructure:
    - [ ] Verify test environment exists and is accessible
    - [ ] Document test environment URLs and access methods
    - [ ] Verify all backend services are deployed in test environment
    - [ ] Test connectivity to test environment services

- [ ] Configure test environment URLs:
    - [ ] Update `API_URLS` in `services/api.ts` for test environment (or use env vars)
    - [ ] Update `baseURL` in `playwright.config.ts` for test environment
    - [ ] Verify test environment service endpoints are correct
    - [ ] Test API connectivity to test environment

- [ ] Create deployment verification:
    - [ ] Create script to verify test environment is ready
    - [ ] Create health check script for test environment services
    - [ ] Document test environment access procedures
    - [ ] Create test environment reset/cleanup procedures

## 4. Playwright E2E Test Suite with Antigravity Browser

This task focuses on implementing a comprehensive Playwright test suite using antigravity web browser for advanced browser automation and testing capabilities.

- [ ] **Antigravity Browser Integration**:
    - [ ] Install and configure antigravity web browser for Playwright
    - [ ] Update `playwright.config.ts` to use antigravity browser as primary test execution engine
    - [ ] Configure antigravity-specific options (selectors, performance monitoring, network interception)
    - [ ] Set up antigravity browser for different environments (local, test, staging)
    - [ ] Document antigravity browser setup and configuration

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

