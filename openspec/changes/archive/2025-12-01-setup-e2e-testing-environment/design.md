# Design: E2E Testing Environment Setup

## Overview

This design document outlines the technical approach for setting up a comprehensive E2E testing environment that enables reliable execution of Playwright tests against both local and deployed test environments. The design emphasizes the use of antigravity web browser for advanced browser automation capabilities.

## Architecture

### Test Execution Flow

```
┌─────────────────┐
│  Test Runner    │
│  (Playwright)   │
└────────┬────────┘
         │
         ├─► Service Health Checks
         │   └─► Verify all backend services are running
         │
         ├─► Test Data Setup
         │   └─► Initialize test data in databases
         │
         ├─► Antigravity Browser
         │   └─► Execute E2E tests with advanced automation
         │
         └─► Test Reporting
             └─► Generate test results and artifacts
```

### Service Dependencies

```
Client Application (Next.js)
    │
    ├─► Adaptive Engine Service
    │   ├─► Content Service
    │   └─► Learner Model Service
    │
    ├─► Scoring Service
    │
    └─► Learner Model Service
```

## Antigravity Web Browser Integration

### Why Antigravity Browser

Antigravity web browser provides advanced capabilities for E2E testing:
- **Advanced Selectors**: More reliable element selection and interaction
- **Performance Monitoring**: Built-in performance metrics collection
- **Network Interception**: Enhanced API mocking and request/response inspection
- **Cross-Browser Compatibility**: Consistent behavior across different browser engines
- **Debugging Tools**: Advanced debugging and trace collection capabilities

### Integration Approach

1. **Playwright Configuration**
   - Configure Playwright to use antigravity browser as the primary test execution engine
   - Set up antigravity-specific options in `playwright.config.ts`
   - Configure antigravity browser for different environments (local, test, staging)

2. **Test Implementation**
   - Leverage antigravity browser's advanced selectors for more reliable element identification
   - Use antigravity's network interception capabilities for API mocking and verification
   - Utilize performance monitoring features to track test execution metrics
   - Implement antigravity-specific debugging tools for test failure analysis

3. **Test Execution**
   - Run tests with antigravity browser in headless and headed modes
   - Collect antigravity browser traces and performance metrics
   - Generate test reports with antigravity-specific insights

## Environment Configuration

### Local Environment
- Services run via Docker Compose or individual scripts
- Client application runs on `localhost:3000`
- Backend services accessible via Docker network
- Test data initialized in local databases

### Test Environment
- Services deployed to test infrastructure
- Client application accessible via test URL
- Backend services accessible via test environment URLs
- Test data managed separately from production

### Configuration Files
- `.env.local` - Local development configuration
- `.env.test` - Test environment configuration
- `.env.staging` - Staging environment configuration (optional)

## Test Data Management

### Test Data Strategy
- **Fixtures**: Pre-defined test data sets for consistent test execution
- **Setup Scripts**: Automated scripts to initialize test data
- **Cleanup Scripts**: Automated scripts to clean up test data after execution
- **Isolation**: Each test should be able to run independently with its own data

### Test User Management
- Create dedicated test user accounts
- Manage test user sessions and authentication
- Clean up test user data after test execution

## Service Health Verification

### Health Check Strategy
1. **Pre-Test Verification**
   - Check all backend services are running
   - Verify service health endpoints respond correctly
   - Validate database connectivity
   - Confirm RabbitMQ is accessible

2. **Service Health Endpoints**
   - Content Service: `/health`
   - Scoring Service: `/health`
   - Learner Model Service: `/health`
   - Adaptive Engine Service: `/health`

3. **Failure Handling**
   - If services are not healthy, provide clear error messages
   - Suggest remediation steps
   - Log service status for debugging

## Test Execution Scripts

### Script Structure
- `scripts/start-services.sh` - Start all required services
- `scripts/stop-services.sh` - Stop all services
- `scripts/verify-services.sh` - Verify service health
- `scripts/run-e2e-local.sh` - Run E2E tests locally
- `scripts/run-e2e-test-env.sh` - Run E2E tests against test environment
- `scripts/setup-test-data.sh` - Initialize test data
- `scripts/cleanup-test-data.sh` - Clean up test data

## Test Reporting

### Report Types
- **HTML Reports**: Playwright HTML reporter for visual test results
- **JSON Reports**: Machine-readable test results for CI/CD integration
- **Artifacts**: Screenshots, videos, and traces for failed tests
- **Performance Metrics**: Antigravity browser performance data

### Artifact Collection
- Screenshots on test failure
- Video recordings of test execution
- Browser traces for debugging
- Network logs for API call analysis
- Antigravity browser-specific metrics

## Error Handling and Debugging

### Test Failure Analysis
1. **Automatic Artifact Collection**: Screenshots, videos, traces
2. **Service Status Logging**: Log service health at time of failure
3. **Network Request Logging**: Log all API calls and responses
4. **Antigravity Browser Debugging**: Use antigravity-specific debugging tools

### Common Issues and Solutions
- **Service Unavailable**: Check service health, verify Docker containers
- **Test Timeout**: Increase timeout values, check network connectivity
- **Flaky Tests**: Add proper waits, use antigravity browser's advanced selectors
- **Data Issues**: Verify test data setup, check database state

## Security Considerations

### Test Environment Security
- Test environment should be isolated from production
- Test data should not contain sensitive information
- Test credentials should be managed securely
- Environment files with secrets should be excluded from version control

## Performance Optimization

### Test Execution Optimization
- Run tests in parallel where possible
- Use test sharding for large test suites
- Optimize test data setup to minimize initialization time
- Leverage antigravity browser's performance optimizations

### Resource Management
- Limit concurrent test execution to avoid resource exhaustion
- Clean up resources after test execution
- Monitor system resources during test execution

