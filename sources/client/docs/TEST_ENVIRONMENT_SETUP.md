# Test Environment Setup Guide

This document describes how to set up, verify, and access a deployed test environment for E2E testing.

## Overview

The test environment is a deployed instance of the Intelligent Tutoring System that mirrors the production environment but is used exclusively for testing. E2E tests can be run against this environment to verify system behavior in a production-like setting.

## Test Environment Infrastructure

### Required Services

The test environment must have the following services deployed and accessible:

1. **Content Service** - Provides questions and content
   - Health endpoint: `{CONTENT_URL}/health`
   - API endpoint: `{CONTENT_URL}/api/content/*`

2. **Scoring Service** - Handles answer scoring
   - Health endpoint: `{SCORING_URL}/health`
   - API endpoint: `{SCORING_URL}/api/scoring/*`

3. **Learner Model Service** - Manages user mastery data
   - Health endpoint: `{LEARNER_URL}/health`
   - API endpoint: `{LEARNER_URL}/internal/learner/*`

4. **Adaptive Engine Service** - Provides adaptive learning recommendations
   - Health endpoint: `{ADAPTIVE_BASE_URL}/health`
   - API endpoint: `{ADAPTIVE_URL}/api/adaptive/*`

5. **Client Application** (optional) - Frontend application
   - Accessible at: `{CLIENT_URL}`

### Infrastructure Requirements

- **Database**: PostgreSQL instances for each service
- **Message Queue**: RabbitMQ for asynchronous communication
- **Network**: Services must be accessible from the test execution environment
- **Security**: Appropriate firewall rules and access controls

## Accessing Test Environment

### Environment Configuration

Test environment URLs are configured in `.env.test` file:

```bash
NEXT_PUBLIC_ENV=test
NEXT_PUBLIC_CONTENT_API_URL=http://test-content-service:8081
NEXT_PUBLIC_SCORING_API_URL=http://test-scoring-service:8082
NEXT_PUBLIC_LEARNER_API_URL=http://test-learner-service:8083
NEXT_PUBLIC_ADAPTIVE_API_URL=http://test-adaptive-engine:8084/api/adaptive
NEXT_PUBLIC_CLIENT_URL=http://test-client:3000
```

**Note**: Replace placeholder URLs with actual test environment URLs.

### Setting Up Configuration

1. **Copy the template:**
   ```bash
   cp .env.test.example .env.test
   ```

2. **Update with actual URLs:**
   ```bash
   # Edit .env.test with your test environment URLs
   nano .env.test
   ```

3. **Or use the switch script:**
   ```bash
   ./scripts/switch-env.sh test
   # Then edit .env.local with test environment URLs
   ```

## Verifying Test Environment

### Quick Verification

Use the verification script to check if the test environment is accessible:

```bash
# Verify test environment deployment
./scripts/verify-test-env.sh test

# Or for staging
./scripts/verify-test-env.sh staging
```

This script:
- Loads environment configuration
- Checks connectivity to all services
- Verifies health endpoints
- Reports any inaccessible services

### Detailed API Connectivity Test

Test actual API endpoints to verify services are responding correctly:

```bash
# Test API connectivity
./scripts/test-env-connectivity.sh test

# Or for staging
./scripts/test-env-connectivity.sh staging
```

This script:
- Tests health endpoints
- Tests actual API endpoints
- Verifies service responses
- Reports detailed connectivity status

### Manual Verification

You can also verify manually:

```bash
# Check Content Service
curl http://test-content-service:8081/health

# Check Scoring Service
curl http://test-scoring-service:8082/health

# Check Learner Model Service
curl http://test-learner-service:8083/health

# Check Adaptive Engine Service
curl http://test-adaptive-engine:8084/health
```

## Running E2E Tests Against Test Environment

### Using npm Scripts

```bash
# Run E2E tests against test environment
npm run e2e:run-test-env test

# Or for staging
npm run e2e:run-test-env staging
```

### Using Scripts Directly

```bash
# Run E2E tests against test environment
./scripts/run-e2e-test-env.sh test
```

This script:
1. Loads environment configuration
2. Validates environment variables
3. Verifies test environment connectivity
4. Runs Playwright E2E tests

## Test Environment Reset and Cleanup

### Test Data Management

Test environment should have its own test data that is separate from production:

1. **Test Users**: Dedicated test user accounts
2. **Test Questions**: Pre-defined test questions
3. **Test Mastery Data**: Initial mastery scores for test users

### Cleanup Procedures

Before running E2E tests, you may want to reset test data:

```bash
# Note: Cleanup scripts should be adapted for test environment
# They may need to connect to test environment databases

# Example: Connect to test environment database
psql -h test-db-host -U postgres -d content_db -f cleanup.sql
```

**Important**: Ensure cleanup procedures do not affect production data.

## Troubleshooting

### Services Not Accessible

**Problem**: Verification script reports services as not accessible.

**Solutions**:
1. Verify test environment is deployed and running
2. Check service URLs in `.env.test` are correct
3. Verify network connectivity from your machine to test environment
4. Check firewall rules allow access to test environment
5. Verify services are healthy in the test environment

### API Endpoints Returning Errors

**Problem**: API connectivity tests fail or return errors.

**Solutions**:
1. Verify services are properly deployed and configured
2. Check service logs for errors
3. Verify database connections in test environment
4. Check service dependencies (PostgreSQL, RabbitMQ) are running
5. Verify environment variables in test environment services

### Environment Configuration Issues

**Problem**: Tests are using wrong URLs or configuration.

**Solutions**:
1. Verify `.env.test` file exists and has correct values
2. Run `npm run e2e:validate-env` to validate configuration
3. Check environment variables are prefixed with `NEXT_PUBLIC_`
4. Restart Next.js dev server if running locally

### Network Connectivity Issues

**Problem**: Cannot reach test environment services.

**Solutions**:
1. Verify you have network access to test environment
2. Check VPN connection if required
3. Verify DNS resolution for test environment hostnames
4. Test connectivity with `ping` or `telnet`
5. Check proxy settings if behind corporate proxy

## Best Practices

1. **Separate Test Data**: Always use separate test data that does not affect production
2. **Regular Verification**: Run verification scripts before running E2E tests
3. **Documentation**: Keep test environment URLs and access methods documented
4. **Access Control**: Ensure proper access controls are in place
5. **Monitoring**: Monitor test environment health and performance
6. **Cleanup**: Regularly clean up test data to prevent accumulation

## Related Documentation

- [Environment Configuration Guide](ENVIRONMENT_CONFIGURATION.md) - Detailed environment configuration
- [Testing Guide](../TESTING.md) - General testing documentation
- [E2E Test Data State](E2E_TEST_DATA_STATE.md) - Test data requirements

## Scripts Reference

- `scripts/verify-test-env.sh` - Verify test environment deployment
- `scripts/test-env-connectivity.sh` - Test API connectivity
- `scripts/run-e2e-test-env.sh` - Run E2E tests against test environment
- `scripts/validate-env.sh` - Validate environment configuration

